package router

import (
	"fmt"
	"net/http"
	"time"
	"yazmeyaa_projects/controller"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(projectsController *controller.ProjectsController, authController *controller.AuthController) *gin.Engine {
	router := gin.Default()
	const secret string = "tTkOJFQu4S"

	token, err := helper.CreateAccessToken(secret, int(time.Hour.Nanoseconds())*168)
	if err != nil {
		helper.ErrorPanic(err)
	}

	fmt.Printf("jwt is: %s\n", token)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Projects API")
	})

	baseRouter := router.Group("/api")

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/login", authController.Login)
	authRouter.POST("/register", middlewares.AuthJWTMiddleware(secret), authController.Register)

	projectsRouter := baseRouter.Group("/projects")
	projectsRouter.Use(middlewares.AuthJWTMiddleware(secret))
	projectsRouter.GET("", projectsController.GetAll)
	projectsRouter.GET("/:id", projectsController.GetById)
	projectsRouter.POST("", projectsController.Create)
	projectsRouter.DELETE("/:id", projectsController.Delete)
	projectsRouter.PATCH("/:id", projectsController.Update)

	return router
}
