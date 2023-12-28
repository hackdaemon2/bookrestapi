package services

import (
	"encoding/json"
	"fmt"
	"log"
	"restapi/app/models"
)

func FormulateErrorResponse(errorMessage string, errorType string, errorMap map[string]string) models.BaseErrorType {
	response := models.BaseErrorType{
		ErrorMsg:  &errorMessage,
		Error:     true,
		ErrorType: errorType,
	}

	if len(errorMap) != 0 {
		response.ErrorDetail = errorMap
	}

	jsonData, _ := json.Marshal(response)
	log.Println(fmt.Sprintf(models.ResponseLogMessage, string(jsonData)))
	return response
}
