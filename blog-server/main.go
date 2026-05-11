package main

import (
	"blog-server/config"
	"blog-server/controller/admin"
	"blog-server/controller/pub"
	"blog-server/middleware"
	"blog-server/repository"
	"blog-server/service"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db, err := repository.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// 初始化仓储层
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	tagRepo := repository.NewTagRepository(db)

	// 初始化服务层
	userService := service.NewUserService(userRepo)
	articleService := service.NewArticleService(articleRepo, categoryRepo, tagRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	tagService := service.NewTagService(tagRepo)
	uploadService := service.NewUploadService(cfg.OSS)

	// 初始化控制器
	authController := admin.NewAuthController(userService)
	adminArticleController := admin.NewArticleController(articleService)
	adminCategoryController := admin.NewCategoryController(categoryService)
	adminTagController := admin.NewTagController(tagService)
	uploadController := admin.NewUploadController(uploadService)

	pubArticleController := pub.NewArticleController(articleService)
	pubCategoryController := pub.NewCategoryController(categoryService)
	pubTagController := pub.NewTagController(tagService)

	// 创建 Gin 引擎
	r := gin.Default()

	// CORS 中间件
	r.Use(middleware.CORS())

	// 静态文件服务（前端构建产物）
	frontendPath := filepath.Join(".", "frontend")
	if _, err := os.Stat(frontendPath); err == nil {
		r.Static("/assets", filepath.Join(frontendPath, "assets"))
		r.StaticFile("/favicon.ico", filepath.Join(frontendPath, "favicon.ico"))
	}

	// API 路由
	api := r.Group("/api")
	{
		// 公开接口
		pubGroup := api.Group("/pub")
		{
			pubGroup.GET("/articles", pubArticleController.List)
			pubGroup.GET("/articles/:id", pubArticleController.GetByID)
			pubGroup.GET("/articles/slug/:slug", pubArticleController.GetBySlug)
			pubGroup.GET("/categories", pubCategoryController.List)
			pubGroup.GET("/tags", pubTagController.List)
		}

		// 管理接口（需要 JWT 鉴权）
		adminGroup := api.Group("/admin")
		{
			// 登录接口（无需鉴权）
			adminGroup.POST("/login", authController.Login)

			// 需要鉴权的接口
			authRequired := adminGroup.Group("")
			authRequired.Use(middleware.JWTAuth(cfg.JWTSecret))
			{
				// 文章管理
				authRequired.GET("/articles", adminArticleController.List)
				authRequired.POST("/articles", adminArticleController.Create)
				authRequired.PUT("/articles/:id", adminArticleController.Update)
				authRequired.DELETE("/articles/:id", adminArticleController.Delete)

				// 分类管理
				authRequired.GET("/categories", adminCategoryController.List)
				authRequired.POST("/categories", adminCategoryController.Create)
				authRequired.PUT("/categories/:id", adminCategoryController.Update)
				authRequired.DELETE("/categories/:id", adminCategoryController.Delete)

				// 标签管理
				authRequired.GET("/tags", adminTagController.List)
				authRequired.POST("/tags", adminTagController.Create)
				authRequired.PUT("/tags/:id", adminTagController.Update)
				authRequired.DELETE("/tags/:id", adminTagController.Delete)

				// 文件上传
				authRequired.POST("/upload", uploadController.Upload)
			}
		}
	}

	// 前端路由（SPA 回退）
	r.NoRoute(func(c *gin.Context) {
		// 如果是 API 请求，返回 404
		if len(c.Request.URL.Path) >= 4 && c.Request.URL.Path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"error": "API not found"})
			return
		}
		// 否则返回前端的 index.html
		c.File(filepath.Join(frontendPath, "index.html"))
	})

	// 启动服务器
	addr := ":" + cfg.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
