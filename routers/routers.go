// 路由层 管理程序的路由信息
package routers

import (
	"km_backend/routers/auth"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegistrerRouters(r *gin.Engine) {
	// 1. 登录: /api/auth/login
	// 2. 退出: /api/auth/logout
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
}
