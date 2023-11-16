package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
)

type Validation struct {
	errorMessage string
}

var once sync.Once
var validationInstance *Validation

var validate = GetValidationInstance()

func GetValidationInstance() *Validation {
	if validationInstance == nil {
		once.Do(func() { validationInstance = new(Validation) })
	}
	return validationInstance
}

func (validation *Validation) ValidateStruct(payload any) *map[string]string {
	err := validator.New().Struct(payload)
	validationErrors := make(map[string]string)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := validation.extractFieldName(err.Namespace(), payload)
			validationErrors[fieldName] = fmt.Sprintf("JSON field '%s' is invalid: %s[%s]", fieldName, err.Tag(), err.Param())
		}
	}
	return &validationErrors
}

func (validation *Validation) extractFieldName(namespace string, payload any) string {
	parts := strings.Split(namespace, ".")
	field := parts[len(parts)-1]
	return validation.getJSONFieldName(field, reflect.TypeOf(payload))
}

func (validation *Validation) getJSONFieldName(fieldName string, structType reflect.Type) string {
	field, found := structType.FieldByName(fieldName)
	if !found {
		return ""
	}
	jsonTag := field.Tag.Get("json")
	return jsonTag
}
