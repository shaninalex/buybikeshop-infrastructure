package partners

import (
	"buybikeshop/apps/datasource/app/models"
	pb "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"context"
)

type Adapter struct {
	repository *Repository
}

func ProvideAdapter(repository *Repository) *Adapter {
	return &Adapter{
		repository: repository,
	}
}

func (s Adapter) PartnersList(ctx context.Context, request *pb.PartnersListRequest) (*pb.PartnersListReply, error) {
	panic("implement me")
}

func (s Adapter) PartnersSave(ctx context.Context, request *pb.PartnersSaveRequest) (*pb.PartnersSaveReply, error) {
	panic("implement me")
}

func (s Adapter) PartnersDelete(ctx context.Context, request *pb.PartnersDeleteRequest) (*pb.PartnersDeleteReply, error) {
	panic("implement me")
}

func (s Adapter) PartnerRoleList(ctx context.Context, request *pb.PartnerRoleListRequest) (*pb.PartnerRoleListReply, error) {
	pr, err := s.repository.RolesGet(ctx)
	if err != nil {
		return nil, err
	}

	roles := make([]*pb.PartnerRole, len(pr))
	for i, role := range pr {
		roles[i] = &pb.PartnerRole{
			Id:   role.Id,
			Role: role.Role,
		}
	}

	return &pb.PartnerRoleListReply{
		Roles: roles,
	}, nil
}

func (s Adapter) PartnerRoleSave(ctx context.Context, request *pb.PartnerRoleSaveRequest) (*pb.PartnerRoleSaveReply, error) {
	r, err := s.repository.RolesSave(ctx, &models.PartnerRole{
		Id:   request.Role.Id,
		Role: request.Role.Role,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PartnerRoleSaveReply{
		Role: &pb.PartnerRole{
			Id:   r.Id,
			Role: r.Role,
		},
	}, nil
}

func (s Adapter) PartnerRoleDelete(ctx context.Context, request *pb.PartnerRoleDeleteRequest) (*pb.PartnerRoleDeleteReply, error) {
	if err := s.repository.RolesDelete(ctx, request.PartnerId); err != nil {
		return &pb.PartnerRoleDeleteReply{
			Status:  false,
			Message: err.Error(),
		}, err
	}
	return &pb.PartnerRoleDeleteReply{
		Status:  true,
		Message: "Deleted",
	}, nil
}
