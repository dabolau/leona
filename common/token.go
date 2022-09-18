package common

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 自定义结构体
type JwtClaims struct {
	Username             string `json:"username"` // 用户名
	jwt.RegisteredClaims        // 包的默认结构体
}

var (
	// 签发人
	// JwtIssuer = "默认签发人"
	// 密钥
	JwtSecret = []byte("春花秋月何时了")
)

// 生成令牌
// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#NewWithClaims
func GenerateToken(username string) (string, error) {
	// 自定义声明
	jwtClaims := JwtClaims{
		Username: username, // 用户名
		RegisteredClaims: jwt.RegisteredClaims{
			// Issuer:    JwtIssuer,                                                             // 签发人
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * (30 * 24 * 60 * 60))), // 过期时间，30天
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	// 使用指定的密钥（JwtSecret）签名并获得完整的编码后的令牌字符串
	tokenString, err := token.SignedString(JwtSecret)
	return tokenString, err
}

// 解析令牌
// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#ParseWithClaims
func ParseToken(tokenString string) (*JwtClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	// 解析声明
	jwtClaims, ok := token.Claims.(*JwtClaims)
	// 校验令牌
	if ok && token.Valid {
		return jwtClaims, nil
	}
	return nil, err
}
