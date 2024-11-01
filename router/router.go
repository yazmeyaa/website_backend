package router

import (
	"net/http"
	"yazmeyaa_projects/config"
	"yazmeyaa_projects/controller"
	"yazmeyaa_projects/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(projectsController *controller.ProjectsController, authController *controller.AuthController, staticFilesController *controller.StaticFilesController, schemaController *controller.SchemasController) *gin.Engine {
	router := gin.Default()
	appConfig := config.NewAppConfig()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"POST", "DELETE", "GET", "PUT", "PATCH"}
	corsConfig.AllowHeaders = []string{"authorization", "x-token", "X-Token", "Authorization", "content-type"}

	router.Use(cors.New(corsConfig))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Projects API")
	})

	baseRouter := router.Group("/api")

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/login", authController.Login)
	authRouter.POST("/register", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), authController.Register)
	authRouter.POST("/validate", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), authController.ValidateJWT)

	projectsRouter := baseRouter.Group("/projects")
	projectsRouter.GET("", projectsController.GetAll)
	projectsRouter.GET("/:id", projectsController.GetById)
	projectsRouter.POST("", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), projectsController.Create)
	projectsRouter.DELETE("/:id", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), projectsController.Delete)
	projectsRouter.PATCH("/:id", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), projectsController.Update)

	staticFilesRouter := baseRouter.Group("/static")
	staticFilesRouter.GET("/:id", staticFilesController.GetFile)
	staticFilesRouter.POST("/", staticFilesController.UploadFile)

	schemaRouter := baseRouter.Group("/schema")
	schemaRouter.GET("/:name", schemaController.GetSchemaByName)

	return router
}
