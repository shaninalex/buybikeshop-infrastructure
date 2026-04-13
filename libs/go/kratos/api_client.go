package kratos

import (
	"buybikeshop/libs/go/ptr"
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type EmployeeCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Photo    string `json:"photo"`
	Password string `json:"password"`
}

func (s *EmployeeCreate) ApplyDefaults() {
	if s.Photo == "" {
		s.Photo = "/images/default-avatar.png"
	}
}

type ApiClient interface {
	CreateIdentity(ctx context.Context, data EmployeeCreate) (*ory.Identity, error)
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

func (s KratosApiClient) CreateIdentity(ctx context.Context, data EmployeeCreate) (*ory.Identity, error) {
	body := ory.NewCreateIdentityBodyWithDefaults()
	body.Traits = map[string]any{
		"name":  data.Name,
		"email": data.Email,
		"phone": data.Phone,
		"dob":   data.Dob,
		"photo": data.Photo,
	}

	creds := ory.NewIdentityWithCredentials()
	pswd := ory.NewIdentityWithCredentialsPassword()
	pswd.SetConfig(ory.IdentityWithCredentialsPasswordConfig{
		Password: ptr.Ptr(data.Password),
	})

	creds.Password = pswd
	body.SetCredentials(*creds)

	d := s.client.IdentityAPI.CreateIdentity(ctx).CreateIdentityBody(*body)
	identity, r, err := s.client.IdentityAPI.CreateIdentityExecute(d)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusCreated {
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
