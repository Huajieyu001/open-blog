package service

import (
	"blog-server/config"
	"blog-server/model"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// UploadService 文件上传服务
type UploadService struct {
	config config.OSSConfig
}

// NewUploadService 创建文件上传服务
func NewUploadService(config config.OSSConfig) *UploadService {
	return &UploadService{config: config}
}

// Upload 上传文件（暂时返回本地路径，后续集成OSS）
func (s *UploadService) Upload(file io.Reader, filename string) (*model.UploadResponse, error) {
	// 生成唯一文件名
	ext := filepath.Ext(filename)
	newFilename := fmt.Sprintf("uploads/%s/%s%s",
		time.Now().Format("2006/01/02"),
		uuid.New().String(),
		ext,
	)

	// TODO: 实际上传到OSS
	// 暂时返回占位URL
	url := fmt.Sprintf("/%s", newFilename)

	return &model.UploadResponse{
		URL:      url,
		Filename: newFilename,
	}, nil
}
