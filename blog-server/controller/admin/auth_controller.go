package admin

import (
	"blog-server/middleware"
	"blog-server/model"
	"blog-server/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthController 认证控制器
type AuthController struct {
	userService *service.UserService
	jwtSecret   string
}

// NewAuthController 创建认证控制器
func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
		jwtSecret:   "your-jwt-secret-key-change-in-production", // 应从配置读取
	}
}

// Login 管理员登录
func (c *AuthController) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 验证用户名和密码
	user, err := c.userService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成JWT Token
	claims := &middleware.JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.jwtSecret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, model.LoginResponse{
		Token: tokenString,
		User:  *user,
	})
}
