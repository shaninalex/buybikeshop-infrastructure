package transport

import (
	"encoding/json"
	"errors"
)

type ApiError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func (e ApiError) Error() string {
	return e.Message
}

func ToApiError(err error) ApiError {
	return ApiError{
		Message: err.Error(),
	}
}

// FromOryError converts an ory-generated OpenAPI error (kratos, keto, hydra)
// into an ApiError. It duck-types on the `Body() []byte` method so the same
// helper works for every ory client's GenericOpenAPIError.
func FromOryError(err error) ApiError {
	var be interface{ Body() []byte }
	if errors.As(err, &be) {
		var payload struct {
			Error struct {
				Code    int    `json:"code"`
				Status  string `json:"status"`
				Reason  string `json:"reason"`
				Message string `json:"message"`
			} `json:"error"`
		}
		if jsonErr := json.Unmarshal(be.Body(), &payload); jsonErr == nil && payload.Error.Message != "" {
			return ApiError{
				Message: payload.Error.Message,
				Reason:  payload.Error.Reason,
				Code:    payload.Error.Code,
				Status:  payload.Error.Status,
			}
		}
	}
	return ToApiError(err)
}
