package controller

import (
	"github.com/gin-gonic/gin"
	"restapi/app/services"
)

type BookController struct {
	bookService services.IBookService
}

func NewBookController(bookService services.IBookService) *BookController {
	return &BookController{bookService}
}

func (bookController *BookController) All(context *gin.Context) {
	bookController.bookService.FindAll(context)
}

func (bookController *BookController) Create(context *gin.Context) {
	bookController.bookService.Create(context)
}

func (bookController *BookController) Update(context *gin.Context) {
	bookController.bookService.Update(context)
}

func (bookController *BookController) Read(context *gin.Context) {
	bookController.bookService.FindOne(context)
}

func (bookController *BookController) Delete(context *gin.Context) {
	bookController.bookService.Delete(context)
}
