package router

import (
	"fmt"
	"net/http"
	"time"
	"yazmeyaa_projects/config"
	"yazmeyaa_projects/controller"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(projectsController *controller.ProjectsController, authController *controller.AuthController) *gin.Engine {
	router := gin.Default()
	appConfig := config.NewAppConfig()

	token, err := helper.CreateAccessToken(appConfig.JWT.Secret, int(time.Hour.Nanoseconds())*168)
	if err != nil {
		helper.ErrorPanic(err)
	}

	fmt.Printf("jwt is: %s\n", token)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, "POST", "DELETE", "GET", "PUT")

	router.Use(cors.New(corsConfig))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Projects API")
	})

	baseRouter := router.Group("/api")

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/login", authController.Login)
	authRouter.POST("/register", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), authController.Register)

	projectsRouter := baseRouter.Group("/projects")
	projectsRouter.GET("", projectsController.GetAll)
	projectsRouter.GET("/:id", projectsController.GetById)
	projectsRouter.POST("", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), projectsController.Create)
	projectsRouter.DELETE("/:id", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), projectsController.Delete)
	projectsRouter.PATCH("/:id", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), projectsController.Update)

	return router
}
