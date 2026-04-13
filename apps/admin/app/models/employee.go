package models

import (
	"encoding/json"

	ory "github.com/ory/kratos-client-go"
)

type Employee struct {
	Identity ory.Identity `json:"identity"`
}

func (e *Employee) Id() string {
	return e.Identity.GetId()
}

func (e *Employee) GetIdentity() ory.Identity {
	return e.Identity
}

type EmployeeTraits struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Dob   string `json:"dob"`
	Photo string `json:"photo"`
}

func (e *Employee) Traits() (EmployeeTraits, error) {
	var traits EmployeeTraits

	if e.Identity.Traits == nil {
		return traits, nil
	}

	data, err := json.Marshal(e.Identity.Traits)
	if err != nil {
		return traits, err
	}

	err = json.Unmarshal(data, &traits)
	return traits, err
}

type EmployeeCreateModel struct {
	Traits   EmployeeTraits `json:"traits"`
	Password string         `json:"password"`
	Role     string         `json:"role"`
}
