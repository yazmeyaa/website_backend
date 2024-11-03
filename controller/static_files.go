package controller

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"yazmeyaa_projects/data/response"
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
		"path": savedFile.FileName,
	})
}

func (ic *StaticFilesController) GetFile(ctx *gin.Context) {
	fileName := ctx.Param("fileName")

	file, err := ic.staticFileService.GetByFileName(fileName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	uploadDir := "uploads"
	filePath := filepath.Join(uploadDir, file.FileName)

	ctx.File(filePath)
}

func (ic *StaticFilesController) GetAll(ctx *gin.Context) {
	files, err := ic.staticFileService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	data := make([]response.GetFilesResponse, len(files))

	for idx, val := range files {
		data[idx] = response.GetFilesResponse{
			ID:       val.ID,
			FileName: val.FileName,
		}
	}

	resp := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, resp)
}
