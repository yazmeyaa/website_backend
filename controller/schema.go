package controller

import (
	"encoding/json"
	"net/http"
	"yazmeyaa_projects/schemas"

	"github.com/gin-gonic/gin"
)

type SchemasController struct {
	projectSchema    *schemas.ProjectSchema
	staticFileSchema *schemas.StaticFileSchema
}

func NewSchemasController() *SchemasController {
	projectSchema := schemas.NewProjectSchema()
	staticFileSchema := schemas.NewStaticFileSchema()
	return &SchemasController{
		projectSchema:    projectSchema,
		staticFileSchema: staticFileSchema,
	}
}

func (c *SchemasController) GetSchemaByName(ctx *gin.Context) {
	name := ctx.Param("name")
	encoder := json.NewEncoder(ctx.Writer)
	ctx.Header("Content-Type", "application/json")

	switch name {
	case "project":
		if err := encoder.Encode(c.projectSchema); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Cannot encode schema.",
			})
		}
		return
	case "staticFile":
		if err := encoder.Encode(c.staticFileSchema); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Cannot encode schema.",
			})
		}
		return
	default:
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Cannot get schema",
		})
		return
	}
}
