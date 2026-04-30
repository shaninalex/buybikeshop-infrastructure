package employee

import (
	"buybikeshop/apps/datasource/app/models"
	pb "buybikeshop/gen/grpc-buybikeshop-go/employee"
	"context"

	"github.com/google/uuid"
)

type Adapter struct {
	repository *Repository
}

func ProvideAdapter(catalogRepository *Repository) *Adapter {
	return &Adapter{
		repository: catalogRepository,
	}
}

func (c Adapter) ProductSave(ctx context.Context, request *pb.SaveEmployeeRequest) (*pb.SaveEmployeeResponse, error) {
	department, err := c.repository.DepartmentGet(ctx, request.GetDepartment())
	if err != nil {
		return nil, err
	}

	empl := models.Employee{
		EmployeeId:   uuid.MustParse(request.EmployeeId),
		DepartmentId: department.Id,
	}

	p, err := c.repository.EmployeeSave(ctx, empl)
	if err != nil {
		return nil, err
	}
	return &pb.SaveEmployeeResponse{
		Employee: &pb.Employee{
			Id:           p.EmployeeId.String(),
			DepartmentId: p.DepartmentId,
		},
	}, nil
}

func (c Adapter) ProductGet(ctx context.Context, request *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {
	p, err := c.repository.EmployeeGet(ctx, uuid.MustParse(request.GetEmployeeId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetEmployeeResponse{
		Employee: &pb.Employee{
			Id:           p.EmployeeId.String(),
			DepartmentId: p.DepartmentId,
		},
	}, nil
}
