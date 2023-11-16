package services

import "github.com/gin-gonic/gin"

type IBookService interface {
	FindAll(ctx *gin.Context)
	FindOne(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}
