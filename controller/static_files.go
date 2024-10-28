package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
)

type StaticFilesController struct {
	staticFileService service.StaticFileService
}

func NewStaticFilesController(staticFileService service.StaticFileService) *StaticFilesController {
	return &StaticFilesController{
		staticFileService: staticFileService,
	}
}

func (ic *StaticFilesController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Printf(">>>>ERROR: %s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	fileData, err := file.Open()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer fileData.Close()

	data, err := io.ReadAll(fileData)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	savedFile, err := ic.staticFileService.Create(file.Filename, data)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":   savedFile.ID,
		"path": savedFile.Path,
	})
}

func (ic *StaticFilesController) GetFile(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "wrong id provided"})
		return
	}

	file, err := ic.staticFileService.GetById(idInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	ctx.File(file.Path)
}
