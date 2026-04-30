package employee

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/employee"
	"context"
)

type Server struct {
	adapter *Adapter
	pb.UnimplementedEmployeeServiceServer
}

func ProvideServer(
	adapter *Adapter,
) *Server {
	return &Server{
		adapter: adapter,
	}
}

var _ pb.EmployeeServiceServer = (*Server)(nil)

func (s *Server) GetEmployee(ctx context.Context, data *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {
	return s.adapter.ProductGet(ctx, data)
}

func (s *Server) SaveEmployee(ctx context.Context, data *pb.SaveEmployeeRequest) (*pb.SaveEmployeeResponse, error) {
	return s.adapter.ProductSave(ctx, data)
}
