package partners

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"context"
)

func ProvideServer(
	adapter *Adapter,
) *Server {
	return &Server{
		adapter: adapter,
	}
}

type Server struct {
	adapter *Adapter
	pb.UnimplementedPartnersServer
}

var _ pb.PartnersServer = &Server{}

func (s Server) Partner(ctx context.Context, request *pb.PartnerRequest) (*pb.PartnerReply, error) {
	return s.adapter.Partner(ctx, request)
}

func (s Server) PartnersList(ctx context.Context, request *pb.PartnersListRequest) (*pb.PartnersListReply, error) {
	return s.adapter.PartnersList(ctx, request)
}

func (s Server) PartnersSave(ctx context.Context, request *pb.PartnersSaveRequest) (*pb.PartnersSaveReply, error) {
	return s.adapter.PartnersSave(ctx, request)
}

func (s Server) PartnersDelete(ctx context.Context, request *pb.PartnersDeleteRequest) (*pb.PartnersDeleteReply, error) {
	return s.adapter.PartnersDelete(ctx, request)
}

func (s Server) PartnerRoleList(ctx context.Context, request *pb.PartnerRoleListRequest) (*pb.PartnerRoleListReply, error) {
	return s.adapter.PartnerRoleList(ctx, request)
}

func (s Server) PartnerRoleSave(ctx context.Context, request *pb.PartnerRoleSaveRequest) (*pb.PartnerRoleSaveReply, error) {
	return s.adapter.PartnerRoleSave(ctx, request)
}

func (s Server) PartnerRoleDelete(ctx context.Context, request *pb.PartnerRoleDeleteRequest) (*pb.PartnerRoleDeleteReply, error) {
	return s.adapter.PartnerRoleDelete(ctx, request)
}
