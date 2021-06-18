package api

import (
	"context"

	pb "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	pb.UnimplementedOcpRoleApiServer
}

func NewOcpRoleApi() *api {
	return &api{}
}

func (a *api) ListRolesV1(ctx context.Context, req *pb.ListRolesV1Request) (*pb.ListRolesV1Response, error) {
	log.Info().Msgf("ListRolesV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, "")
}

func (a *api) DescribeRolesV1(ctx context.Context, req *pb.DescribeRoleV1Request) (*pb.DescribeRoleV1Response, error) {
	log.Info().Msgf("DescribeRolesV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, "")
}

func (a *api) CreateRoleV1(ctx context.Context, req *pb.CreateRoleV1Request) (*pb.CreateRoleV1Response, error) {
	log.Info().Msgf("CreateRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, "")
}

func (a *api) RemoveRoleV1(ctx context.Context, req *pb.RemoveRoleV1Request) (*pb.RemoveRoleV1Response, error) {
	log.Info().Msgf("RemoveRoleV1: req: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, "")
}
