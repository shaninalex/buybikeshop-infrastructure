package kratos

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type ApiClient interface {
	CreateIdentity(ctx context.Context, data ory.CreateIdentityBody) (*ory.Identity, error)
	ListIdentities(ctx context.Context) ([]ory.Identity, error)
	DeleteIdentity(ctx context.Context, id uuid.UUID) (bool, error)
}

var (
	IdentityApiErrorCreate       = errors.New("unable to create identity")
	IdentityApiErrorMissingEmail = errors.New(`missing "email" key`)
	IdentityApiErrorNotFound     = errors.New("identity not found")
)

var _ ApiClient = (*KratosApiClient)(nil)

func ProvideApiClient(client *ory.APIClient) ApiClient {
	return &KratosApiClient{client: client}
}

type KratosApiClient struct {
	client *ory.APIClient
}

func (s KratosApiClient) CreateIdentity(ctx context.Context, data ory.CreateIdentityBody) (*ory.Identity, error) {
	d := s.client.IdentityAPI.CreateIdentity(ctx).CreateIdentityBody(data)
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
		return nil, IdentityApiErrorCreate
	}
	return identity, nil
}

func (s KratosApiClient) ListIdentities(ctx context.Context) ([]ory.Identity, error) {
	identities, _, err := s.client.IdentityAPI.ListIdentitiesExecute(s.client.IdentityAPI.ListIdentities(ctx))
	if err != nil {
		return nil, err
	}
	return identities, nil
}

func (s KratosApiClient) DeleteIdentity(ctx context.Context, id uuid.UUID) (bool, error) {
	r, err := s.client.IdentityAPI.DeleteIdentityExecute(s.client.IdentityAPI.DeleteIdentity(ctx, id.String()))
	defer r.Body.Close()
	if err != nil {
		return false, err
	}
	return r.StatusCode == http.StatusNoContent, nil
}
