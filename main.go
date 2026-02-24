// 项目的总入口
package main

import (
	"km_backend/config"
	"km_backend/middlerwares"
	"km_backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 注册gin引擎
	r := gin.Default()
	// 使用中间件鉴权
	r.Use(middlerwares.JwtAuth)
	// 注册路由
	routers.RegistrerRouters(r)
	// 启动程序
	r.Run(config.Port)
}
