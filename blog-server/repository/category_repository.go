package repository

import (
	"blog-server/model"
	"database/sql"
	"time"
)

// CategoryRepository 分类仓储
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository 创建分类仓储
func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// FindByID 根据ID查找分类
func (r *CategoryRepository) FindByID(id int64) (*model.Category, error) {
	category := &model.Category{}
	err := r.db.QueryRow(
		"SELECT id, name, slug, description, sort_order, created_at, updated_at FROM categories WHERE id = ?",
		id,
	).Scan(&category.ID, &category.Name, &category.Slug, &category.Description,
		&category.SortOrder, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return category, nil
}

// FindBySlug 根据slug查找分类
func (r *CategoryRepository) FindBySlug(slug string) (*model.Category, error) {
	category := &model.Category{}
	err := r.db.QueryRow(
		"SELECT id, name, slug, description, sort_order, created_at, updated_at FROM categories WHERE slug = ?",
		slug,
	).Scan(&category.ID, &category.Name, &category.Slug, &category.Description,
		&category.SortOrder, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return category, nil
}

// List 分类列表
func (r *CategoryRepository) List() ([]model.Category, error) {
	rows, err := r.db.Query(
		"SELECT id, name, slug, description, sort_order, created_at, updated_at FROM categories ORDER BY sort_order ASC, id ASC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []model.Category{}
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Slug, &category.Description,
			&category.SortOrder, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// Create 创建分类
func (r *CategoryRepository) Create(category *model.Category) error {
	now := time.Now()
	result, err := r.db.Exec(
		"INSERT INTO categories (name, slug, description, sort_order, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		category.Name, category.Slug, category.Description, category.SortOrder, now, now,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	category.ID = id
	category.CreatedAt = now
	category.UpdatedAt = now
	return nil
}

// Update 更新分类
func (r *CategoryRepository) Update(category *model.Category) error {
	now := time.Now()
	_, err := r.db.Exec(
		"UPDATE categories SET name = ?, slug = ?, description = ?, sort_order = ?, updated_at = ? WHERE id = ?",
		category.Name, category.Slug, category.Description, category.SortOrder, now, category.ID,
	)
	if err != nil {
		return err
	}
	category.UpdatedAt = now
	return nil
}

// Delete 删除分类
func (r *CategoryRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}
