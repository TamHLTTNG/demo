package router

import (
	"go-auth-poc/pkg/controller"

	_ "go-auth-poc/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(c controller.AppController) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	// r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/swagger"})))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth
	auth := r.Group("")
	{
		auth.POST("/get-jwt", func(context *gin.Context) { c.Auth.POSTTokenVerify(context) })
	}

	// Set default binding to json
	gin.SetMode(gin.ReleaseMode)

	return r
}
