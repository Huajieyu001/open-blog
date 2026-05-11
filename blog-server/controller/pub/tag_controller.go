package pub

import (
	"blog-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TagController 标签控制器
type TagController struct {
	tagService *service.TagService
}

// NewTagController 创建标签控制器
func NewTagController(tagService *service.TagService) *TagController {
	return &TagController{tagService: tagService}
}

// List 标签列表
func (c *TagController) List(ctx *gin.Context) {
	tags, err := c.tagService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tags"})
		return
	}

	ctx.JSON(http.StatusOK, tags)
}
