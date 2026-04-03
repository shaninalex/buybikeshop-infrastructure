package partners

import (
	"buybikeshop/apps/datasource/app/models"
	pb "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"context"
)

type Adapter struct {
	repositoryRoles    *RepositoryRoles
	repositoryPartners *RepositoryPartners
}

func ProvideAdapter(repositoryRoles *RepositoryRoles, repositoryPartners *RepositoryPartners) *Adapter {
	return &Adapter{
		repositoryRoles:    repositoryRoles,
		repositoryPartners: repositoryPartners,
	}
}

func (s Adapter) Partner(ctx context.Context, request *pb.PartnerRequest) (*pb.PartnerReply, error) {
	partner, err := s.repositoryPartners.Partner(ctx, request.PartnerId)
	if err != nil {
		return nil, err
	}
	return &pb.PartnerReply{
		Partner: models.ToPbPartner(partner),
	}, err
}

func (s Adapter) PartnersList(ctx context.Context, request *pb.PartnersListRequest) (*pb.PartnersListReply, error) {
	partners, err := s.repositoryPartners.PartnersList(ctx)
	if err != nil {
		return &pb.PartnersListReply{
			Partners: []*pb.Partner{},
			Total:    0,
		}, err
	}
	return &pb.PartnersListReply{
		Partners: models.ToPbPartners(partners),
		Total:    int64(len(partners)),
	}, err
}

func (s Adapter) PartnersSave(ctx context.Context, request *pb.PartnersSaveRequest) (*pb.PartnersSaveReply, error) {
	payload := models.ToModelPartner(request.Partner)
	partner, err := s.repositoryPartners.PartnersSave(ctx, payload)
	if err != nil {
		return &pb.PartnersSaveReply{
			Partner: nil,
			Status:  false,
		}, err
	}
	return &pb.PartnersSaveReply{
		Partner: models.ToPbPartner(partner),
		Status:  false,
	}, err
}

func (s Adapter) PartnersDelete(ctx context.Context, request *pb.PartnersDeleteRequest) (*pb.PartnersDeleteReply, error) {
	if err := s.repositoryPartners.PartnersDelete(ctx, request.PartnerId); err != nil {
		return &pb.PartnersDeleteReply{
			Status: false,
		}, err
	}
	return &pb.PartnersDeleteReply{
		Status: true,
	}, nil
}

func (s Adapter) PartnerRoleList(ctx context.Context, request *pb.PartnerRoleListRequest) (*pb.PartnerRoleListReply, error) {
	pr, err := s.repositoryRoles.RolesGet(ctx)
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
	r, err := s.repositoryRoles.RolesSave(ctx, &models.Role{
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
	if err := s.repositoryRoles.RolesDelete(ctx, request.RoleId); err != nil {
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
