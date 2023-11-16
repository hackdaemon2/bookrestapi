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
	return pagination.init(&models.ResourcePaginationConfigDTO{
		PageText:    models.DefaultPageText,
		SizeText:    models.DefaultSizeText,
		Page:        models.DefaultPage,
		PageSize:    models.DefaultPageSize,
		MinPageSize: models.DefaultMinPageSize,
		MaxPageSize: models.DefaultMaxPageSize,
	})
}

// init a new pagination middleware with custom values
func (pagination *Pagination) init(resourcePaginationDTO *models.ResourcePaginationConfigDTO) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extract the page from the query string and convert it to an integer
		pageStr := context.DefaultQuery(resourcePaginationDTO.PageText, resourcePaginationDTO.Page)
		page, pageError := strconv.Atoi(pageStr)

		var errorResponse models.BaseErrorDTO
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
		sizeStr := context.DefaultQuery(resourcePaginationDTO.SizeText, resourcePaginationDTO.PageSize)
		size, sizeError := strconv.Atoi(sizeStr)

		if sizeError != nil {
			errMessage = "page size must be positive"
			errorResponse = services.FormulateErrorResponse(errMessage, models.ValidationError, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Validate for min and max page size
		if size < resourcePaginationDTO.MinPageSize || size > resourcePaginationDTO.MaxPageSize {
			minSize := strconv.Itoa(resourcePaginationDTO.MinPageSize)
			maxSize := strconv.Itoa(resourcePaginationDTO.MaxPageSize)
			errorMessage := fmt.Sprintf("page size must be between %s and %s", minSize, maxSize)
			errorResponse = services.FormulateErrorResponse(errorMessage, models.ValidationError, nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Set the page and size in the gin context
		context.Set(resourcePaginationDTO.PageText, page)
		context.Set(resourcePaginationDTO.SizeText, size)

		context.Next()
	}
}
