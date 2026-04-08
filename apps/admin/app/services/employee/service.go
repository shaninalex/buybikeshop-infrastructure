package employee

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type Service interface {
	Create(ctx context.Context, dataPath string) (*ory.Identity, error)
	List(ctx context.Context) ([]ory.Identity, error)
	Delete(ctx context.Context, id uuid.UUID) error
	EditIdentity(ctx context.Context, id uuid.UUID, patch ory.UpdateIdentityBody) (*ory.Identity, error)
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
	f, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var payload ory.CreateIdentityBody
	if err = json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}
	d := s.client.IdentityAPI.CreateIdentity(ctx).CreateIdentityBody(payload)
	i, r, err := s.client.IdentityAPI.CreateIdentityExecute(d)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(r.Body)
		log.Println(string(b))
		return nil, errors.New(string(b))
	}

	return i, nil
}

func (s serviceImpl) List(ctx context.Context) ([]ory.Identity, error) {
	identities, _, err := s.client.IdentityAPI.ListIdentitiesExecute(s.client.IdentityAPI.ListIdentities(ctx))
	if err != nil {
		return nil, err
	}
	return identities, nil
}

func (s serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	r, err := s.client.IdentityAPI.DeleteIdentityExecute(s.client.IdentityAPI.DeleteIdentity(ctx, id.String()))
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s serviceImpl) EditIdentity(ctx context.Context, id uuid.UUID, patch ory.UpdateIdentityBody) (*ory.Identity, error) {
	d := s.client.IdentityAPI.UpdateIdentity(ctx, id.String()).UpdateIdentityBody(patch)
	identity, _, err := s.client.IdentityAPI.UpdateIdentityExecute(d)
	if err != nil {
		return nil, err
	}
	return identity, nil
}
