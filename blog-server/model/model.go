package model

import "time"

// User 管理员用户
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Category 分类
type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Tag 标签
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Article 文章
type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Summary     string    `json:"summary"`
	ContentMD   string    `json:"content_md"`
	ContentHTML string    `json:"content_html"`
	CoverImage  string    `json:"cover_image"`
	Status      int       `json:"status"` // 0:草稿 1:发布
	IsTop       int       `json:"is_top"` // 0:否 1:置顶
	ViewCount   int       `json:"view_count"`
	CategoryID  *int64    `json:"category_id"`
	AuthorID    *int64    `json:"author_id"`
	PublishTime *time.Time `json:"publish_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Category    *Category `json:"category,omitempty"`
	Author      *User     `json:"author,omitempty"`
	Tags        []Tag     `json:"tags,omitempty"`
}

// ArticleListRequest 文章列表请求
type ArticleListRequest struct {
	Page       int    `form:"page,default=1"`
	PageSize   int    `form:"page_size,default=10"`
	Status     *int   `form:"status"`
	CategoryID *int64 `form:"category_id"`
	TagID      *int64 `form:"tag_id"`
	Keyword    string `form:"keyword"`
}

// ArticleListResponse 文章列表响应
type ArticleListResponse struct {
	Total    int64     `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	List     []Article `json:"list"`
}

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title      string  `json:"title" binding:"required"`
	Slug       string  `json:"slug"`
	Summary    string  `json:"summary"`
	ContentMD  string  `json:"content_md" binding:"required"`
	CoverImage string  `json:"cover_image"`
	Status     int     `json:"status"`
	IsTop      int     `json:"is_top"`
	CategoryID *int64  `json:"category_id"`
	TagIDs     []int64 `json:"tag_ids"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Title      string  `json:"title"`
	Slug       string  `json:"slug"`
	Summary    string  `json:"summary"`
	ContentMD  string  `json:"content_md"`
	CoverImage string  `json:"cover_image"`
	Status     *int    `json:"status"`
	IsTop      *int    `json:"is_top"`
	CategoryID *int64  `json:"category_id"`
	TagIDs     []int64 `json:"tag_ids"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateCategoryRequest 创建分类请求
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

// UpdateCategoryRequest 更新分类请求
type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	SortOrder   *int   `json:"sort_order"`
}

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// UploadResponse 上传响应
type UploadResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
}
