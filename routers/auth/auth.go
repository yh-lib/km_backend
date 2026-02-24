package auth

import (
	"km_backend/controllers/auth"

	"github.com/gin-gonic/gin"
)

// 实现登录接口
func login(loginGroup *gin.RouterGroup) {
	loginGroup.POST("/login", auth.Login)
}

// 实现退出接口
func logout(logoutGroup *gin.RouterGroup) {
	logoutGroup.GET("/logout", auth.Logout)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	authGroup := g.Group("auth")
	// 登录的功能
	login(authGroup)
	// 退出接口
	logout(authGroup)
}
