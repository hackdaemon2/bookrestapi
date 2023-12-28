package pagination

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi/app/models"
	"restapi/app/services"
	"strconv"
)

type Pagination struct{}

// Default Create a new pagination middleware with default values
func (pagination *Pagination) Default() gin.HandlerFunc {
	return pagination.init(&models.ResourcePaginationConfigType{
		PageText:    models.DefaultPageText,
		SizeText:    models.DefaultSizeText,
		Page:        models.DefaultPage,
		PageSize:    models.DefaultPageSize,
		MinPageSize: models.DefaultMinPageSize,
		MaxPageSize: models.DefaultMaxPageSize,
	})
}

// init a new pagination middleware with custom values
func (pagination *Pagination) init(resourcePaginationType *models.ResourcePaginationConfigType) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extract the page from the query string and convert it to an integer
		pageStr := context.DefaultQuery(resourcePaginationType.PageText, resourcePaginationType.Page)
		page, pageError := strconv.Atoi(pageStr)

		var errorResponse models.BaseErrorType
		var errMessage string

		if pageError != nil {
			errMessage = "page number must be an integer"
			errorResponse = services.FormulateErrorResponse(errMessage, models.ValidationError, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Validate for positive page number
		if page < 1 {
			errMessage = "page number must be greater than zero"
			errorResponse = services.FormulateErrorResponse(errMessage, models.ValidationError, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Extract the size from the query string and convert it to an integer
		sizeStr := context.DefaultQuery(resourcePaginationType.SizeText, resourcePaginationType.PageSize)
		size, sizeError := strconv.Atoi(sizeStr)

		if sizeError != nil {
			errMessage = "page size must be positive"
			errorResponse = services.FormulateErrorResponse(errMessage, models.ValidationError, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Validate for min and max page size
		if size < resourcePaginationType.MinPageSize || size > resourcePaginationType.MaxPageSize {
			minSize := strconv.Itoa(resourcePaginationType.MinPageSize)
			maxSize := strconv.Itoa(resourcePaginationType.MaxPageSize)
			errorMessage := fmt.Sprintf("page size must be between %s and %s", minSize, maxSize)
			errorResponse = services.FormulateErrorResponse(errorMessage, models.ValidationError, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Set the page and size in the gin context
		context.Set(resourcePaginationType.PageText, page)
		context.Set(resourcePaginationType.SizeText, size)

		context.Next()
	}
}
