package repository

import (
	"blog-server/model"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// ArticleRepository 文章仓储
type ArticleRepository struct {
	db *sql.DB
}

// NewArticleRepository 创建文章仓储
func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// FindByID 根据ID查找文章
func (r *ArticleRepository) FindByID(id int64) (*model.Article, error) {
	article := &model.Article{}
	var categoryID, authorID sql.NullInt64
	var publishTime sql.NullTime

	err := r.db.QueryRow(`
		SELECT id, title, slug, summary, content_md, content_html, cover_image, 
		       status, is_top, view_count, category_id, author_id, publish_time, 
		       created_at, updated_at 
		FROM articles WHERE id = ?
	`, id).Scan(
		&article.ID, &article.Title, &article.Slug, &article.Summary,
		&article.ContentMD, &article.ContentHTML, &article.CoverImage,
		&article.Status, &article.IsTop, &article.ViewCount,
		&categoryID, &authorID, &publishTime,
		&article.CreatedAt, &article.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if categoryID.Valid {
		article.CategoryID = &categoryID.Int64
	}
	if authorID.Valid {
		article.AuthorID = &authorID.Int64
	}
	if publishTime.Valid {
		article.PublishTime = &publishTime.Time
	}

	return article, nil
}

// FindBySlug 根据slug查找文章
func (r *ArticleRepository) FindBySlug(slug string) (*model.Article, error) {
	article := &model.Article{}
	var categoryID, authorID sql.NullInt64
	var publishTime sql.NullTime

	err := r.db.QueryRow(`
		SELECT id, title, slug, summary, content_md, content_html, cover_image, 
		       status, is_top, view_count, category_id, author_id, publish_time, 
		       created_at, updated_at 
		FROM articles WHERE slug = ?
	`, slug).Scan(
		&article.ID, &article.Title, &article.Slug, &article.Summary,
		&article.ContentMD, &article.ContentHTML, &article.CoverImage,
		&article.Status, &article.IsTop, &article.ViewCount,
		&categoryID, &authorID, &publishTime,
		&article.CreatedAt, &article.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if categoryID.Valid {
		article.CategoryID = &categoryID.Int64
	}
	if authorID.Valid {
		article.AuthorID = &authorID.Int64
	}
	if publishTime.Valid {
		article.PublishTime = &publishTime.Time
	}

	return article, nil
}

// List 文章列表
func (r *ArticleRepository) List(req *model.ArticleListRequest) (*model.ArticleListResponse, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if req.Status != nil {
		conditions = append(conditions, "a.status = ?")
		args = append(args, *req.Status)
	}
	if req.CategoryID != nil {
		conditions = append(conditions, "a.category_id = ?")
		args = append(args, *req.CategoryID)
	}
	if req.TagID != nil {
		conditions = append(conditions, "EXISTS (SELECT 1 FROM article_tag at WHERE at.article_id = a.id AND at.tag_id = ?)")
		args = append(args, *req.TagID)
	}
	if req.Keyword != "" {
		conditions = append(conditions, "(a.title LIKE ? OR a.summary LIKE ?)")
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总数
	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM articles a %s", whereClause)
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// 查询列表
	offset := (req.Page - 1) * req.PageSize
	listQuery := fmt.Sprintf(`
		SELECT a.id, a.title, a.slug, a.summary, a.cover_image, 
		       a.status, a.is_top, a.view_count, a.category_id, a.author_id, 
		       a.publish_time, a.created_at, a.updated_at,
		       c.id, c.name, c.slug
		FROM articles a 
		LEFT JOIN categories c ON a.category_id = c.id 
		%s 
		ORDER BY a.is_top DESC, a.publish_time DESC, a.created_at DESC 
		LIMIT ? OFFSET ?
	`, whereClause)

	listArgs := append(args, req.PageSize, offset)
	rows, err := r.db.Query(listQuery, listArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []model.Article{}
	for rows.Next() {
		article := model.Article{}
		var categoryID, authorID sql.NullInt64
		var publishTime sql.NullTime
		var catID, catName, catSlug sql.NullString

		err := rows.Scan(
			&article.ID, &article.Title, &article.Slug, &article.Summary,
			&article.CoverImage, &article.Status, &article.IsTop, &article.ViewCount,
			&categoryID, &authorID, &publishTime,
			&article.CreatedAt, &article.UpdatedAt,
			&catID, &catName, &catSlug,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			article.CategoryID = &categoryID.Int64
		}
		if authorID.Valid {
			article.AuthorID = &authorID.Int64
		}
		if publishTime.Valid {
			article.PublishTime = &publishTime.Time
		}
		if catID.Valid {
			catIDInt := int64(0)
			fmt.Sscanf(catID.String, "%d", &catIDInt)
			article.Category = &model.Category{
				ID:   catIDInt,
				Name: catName.String,
				Slug: catSlug.String,
			}
		}

		articles = append(articles, article)
	}

	return &model.ArticleListResponse{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		List:     articles,
	}, nil
}

// Create 创建文章
func (r *ArticleRepository) Create(article *model.Article) error {
	now := time.Now()
	var publishTime interface{}
	if article.Status == 1 {
		publishTime = now
	}

	result, err := r.db.Exec(`
		INSERT INTO articles (title, slug, summary, content_md, content_html, cover_image, 
		                      status, is_top, view_count, category_id, author_id, publish_time, 
		                      created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		article.Title, article.Slug, article.Summary, article.ContentMD, article.ContentHTML,
		article.CoverImage, article.Status, article.IsTop, 0,
		article.CategoryID, article.AuthorID, publishTime, now, now,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	article.ID = id
	article.ViewCount = 0
	article.CreatedAt = now
	article.UpdatedAt = now
	if article.Status == 1 {
		article.PublishTime = &now
	}
	return nil
}

// Update 更新文章
func (r *ArticleRepository) Update(article *model.Article) error {
	now := time.Now()
	_, err := r.db.Exec(`
		UPDATE articles SET title = ?, slug = ?, summary = ?, content_md = ?, content_html = ?, 
		                   cover_image = ?, status = ?, is_top = ?, category_id = ?, updated_at = ? 
		WHERE id = ?
	`,
		article.Title, article.Slug, article.Summary, article.ContentMD, article.ContentHTML,
		article.CoverImage, article.Status, article.IsTop, article.CategoryID, now, article.ID,
	)
	if err != nil {
		return err
	}
	article.UpdatedAt = now
	return nil
}

// Delete 删除文章
func (r *ArticleRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}

// IncrementViewCount 增加阅读量
func (r *ArticleRepository) IncrementViewCount(id int64) error {
	_, err := r.db.Exec("UPDATE articles SET view_count = view_count + 1 WHERE id = ?", id)
	return err
}

// SetTags 设置文章标签
func (r *ArticleRepository) SetTags(articleID int64, tagIDs []int64) error {
	// 删除原有标签
	_, err := r.db.Exec("DELETE FROM article_tag WHERE article_id = ?", articleID)
	if err != nil {
		return err
	}

	// 添加新标签
	for _, tagID := range tagIDs {
		_, err := r.db.Exec(
			"INSERT INTO article_tag (article_id, tag_id, created_at) VALUES (?, ?, ?)",
			articleID, tagID, time.Now(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTags 获取文章标签
func (r *ArticleRepository) GetTags(articleID int64) ([]model.Tag, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.name, t.slug, t.created_at, t.updated_at 
		FROM tags t 
		JOIN article_tag at ON t.id = at.tag_id 
		WHERE at.article_id = ?
		ORDER BY t.name
	`, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []model.Tag{}
	for rows.Next() {
		tag := model.Tag{}
		err := rows.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
