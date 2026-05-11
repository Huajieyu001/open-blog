package service

import (
	"blog-server/model"
	"blog-server/repository"
	"errors"
)

// CategoryService 分类服务
type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

// NewCategoryService 创建分类服务
func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

// GetByID 根据ID获取分类
func (s *CategoryService) GetByID(id int64) (*model.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

// GetBySlug 根据slug获取分类
func (s *CategoryService) GetBySlug(slug string) (*model.Category, error) {
	category, err := s.categoryRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

// List 分类列表
func (s *CategoryService) List() ([]model.Category, error) {
	return s.categoryRepo.List()
}

// Create 创建分类
func (s *CategoryService) Create(req *model.CreateCategoryRequest) (*model.Category, error) {
	// 检查slug是否已存在
	existing, err := s.categoryRepo.FindBySlug(req.Slug)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("slug already exists")
	}

	category := &model.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

// Update 更新分类
func (s *CategoryService) Update(id int64, req *model.UpdateCategoryRequest) (*model.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	// 检查slug是否已被其他分类使用
	if req.Slug != "" && req.Slug != category.Slug {
		existing, err := s.categoryRepo.FindBySlug(req.Slug)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, errors.New("slug already exists")
		}
		category.Slug = req.Slug
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	if req.SortOrder != nil {
		category.SortOrder = *req.SortOrder
	}

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

// Delete 删除分类
func (s *CategoryService) Delete(id int64) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("category not found")
	}
	return s.categoryRepo.Delete(id)
}
