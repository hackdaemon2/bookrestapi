package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"restapi/app/models"
	"restapi/app/repository"
	"strconv"
	"time"
)

var errorResponse models.BaseErrorDTO

// BookService provides business logic for Book operations
type BookService struct {
	bookRepository repository.IBookRepository
}

// NewBookService creates a new BookService
func NewBookService(bookRepository repository.IBookRepository) *BookService {
	return &BookService{bookRepository}
}

// FindAll Get all Books retrieves all books from the database
func (bookService *BookService) FindAll(context *gin.Context) {
	page := context.GetInt("page")
	size := context.GetInt("size")
	data, err := bookService.bookRepository.FindAll(page, size)

	if err != nil {
		errorResponse = FormulateErrorResponse(err.Error(), models.DataError, nil)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var rowCount int64

	if len(data) > 0 {
		rowCount, _ = bookService.bookRepository.CountAll()
	}

	response := models.BookPaginationResponseDTO{
		Page:     &page,
		Data:     data,
		Size:     &size,
		RowCount: &rowCount,
		BaseErrorDTO: models.BaseErrorDTO{
			Error: false,
		},
	}

	jsonData, _ := json.Marshal(response)
	log.Println(fmt.Sprintf(models.ResponseLogMessage, string(jsonData)))
	context.JSON(http.StatusOK, response)
}

// FindOne retrieves a book by ID from the database
func (bookService *BookService) FindOne(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	log.Printf("Get Book By ID => %d", id)

	book, err := bookService.bookRepository.FindOne(int64(id))

	if err != nil {
		log.Println("Unable to retrieve book")
		errorResponse = FormulateErrorResponse(err.Error(), models.DataError, nil)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := models.BookResponseDTO{}

	if book.ID == 0 {
		log.Println("Book does not exist")
		errMsg := "Book does not exist"
		response.BaseErrorDTO.Error = true
		response.BaseErrorDTO.ErrorMsg = &errMsg
		response.BaseErrorDTO.ErrorType = models.ResourceNotFound
		jsonData, _ := json.Marshal(response)
		log.Println(fmt.Sprintf(models.ResponseLogMessage, string(jsonData)))
		context.JSON(http.StatusNotFound, response)
		return
	}

	response.BaseErrorDTO.Error = false
	response.Data = book

	jsonData, _ := json.Marshal(response)
	log.Println(fmt.Sprintf(models.ResponseLogMessage, string(jsonData)))
	context.JSON(http.StatusOK, response)
}

// Create creates a new book in the database
func (bookService *BookService) Create(context *gin.Context) {
	var bookRequest models.BookRequestDTO

	if err := context.ShouldBindJSON(&bookRequest); err != nil {
		errorResponse = FormulateErrorResponse("Unable to bind request", models.BindingError, nil)
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	jsonData, _ := json.Marshal(bookRequest)
	log.Println(fmt.Sprintf(models.RequestLogMessage, string(jsonData)))

	// Validate the request payload
	if err := validate.ValidateStruct(bookRequest); len(*err) != 0 {
		errorResponse = FormulateErrorResponse("Request validation errors", models.ValidationError, *err)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	book := models.Book{}
	book.Title = bookRequest.Title
	book.Author = bookRequest.Author
	book.CreatedAt = models.CustomDateTimeMySQL{Time: time.Now()}
	book.DeletedAt = models.CustomDateTimeMySQL{}
	book, err := bookService.bookRepository.Create(book)

	if err != nil {
		errorMessage := "Unable to create book on the database"
		log.Println(errorMessage)
		errorResponse = FormulateErrorResponse(errorMessage, models.DataError, nil)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if book.ID == 0 {
		errorResponse = FormulateErrorResponse(models.InsertErrorMessage, models.DataError, nil)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := models.BookResponseDTO{Data: book, BaseErrorDTO: models.BaseErrorDTO{Error: false}}

	jsonData, _ = json.Marshal(response)
	log.Println(fmt.Sprintf(models.ResponseLogMessage, string(jsonData)))
	context.JSON(http.StatusCreated, response)
}

// Update updates a book in the database
func (bookService *BookService) Update(context *gin.Context) {
	var bookRequest models.BookRequestDTO

	if err := context.ShouldBindJSON(&bookRequest); err != nil {
		errorResponse = FormulateErrorResponse("Unable to bind request", models.BindingError, nil)
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	jsonData, _ := json.Marshal(bookRequest)
	log.Println(fmt.Sprintf(models.RequestLogMessage, string(jsonData)))

	// Validate the request payload
	if err := validate.ValidateStruct(bookRequest); len(*err) != 0 {
		errorResponse = FormulateErrorResponse("Request validation errors", models.ValidationError, *err)
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	id, _ := strconv.Atoi(context.Param("id"))

	book, err := bookService.bookRepository.FindOne(int64(id))

	if err != nil {
		errorResponse = FormulateErrorResponse(models.UpdateErrorMessage, models.DataError, nil)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if book.ID == 0 {
		errorResponse = FormulateErrorResponse(models.ResourceNotFoundErrorMessage, models.ResourceNotFound, nil)
		context.JSON(http.StatusNotFound, errorResponse)
		return
	}

	book.Author = bookRequest.Author
	book.Title = bookRequest.Title
	book.UpdatedAt = models.CustomDateTimeMySQL{Time: time.Now()}
	updatedBook, err := bookService.bookRepository.Update(book)

	if err != nil {
		errorResponse = FormulateErrorResponse(models.UpdateErrorMessage, models.DataError, nil)
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := models.BookResponseDTO{Data: updatedBook, BaseErrorDTO: models.BaseErrorDTO{Error: false}}
	jsonData, _ = json.Marshal(response)
	log.Println(fmt.Sprintf(models.ResponseLogMessage, string(jsonData)))
	context.JSON(http.StatusOK, response)
}

// Delete deletes a book from the database by ID
func (bookService *BookService) Delete(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	book, err := bookService.bookRepository.FindOne(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	if book.ID == 0 {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	result, _ := bookService.bookRepository.Delete(&book)

	if result == 0 {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
