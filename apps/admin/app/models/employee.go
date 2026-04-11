package models

import (
	ory "github.com/ory/kratos-client-go"
)

type Employee struct {
	Identity ory.Identity `json:"identity"`
}

type EmployeeCreate struct {
	Identity ory.CreateIdentityBody `json:"identity"`
}
