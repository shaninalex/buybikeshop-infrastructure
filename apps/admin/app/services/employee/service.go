package employee

import (
	"buybikeshop/apps/admin/app/models"
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type Service interface {
	Create(ctx context.Context, data models.EmployeeCreate) (*models.Employee, error)
	List(ctx context.Context) ([]models.Employee, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func ProvideEmployeeService(c *ory.APIClient) Service {
	s := serviceImpl{client: c}
	return &s
}

var _ Service = (*serviceImpl)(nil)

type serviceImpl struct {
	client *ory.APIClient
}

var (
	ErrorCreate = errors.New("unable to create identity")
)

func (s serviceImpl) Create(ctx context.Context, data models.EmployeeCreate) (*models.Employee, error) {
	d := s.client.IdentityAPI.CreateIdentity(ctx).CreateIdentityBody(data.Identity)
	identity, r, err := s.client.IdentityAPI.CreateIdentityExecute(d)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(r.Body)
		return nil, errors.New(string(b))
	}

	if identity == nil {
		return nil, ErrorCreate
	}

	return &models.Employee{Identity: *identity}, nil
}

func (s serviceImpl) List(ctx context.Context) (employees []models.Employee, err error) {
	identities, _, err := s.client.IdentityAPI.ListIdentitiesExecute(s.client.IdentityAPI.ListIdentities(ctx))
	if err != nil {
		return nil, err
	}

	for _, i := range identities {
		employees = append(employees, models.Employee{Identity: i})
	}

	return employees, nil
}

func (s serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	r, err := s.client.IdentityAPI.DeleteIdentityExecute(s.client.IdentityAPI.DeleteIdentity(ctx, id.String()))
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
