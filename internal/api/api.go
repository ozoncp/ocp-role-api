package api

import (
	"context"

	"github.com/ozoncp/ocp-role-api/internal/model"
	"github.com/ozoncp/ocp-role-api/internal/repo"
	pb "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiServer = pb.OcpRoleApiServer

type api struct {
	r repo.Repo
	pb.UnimplementedOcpRoleApiServer
}

var internalError = status.Error(codes.Internal, "Internal server error")

func NewOcpRoleApi(repo repo.Repo) ApiServer {
	return &api{
		r: repo,
	}
}

func (a *api) ListRolesV1(ctx context.Context, req *pb.ListRolesV1Request) (*pb.ListRolesV1Response, error) {
	log.Info().Msgf("ListRolesV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	roles, err := a.r.ListRoles(req.Limit, req.Offset)
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

	role, err := a.r.DescribeRole(req.RoleId)
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

	id, err := a.r.AddRole(&model.Role{
		Service:   req.Service,
		Operation: req.Operation,
	})

	if err != nil {
		return nil, internalError
	}

	return &pb.CreateRoleV1Response{RoleId: id}, nil
}

func (a *api) RemoveRoleV1(ctx context.Context, req *pb.RemoveRoleV1Request) (*pb.RemoveRoleV1Response, error) {
	log.Info().Msgf("RemoveRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	found, err := a.r.RemoveRole(req.RoleId)

	if err != nil {
		return nil, internalError
	}

	return &pb.RemoveRoleV1Response{Found: found}, nil
}
