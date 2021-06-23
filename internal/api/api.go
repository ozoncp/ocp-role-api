package api

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-role-api/internal/model"
	"github.com/ozoncp/ocp-role-api/internal/producer"
	"github.com/ozoncp/ocp-role-api/internal/repo"
	"github.com/ozoncp/ocp-role-api/internal/utils"
	pb "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const chunkSize = 10

type ApiServer = pb.OcpRoleApiServer

type api struct {
	pr producer.Producer
	r  repo.Repo
	pb.UnimplementedOcpRoleApiServer
}

var internalError = status.Error(codes.Internal, "Internal server error")

func NewOcpRoleApi(repo repo.Repo, producer producer.Producer) ApiServer {
	return &api{
		pr: producer,
		r:  repo,
	}
}

func (a *api) ListRolesV1(ctx context.Context, req *pb.ListRolesV1Request) (*pb.ListRolesV1Response, error) {
	log.Info().Msgf("ListRolesV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	roles, err := a.r.ListRoles(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, internalError
	}

	resp := pb.ListRolesV1Response{
		Roles: make([]*pb.Role, len(roles)),
	}

	for i := range roles {
		role := &pb.Role{}
		role.Id = roles[i].Id
		role.Operation = roles[i].Operation
		role.Service = roles[i].Service
		resp.Roles[i] = role
	}

	return &resp, nil
}

func (a *api) DescribeRoleV1(ctx context.Context, req *pb.DescribeRoleV1Request) (*pb.DescribeRoleV1Response, error) {
	log.Info().Msgf("DescribeRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	role, err := a.r.DescribeRole(ctx, req.RoleId)
	if err != nil {
		return nil, internalError
	}

	return &pb.DescribeRoleV1Response{
		Role: &pb.Role{
			Id:        role.Id,
			Service:   role.Service,
			Operation: role.Operation,
		},
	}, nil
}

func (a *api) CreateRoleV1(ctx context.Context, req *pb.CreateRoleV1Request) (*pb.CreateRoleV1Response, error) {
	log.Info().Msgf("CreateRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := a.r.AddRole(
		ctx,
		&model.Role{
			Service:   req.Service,
			Operation: req.Operation,
		},
	)

	if err != nil {
		return nil, internalError
	}

	return &pb.CreateRoleV1Response{RoleId: id}, nil
}

func (a *api) MultiCreateRoleV1(
	ctx context.Context, req *pb.MultiCreateRoleV1Request) (*pb.MultiCreateRoleV1Response, error) {

	log.Info().Msgf("MultiCreateRoleV1: req: %v", req)

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateRoleV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	roles := make([]*model.Role, 0, len(req.Roles))

	for i := range req.Roles {
		roles = append(roles, &model.Role{
			Service:   req.Roles[i].Service,
			Operation: req.Roles[i].Operation,
		})
	}

	chunks := utils.SplitToBulks(roles, chunkSize)

	var resp pb.MultiCreateRoleV1Response

	for i := range chunks {
		err := func() error {
			childSpan := tracer.StartSpan(
				"MultiCreateRoleV1 child", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			ids, err := a.r.AddRoles(ctx, chunks[i])
			if err != nil {
				log.Error().Err(err)
				return err
			}

			resp.RoleIds = append(resp.RoleIds, ids...)

			for j := 0; j < len(ids); j++ {
				if err := a.pr.Send(producer.Event{
					Type: producer.Created,
					Body: map[string]interface{}{"role": chunks[i][j]},
				}); err != nil {
					log.Warn().Msgf("failed to send Created event: %v", err)
				}
			}

			return nil
		}()

		if err != nil {
			break
		}
	}

	return &resp, nil
}

func (a *api) UpdateRoleV1(ctx context.Context, req *pb.UpdateRoleV1Request) (*pb.UpdateRoleV1Response, error) {
	log.Info().Msgf("UpdateRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	role := model.Role{
		Id:        req.RoleId,
		Service:   req.Service,
		Operation: req.Operation,
	}
	found, err := a.r.UpdateRole(ctx, &role)

	if err != nil {
		return nil, internalError
	}

	if err := a.pr.Send(producer.Event{
		Type: producer.Updated,
		Body: map[string]interface{}{"role": role},
	}); err != nil {
		log.Warn().Msgf("failed to send Updated event: %v", err)
	}

	return &pb.UpdateRoleV1Response{Found: found}, nil
}

func (a *api) RemoveRoleV1(ctx context.Context, req *pb.RemoveRoleV1Request) (*pb.RemoveRoleV1Response, error) {
	log.Info().Msgf("RemoveRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	found, err := a.r.RemoveRole(ctx, req.RoleId)

	if err != nil {
		return nil, internalError
	}

	if err := a.pr.Send(producer.Event{
		Type: producer.Deleted,
		Body: map[string]interface{}{"role_id": req.RoleId},
	}); err != nil {
		log.Warn().Msgf("failed to send Updated event: %v", err)
	}

	return &pb.RemoveRoleV1Response{Found: found}, nil
}
