package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restapi/app/models"
	"strings"
	"time"
)

var middlewareInstance *Middleware
var errMsg string

type Middleware struct{}

func GetInstanceMiddleware() *Middleware {
	if middlewareInstance == nil {
		Once.Do(func() { middlewareInstance = new(Middleware) })
	}
	return middlewareInstance
}

func (middleware *Middleware) RequestAuthorization(config models.IAppConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader(models.AuthorizationHeader)

		if authorization == "" {
			errMsg = "Authorization header must be provided"
			context.JSON(401, models.BaseErrorType{Error: true, ErrorMsg: &errMsg})
			context.Abort()
			return
		}

		bearer := authorization[0:7]
		hasPrefixLowerCase := strings.HasPrefix(bearer, models.BearerPrefixLowerCase)
		hasPrefixUpperCase := strings.HasPrefix(bearer, models.BearerPrefixTitleCase)

		if !hasPrefixLowerCase && !hasPrefixUpperCase {
			errMsg = "Authorization header 'Bearer' must be provided"
			context.JSON(401, models.BaseErrorType{Error: true, ErrorMsg: &errMsg})
			context.Abort()
			return
		}

		token, _ := strings.CutPrefix(authorization, bearer)

		if config.BearerToken() != token {
			errMsg = "Authorization token is invalid"
			context.JSON(401, models.BaseErrorType{Error: true, ErrorMsg: &errMsg})
			context.Abort()
			return
		}

		context.Next()
	}
}

func (middleware *Middleware) RequestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		context.Next()
		latency := time.Since(now)
		fmt.Printf("%s %s %s %s \n",
			context.Request.Method,
			context.Request.RequestURI,
			context.Request.Proto,
			latency.String())
	}
}

func (middleware *Middleware) ResponseLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		context.Next()
		fmt.Printf("%d %s %s\n",
			context.Writer.Status(),
			context.Request.Method,
			context.Request.RequestURI)
	}
}
