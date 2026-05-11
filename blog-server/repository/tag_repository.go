package repository

import (
	"blog-server/model"
	"database/sql"
	"time"
)

// TagRepository 标签仓储
type TagRepository struct {
	db *sql.DB
}

// NewTagRepository 创建标签仓储
func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db: db}
}

// FindByID 根据ID查找标签
func (r *TagRepository) FindByID(id int64) (*model.Tag, error) {
	tag := &model.Tag{}
	err := r.db.QueryRow(
		"SELECT id, name, slug, created_at, updated_at FROM tags WHERE id = ?",
		id,
	).Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return tag, nil
}

// FindBySlug 根据slug查找标签
func (r *TagRepository) FindBySlug(slug string) (*model.Tag, error) {
	tag := &model.Tag{}
	err := r.db.QueryRow(
		"SELECT id, name, slug, created_at, updated_at FROM tags WHERE slug = ?",
		slug,
	).Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return tag, nil
}

// List 标签列表
func (r *TagRepository) List() ([]model.Tag, error) {
	rows, err := r.db.Query(
		"SELECT id, name, slug, created_at, updated_at FROM tags ORDER BY name ASC",
	)
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

// Create 创建标签
func (r *TagRepository) Create(tag *model.Tag) error {
	now := time.Now()
	result, err := r.db.Exec(
		"INSERT INTO tags (name, slug, created_at, updated_at) VALUES (?, ?, ?, ?)",
		tag.Name, tag.Slug, now, now,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	tag.ID = id
	tag.CreatedAt = now
	tag.UpdatedAt = now
	return nil
}

// Update 更新标签
func (r *TagRepository) Update(tag *model.Tag) error {
	now := time.Now()
	_, err := r.db.Exec(
		"UPDATE tags SET name = ?, slug = ?, updated_at = ? WHERE id = ?",
		tag.Name, tag.Slug, now, tag.ID,
	)
	if err != nil {
		return err
	}
	tag.UpdatedAt = now
	return nil
}

// Delete 删除标签
func (r *TagRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM tags WHERE id = ?", id)
	return err
}
