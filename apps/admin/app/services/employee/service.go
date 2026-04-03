package employee

import (
	"context"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type Service interface {
	Create(ctx context.Context, dataPath string) (*ory.Identity, error)
	List(ctx context.Context) ([]*ory.Identity, error)
	Delete(ctx context.Context, id uuid.UUID) error
	EditIdentity(ctx context.Context, id uuid.UUID, patch *ory.IdentityPatch) (*ory.IdentityPatchResponse, error)
}

func ProvideEmployeeService(c *ory.APIClient) Service {
	s := serviceImpl{client: c}
	return &s
}

var _ Service = (*serviceImpl)(nil)

type serviceImpl struct {
	client *ory.APIClient
}

func (s serviceImpl) Create(ctx context.Context, dataPath string) (*ory.Identity, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) List(ctx context.Context) ([]*ory.Identity, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) EditIdentity(ctx context.Context, id uuid.UUID, patch *ory.IdentityPatch) (*ory.IdentityPatchResponse, error) {
	//TODO implement me
	panic("implement me")
}
