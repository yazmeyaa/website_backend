package router

import (
	"net/http"
	"yazmeyaa_projects/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(projectsController *controller.ProjectsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	projectsRouter := baseRouter.Group("/projects")
	projectsRouter.POST("", projectsController.Create)

	return router
}
