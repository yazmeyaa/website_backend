package main

import (
	"net/http"
	"yazmeyaa_projects/config"
	"yazmeyaa_projects/controller"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/model"
	"yazmeyaa_projects/repository"
	"yazmeyaa_projects/router"
	"yazmeyaa_projects/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("projects").AutoMigrate(&model.Project{})

	projectsRepository := repository.NewProjectsRepositoryImpl(db)
	projectsService := service.NewProjectsServiceImpl(projectsRepository, validate)
	projectsController := controller.NewProjectsController(projectsService)

	routes := router.NewRouter(projectsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
