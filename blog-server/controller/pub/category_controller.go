package pub

import (
	"blog-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryController 分类控制器
type CategoryController struct {
	categoryService *service.CategoryService
}

// NewCategoryController 创建分类控制器
func NewCategoryController(categoryService *service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// List 分类列表
func (c *CategoryController) List(ctx *gin.Context) {
	categories, err := c.categoryService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get categories"})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}
