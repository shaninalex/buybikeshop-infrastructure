package keto

import (
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/ptr"
	"context"
	"fmt"
	"io"
	"net/http"

	ory "github.com/ory/keto-client-go"
)

type KetoWriter struct {
	ory.APIClient
}

func ProvideKetoWriter(config *config.Config) KetoWriter {
	if config.String("keto.write") == "" {
		panic("keto.write is required")
	}
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: config.String("keto.write"),
		},
	}
	readClient := ory.NewAPIClient(configuration)
	return KetoWriter{APIClient: *readClient}
}

func (s KetoWriter) Write(ctx context.Context, role, object, relation, subjectId string) (*ory.Relationship, error) {
	r := s.APIClient.RelationshipApi.CreateRelationship(ctx).CreateRelationshipBody(ory.CreateRelationshipBody{
		Namespace: ptr.Ptr(role),
		Object:    ptr.Ptr(object),
		Relation:  ptr.Ptr(relation),
		SubjectId: ptr.Ptr(subjectId),
	})
	rel, resp, err := s.APIClient.RelationshipApi.CreateRelationshipExecute(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unable to create relationship: %s", string(b))

	}
	return rel, nil
}
