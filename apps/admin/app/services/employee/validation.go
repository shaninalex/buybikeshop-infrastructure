package employee

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func (s serviceImpl) Validate(ctx context.Context, data EmployeeCreate) error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(data)
}
