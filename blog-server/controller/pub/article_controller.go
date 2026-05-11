package pub

import (
	"blog-server/model"
	"blog-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticleController 文章控制器
type ArticleController struct {
	articleService *service.ArticleService
}

// NewArticleController 创建文章控制器
func NewArticleController(articleService *service.ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

// List 文章列表
func (c *ArticleController) List(ctx *gin.Context) {
	var req model.ArticleListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// 公开接口只显示已发布的文章
	status := 1
	req.Status = &status

	resp, err := c.articleService.List(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetByID 根据ID获取文章
func (c *ArticleController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	article, err := c.articleService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 公开接口只显示已发布的文章
	if article.Status != 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 增加阅读量
	go c.articleService.IncrementViewCount(id)

	ctx.JSON(http.StatusOK, article)
}

// GetBySlug 根据slug获取文章
func (c *ArticleController) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Slug is required"})
		return
	}

	article, err := c.articleService.GetBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 公开接口只显示已发布的文章
	if article.Status != 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 增加阅读量
	go c.articleService.IncrementViewCount(article.ID)

	ctx.JSON(http.StatusOK, article)
}
