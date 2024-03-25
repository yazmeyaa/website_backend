package router

import (
	"net/http"
	"yazmeyaa_projects/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(projectsController *controller.ProjectsController) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	projectsRouter := baseRouter.Group("/projects")
	projectsRouter.GET("", projectsController.GetAll)
	projectsRouter.GET("/:id", projectsController.GetById)
	projectsRouter.POST("", projectsController.Create)
	projectsRouter.DELETE("/:id", projectsController.Delete)
	projectsRouter.PATCH("/:id", projectsController.Update)

	return router
}
