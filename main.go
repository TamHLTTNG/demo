package main

import (
	"fmt"
	"log"

	_ "go-auth-poc/docs"

	"go-auth-poc/pkg/config"
	"go-auth-poc/pkg/controller"
	"go-auth-poc/pkg/domain/model"
	"go-auth-poc/pkg/domain/model/auth"
	"go-auth-poc/pkg/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Auth POC開発 swagger API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for this security definition being used

func main() {
	config.ReadConfig()

	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&model.User{}, &auth.Role{}, &auth.Permission{}, &auth.RolePermission{})

	r := controller.NewRegistry(db)

	route := router.NewRouter(r.NewAppController())

	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := route.Run(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
