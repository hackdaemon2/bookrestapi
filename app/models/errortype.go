package models

type (
	BaseErrorResponder interface {
		SetBaseError(baseError BaseErrorType)
	}

	BaseErrorType struct {
		Error       bool              `json:"error"`
		ErrorMsg    *string           `json:"error_msg,omitempty"`
		ErrorType   string            `json:"error_type,omitempty"`
		ErrorDetail map[string]string `json:"validation_errors,omitempty"`
	}
)
