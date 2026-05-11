package config

import (
	"os"
)

// Config 应用配置
type Config struct {
	Port      string
	DBPath    string
	JWTSecret string
	OSS       OSSConfig
}

// OSSConfig 阿里云 OSS 配置
type OSSConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	Domain          string // 访问域名，如 https://your-bucket.oss-cn-hangzhou.aliyuncs.com
}

// Load 从环境变量加载配置
func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		DBPath:    getEnv("DB_PATH", "./data/blog.db"),
		JWTSecret: getEnv("JWT_SECRET", "your-jwt-secret-key-change-in-production"),
		OSS: OSSConfig{
			Endpoint:        getEnv("OSS_ENDPOINT", ""),
			AccessKeyID:     getEnv("OSS_ACCESS_KEY_ID", ""),
			AccessKeySecret: getEnv("OSS_ACCESS_KEY_SECRET", ""),
			BucketName:      getEnv("OSS_BUCKET_NAME", ""),
			Domain:          getEnv("OSS_DOMAIN", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
