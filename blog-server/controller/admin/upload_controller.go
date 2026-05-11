package admin

import (
	"blog-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadController 上传控制器
type UploadController struct {
	uploadService *service.UploadService
}

// NewUploadController 创建上传控制器
func NewUploadController(uploadService *service.UploadService) *UploadController {
	return &UploadController{uploadService: uploadService}
}

// Upload 上传文件
func (c *UploadController) Upload(ctx *gin.Context) {
	// 获取上传的文件
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// 上传到OSS
	resp, err := c.uploadService.Upload(file, header.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
