package main

import (
	"fmt"
	"net/http"
	"yazmeyaa_projects/config"
	"yazmeyaa_projects/controller"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/model"
	"yazmeyaa_projects/repository"
	"yazmeyaa_projects/router"
	"yazmeyaa_projects/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()
	appConfig := config.NewAppConfig()

	db.AutoMigrate(&model.Project{}, &model.User{}, &model.StaticFile{})

	projectsRepository := repository.NewProjectsRepositoryImpl(db)
	projectsService := service.NewProjectsServiceImpl(projectsRepository, validate)
	projectsController := controller.NewProjectsController(projectsService)

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository, validate)
	authController := controller.NewAuthController(authService)

	staticFileRepository := repository.NewStaticFileRepository(db)
	staticFileService := service.NewStaticFileService(staticFileRepository)
	staticFilesController := controller.NewStaticFilesController(staticFileService)

	schemaController := controller.NewSchemasController()

	routes := router.NewRouter(projectsController, authController, staticFilesController, schemaController)

	Addr := fmt.Sprintf("%s:%s", appConfig.Server.Host, appConfig.Server.Port)

	server := &http.Server{
		Addr:    Addr,
		Handler: routes,
	}

	fmt.Printf("Server started: %s:%s", appConfig.Server.Host, appConfig.Server.Port)

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
