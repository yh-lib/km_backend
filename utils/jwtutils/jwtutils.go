package jwtutils

import (
	"errors"
	"km_backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

// 1. 自定义声明类型
type MyCustomClaims struct {
	// 可根据需要自行添加字段
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// 2. 封装生成token的函数
// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	claims := MyCustomClaims{
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Minute * time.Duration(config.JwtExpTime)))), // token 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "YH", // 签发人
			Subject:   "YH",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(jwtSignKey)
}

// 2. 封装解析token的函数
// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSignKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
