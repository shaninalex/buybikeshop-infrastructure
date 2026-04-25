package keto

import (
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/ptr"
	"context"
	"errors"
	"fmt"

	ory "github.com/ory/keto-client-go"
)

const (
	namespaceGroup = "Group"
	roleMember     = "member"
)

type PermissionManager interface {
	Assign(ctx context.Context, object string, subjectId *string, sub *SubjectSetLocal) error
	Delete()
}

type PermissionCheck interface {
	Check()
}

type Manager struct {
	writeClient *ory.APIClient
	readClient  *ory.APIClient
}

func ProvideManager(config *config.Config) *Manager {
	writeConfig := ory.NewConfiguration()
	writeConfig.Servers = []ory.ServerConfiguration{{URL: config.String("keto.write")}}
	writeClient := ory.NewAPIClient(writeConfig)

	readConfig := ory.NewConfiguration()
	readConfig.Servers = []ory.ServerConfiguration{{URL: config.String("keto.read")}}
	readClient := ory.NewAPIClient(readConfig)

	return &Manager{
		writeClient: writeClient,
		readClient:  readClient,
	}
}

// SubjectSetLocal - is the mirror of ory.SubjectSet to reduce types coupling in the system
type SubjectSetLocal struct {
	Namespace string `json:"namespace"`
	Object    string `json:"object"`
	Relation  string `json:"relation"`
}

func (m *Manager) Assign(ctx context.Context, object string, subjectId *string, sub *SubjectSetLocal) error {
	// NOTE: in most of the cases namespace will be "Group" because this function planing to apply only for users
	// 		 Every user is in some group
	body := ory.CreateRelationshipBody{
		Namespace: ptr.Ptr(namespaceGroup),
		Object:    &object,
		Relation:  ptr.Ptr(roleMember),
	}
	if sub != nil {
		body.SubjectSet = ory.NewSubjectSet(sub.Namespace, sub.Object, sub.Relation)
	}

	if subjectId != nil {
		body.SubjectId = subjectId
	}

	if sub == nil && subjectId == nil {
		return errors.New("subject and subjectId are both nil")
	}

	_, resp, err := m.writeClient.RelationshipApi.CreateRelationship(ctx).CreateRelationshipBody(body).Execute()
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf(resp.Status)
	}

	return nil
}

func (m *Manager) Delete() {

}

func (m *Manager) Check() {

}
