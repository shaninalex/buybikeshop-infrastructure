package models

import "github.com/google/uuid"

type Employee struct {
	EmployeeId   uuid.UUID `json:"employee_id" db:"employee_id"`
	DepartmentId uint64    `json:"department_id" db:"department_id"`
}

type Department struct {
	Id    uint64 `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}
