package router

import (
	"net/http"
	"yazmeyaa_projects/config"
	"yazmeyaa_projects/controller"
	"yazmeyaa_projects/middlewares"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(corsService service.CorsService, projectsController *controller.ProjectsController, authController *controller.AuthController, staticFilesController *controller.StaticFilesController, schemaController *controller.SchemasController, corsController *controller.CorsController) *gin.Engine {
	router := gin.Default()
	appConfig := config.NewAppConfig()

	router.Use(middlewares.DynamicCorsMiddleware(corsService))

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
	staticFilesRouter.GET("/:fileName", staticFilesController.GetFile)
	staticFilesRouter.GET("/files", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), staticFilesController.GetAll)
	staticFilesRouter.POST("/", middlewares.AuthJWTMiddleware(appConfig.JWT.Secret), staticFilesController.UploadFile)

	appSettingsRouter := baseRouter.Group("/settings")
	appSettingsRouter.Use(middlewares.AuthJWTMiddleware(appConfig.JWT.Secret))

	corsRouter := appSettingsRouter.Group("/cors")
	corsRouter.GET("/", corsController.GetAllOrigins)
	corsRouter.POST("/", corsController.AddOrigin)
	corsRouter.PATCH("/:origin/disable", corsController.DisableOrigin)
	corsRouter.PATCH("/:origin/enable", corsController.EnableOrigin)
	corsRouter.DELETE("/:origin", corsController.RemoveOrigin)

	schemaRouter := baseRouter.Group("/schema")
	schemaRouter.GET("/:name", schemaController.GetSchemaByName)

	return router
}
