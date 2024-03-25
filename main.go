package main

import (
	"fmt"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	Addr := fmt.Sprintf(":%s", port)

	server := &http.Server{
		Addr:    Addr,
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
