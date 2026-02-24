// 中间件层
package middlerwares

import (
	"km_backend/config"
	"km_backend/utils/jwtutils"

	"github.com/gin-gonic/gin"
)

func JwtAuth(c *gin.Context) {
	// 1. 除了 login 和 logout 之外的所有请求都需要验证token是否合法
	requestUrl := c.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		c.Next()
		return
	}
	returnData := config.NewRetrunData()
	// 2. 其它接口需要验证token
	tokenString := c.Request.Header.Get("Authorization")
	// token 为空时
	if tokenString == "" {
		returnData.Status = 401
		returnData.Message = "请求未携带token,请登录后再尝试。"
		c.JSON(200, returnData)
		c.Abort()
		return
	}
	// token 不为空时
	claims, err := jwtutils.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "token 验证未通过。"
		c.JSON(200, returnData)
		c.Abort()
		return
	} else {
		returnData.Status = 200
		returnData.Message = "请求成功。"
		c.JSON(200, returnData)
	}
	// 将 claims 赋值到 *gin.Context 中
	c.Set("claims", claims)
	c.Next()
}
