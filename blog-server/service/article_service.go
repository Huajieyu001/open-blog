package service

import (
	"blog-server/model"
	"blog-server/repository"
	"bytes"
	"errors"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

// ArticleService 文章服务
type ArticleService struct {
	articleRepo  *repository.ArticleRepository
	categoryRepo *repository.CategoryRepository
	tagRepo      *repository.TagRepository
}

// NewArticleService 创建文章服务
func NewArticleService(
	articleRepo *repository.ArticleRepository,
	categoryRepo *repository.CategoryRepository,
	tagRepo *repository.TagRepository,
) *ArticleService {
	return &ArticleService{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

// GetByID 根据ID获取文章
func (s *ArticleService) GetByID(id int64) (*model.Article, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("article not found")
	}

	// 获取分类信息
	if article.CategoryID != nil {
		category, err := s.categoryRepo.FindByID(*article.CategoryID)
		if err == nil && category != nil {
			article.Category = category
		}
	}

	// 获取标签信息
	tags, err := s.articleRepo.GetTags(article.ID)
	if err == nil {
		article.Tags = tags
	}

	return article, nil
}

// GetBySlug 根据slug获取文章
func (s *ArticleService) GetBySlug(slug string) (*model.Article, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("article not found")
	}

	// 获取分类信息
	if article.CategoryID != nil {
		category, err := s.categoryRepo.FindByID(*article.CategoryID)
		if err == nil && category != nil {
			article.Category = category
		}
	}

	// 获取标签信息
	tags, err := s.articleRepo.GetTags(article.ID)
	if err == nil {
		article.Tags = tags
	}

	return article, nil
}

// List 文章列表
func (s *ArticleService) List(req *model.ArticleListRequest) (*model.ArticleListResponse, error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	resp, err := s.articleRepo.List(req)
	if err != nil {
		return nil, err
	}

	// 获取每篇文章的标签
	for i := range resp.List {
		tags, err := s.articleRepo.GetTags(resp.List[i].ID)
		if err == nil {
			resp.List[i].Tags = tags
		}
	}

	return resp, nil
}

// Create 创建文章
func (s *ArticleService) Create(authorID int64, req *model.CreateArticleRequest) (*model.Article, error) {
	// 验证分类是否存在
	if req.CategoryID != nil {
		category, err := s.categoryRepo.FindByID(*req.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, errors.New("category not found")
		}
	}

	// 渲染Markdown为HTML
	contentHTML := renderMarkdown(req.ContentMD)

	article := &model.Article{
		Title:       req.Title,
		Slug:        req.Slug,
		Summary:     req.Summary,
		ContentMD:   req.ContentMD,
		ContentHTML: contentHTML,
		CoverImage:  req.CoverImage,
		Status:      req.Status,
		IsTop:       req.IsTop,
		CategoryID:  req.CategoryID,
		AuthorID:    &authorID,
	}

	// 如果状态为发布，设置发布时间
	if article.Status == 1 {
		now := time.Now()
		article.PublishTime = &now
	}

	// 创建文章
	if err := s.articleRepo.Create(article); err != nil {
		return nil, err
	}

	// 设置标签
	if len(req.TagIDs) > 0 {
		if err := s.articleRepo.SetTags(article.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	// 获取完整文章信息
	return s.GetByID(article.ID)
}

// Update 更新文章
func (s *ArticleService) Update(id int64, req *model.UpdateArticleRequest) (*model.Article, error) {
	// 获取原文章
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("article not found")
	}

	// 验证分类是否存在
	if req.CategoryID != nil {
		category, err := s.categoryRepo.FindByID(*req.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, errors.New("category not found")
		}
	}

	// 更新字段
	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Slug != "" {
		article.Slug = req.Slug
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	if req.ContentMD != "" {
		article.ContentMD = req.ContentMD
		article.ContentHTML = renderMarkdown(req.ContentMD)
	}
	if req.CoverImage != "" {
		article.CoverImage = req.CoverImage
	}
	if req.Status != nil {
		// 如果从草稿变为发布，设置发布时间
		if article.Status == 0 && *req.Status == 1 {
			now := time.Now()
			article.PublishTime = &now
		}
		article.Status = *req.Status
	}
	if req.IsTop != nil {
		article.IsTop = *req.IsTop
	}
	if req.CategoryID != nil {
		article.CategoryID = req.CategoryID
	}

	// 更新文章
	if err := s.articleRepo.Update(article); err != nil {
		return nil, err
	}

	// 更新标签
	if req.TagIDs != nil {
		if err := s.articleRepo.SetTags(article.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	// 获取完整文章信息
	return s.GetByID(article.ID)
}

// Delete 删除文章
func (s *ArticleService) Delete(id int64) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return err
	}
	if article == nil {
		return errors.New("article not found")
	}
	return s.articleRepo.Delete(id)
}

// IncrementViewCount 增加阅读量
func (s *ArticleService) IncrementViewCount(id int64) error {
	return s.articleRepo.IncrementViewCount(id)
}

// renderMarkdown 渲染Markdown为HTML
func renderMarkdown(content string) string {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return content
	}
	return buf.String()
}
