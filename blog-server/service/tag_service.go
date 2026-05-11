package service

import (
	"blog-server/model"
	"blog-server/repository"
	"errors"
)

// TagService 标签服务
type TagService struct {
	tagRepo *repository.TagRepository
}

// NewTagService 创建标签服务
func NewTagService(tagRepo *repository.TagRepository) *TagService {
	return &TagService{tagRepo: tagRepo}
}

// GetByID 根据ID获取标签
func (s *TagService) GetByID(id int64) (*model.Tag, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, errors.New("tag not found")
	}
	return tag, nil
}

// GetBySlug 根据slug获取标签
func (s *TagService) GetBySlug(slug string) (*model.Tag, error) {
	tag, err := s.tagRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, errors.New("tag not found")
	}
	return tag, nil
}

// List 标签列表
func (s *TagService) List() ([]model.Tag, error) {
	return s.tagRepo.List()
}

// Create 创建标签
func (s *TagService) Create(req *model.CreateTagRequest) (*model.Tag, error) {
	// 检查slug是否已存在
	existing, err := s.tagRepo.FindBySlug(req.Slug)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("slug already exists")
	}

	tag := &model.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}

	if err := s.tagRepo.Create(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

// Update 更新标签
func (s *TagService) Update(id int64, req *model.UpdateTagRequest) (*model.Tag, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, errors.New("tag not found")
	}

	// 检查slug是否已被其他标签使用
	if req.Slug != "" && req.Slug != tag.Slug {
		existing, err := s.tagRepo.FindBySlug(req.Slug)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, errors.New("slug already exists")
		}
		tag.Slug = req.Slug
	}

	if req.Name != "" {
		tag.Name = req.Name
	}

	if err := s.tagRepo.Update(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

// Delete 删除标签
func (s *TagService) Delete(id int64) error {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return err
	}
	if tag == nil {
		return errors.New("tag not found")
	}
	return s.tagRepo.Delete(id)
}
