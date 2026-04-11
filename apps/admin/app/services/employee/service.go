package employee

import (
	"buybikeshop/apps/admin/app/models"
	"buybikeshop/libs/go/kratos"
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, data models.EmployeeCreate) (*models.Employee, error)
	List(ctx context.Context) ([]models.Employee, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func ProvideEmployeeService(c kratos.ApiClient) Service {
	s := serviceImpl{client: c}
	return &s
}

var _ Service = (*serviceImpl)(nil)

type serviceImpl struct {
	client kratos.ApiClient
}

var (
	ErrorCreate = errors.New("unable to create identity")
)

func (s serviceImpl) Create(ctx context.Context, data models.EmployeeCreate) (*models.Employee, error) {
	identity, err := s.client.CreateIdentity(ctx, data.Identity)
	if err != nil {
		return nil, err
	}

	return &models.Employee{Identity: *identity}, nil
}

func (s serviceImpl) List(ctx context.Context) (employees []models.Employee, err error) {
	identities, err := s.client.ListIdentities(ctx)
	if err != nil {
		return nil, err
	}

	for _, i := range identities {
		employees = append(employees, models.Employee{Identity: i})
	}

	return employees, nil
}

func (s serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	r, err := s.client.DeleteIdentity(ctx, id)
	if err != nil {
		return err
	}
	if !r {
		return ErrorCreate
	}
	return nil
}
