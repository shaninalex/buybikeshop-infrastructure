package mock

import (
	"buybikeshop/libs/go/kratos"
	"context"
	"encoding/json"
	"io"
	"maps"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type KratosApiClient struct {
	identities map[uuid.UUID]ory.Identity
	mu         sync.RWMutex
}

var _ kratos.ApiClient = (*KratosApiClient)(nil)

func ProvideKratosApiClient() kratos.ApiClient {
	f, err := os.Open("libs/go/mock/data/identities.json")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	identities := make(map[uuid.UUID]ory.Identity)
	if err = json.Unmarshal(b, &identities); err != nil {
		panic(err)
	}

	return &KratosApiClient{
		identities: identities,
	}
}

func (m *KratosApiClient) CreateIdentity(ctx context.Context, data ory.CreateIdentityBody) (*ory.Identity, error) {
	created_at := time.Now()
	id := uuid.New()
	email, ok := data.Traits["email"].(string)
	if !ok {
		return nil, kratos.IdentityApiErrorMissingEmail
	}

	identity := &ory.Identity{
		Id:        id.String(),
		Traits:    data.Traits,
		CreatedAt: &created_at,
	}

	cred := ory.NewIdentityCredentials()
	cred.SetType("password")
	cred.SetIdentifiers([]string{email})
	cred.SetVersion(0)

	identity.SetCredentials(map[string]ory.IdentityCredentials{"password": *cred})

	m.mu.Lock()
	m.identities[id] = *identity
	defer m.mu.Unlock()

	return identity, nil
}

func (m *KratosApiClient) ListIdentities(ctx context.Context) ([]ory.Identity, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return slices.Collect(maps.Values(m.identities)), nil
}

func (m *KratosApiClient) DeleteIdentity(ctx context.Context, id uuid.UUID) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.identities[id]; !ok {
		return false, kratos.IdentityApiErrorNotFound
	}

	delete(m.identities, id)
	return true, nil
}
