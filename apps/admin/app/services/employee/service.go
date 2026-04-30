package employee

import (
	"buybikeshop/apps/admin/app/models"
	"buybikeshop/libs/go/bus"
	"buybikeshop/libs/go/kratos"
	"context"
	"errors"

	"github.com/google/uuid"
)

type EmployeeCreate struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone"`
	Dob        string `json:"dob"`
	Photo      string `json:"photo"`
	Password   string `json:"password" validate:"required"`
	Group      string `json:"group" validate:"required"`
	Department string `json:"department" validate:"required"`
}

func (s *EmployeeCreate) ApplyDefaults() {
	if s.Photo == "" {
		s.Photo = "/images/default-avatar.png"
	}
}

type Service interface {
	Create(ctx context.Context, data EmployeeCreate) (*models.Employee, error)
	List(ctx context.Context) ([]models.Employee, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Validate(ctx context.Context, data EmployeeCreate) error
	Patch(ctx context.Context, id uuid.UUID, data EmployeeCreate) (*models.Employee, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Employee, error)
}

func ProvideEmployeeService(
	c kratos.ApiClient,
	bus bus.Bus,
) Service {
	s := serviceImpl{
		client: c,
		bus:    bus,
	}
	return &s
}

var _ Service = (*serviceImpl)(nil)

type serviceImpl struct {
	client kratos.ApiClient
	bus    bus.Bus
}

var (
	ErrorCreate = errors.New("unable to create identity")
)

func (s serviceImpl) Create(ctx context.Context, data EmployeeCreate) (*models.Employee, error) {
	identity, err := s.client.CreateIdentity(ctx, kratos.IdentityCreate{
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Dob:      data.Dob,
		Photo:    data.Photo,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}

	s.bus.Dispatch(ctx, bus.EmployeeCreatedEventType, data)

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

func (s serviceImpl) Get(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	identity, err := s.client.GetIdentity(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Employee{Identity: *identity}, nil
}

func (s serviceImpl) Patch(ctx context.Context, id uuid.UUID, data EmployeeCreate) (*models.Employee, error) {
	identity, err := s.client.PatchIdentity(ctx, id, kratos.IdentityCreate{
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Dob:      data.Dob,
		Photo:    data.Photo,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}
	return &models.Employee{Identity: *identity}, nil
}
