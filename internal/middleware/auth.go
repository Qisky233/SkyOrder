package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func JWTAuth() gin.HandlerFunc {
	// 从配置文件中读取 JWT 密钥
	secret := viper.GetString("jwt.secret")
	if secret == "" {
		panic("JWT secret key is not set in the configuration")
	}

	return func(c *gin.Context) {
		// 获取 Authorization 头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "未授权访问",
			})
			c.Abort()
			return
		}

		// 检查 Authorization 头的格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "无效的认证格式",
			})
			c.Abort()
			return
		}

		// 解析 JWT Token
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			// 确保 Token 使用了正确的签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "无效的token",
			})
			c.Abort()
			return
		}

		// 提取 Claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "无效的token claims",
			})
			c.Abort()
			return
		}

		// 将 Claims 设置到上下文中
		if adminID, ok := claims["admin_id"].(float64); ok {
			c.Set("admin_id", uint(adminID))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "admin_id is missing or invalid",
			})
			c.Abort()
			return
		}

		if role, ok := claims["role"].(string); ok {
			c.Set("role", role)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "role is missing or invalid",
			})
			c.Abort()
			return
		}

		// 继续处理请求
		c.Next()
	}
}
