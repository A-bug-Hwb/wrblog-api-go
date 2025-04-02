package token

import (
	"github.com/golang-jwt/jwt"
	"wrblog-api-go/config"
)

// CustomClaims 自定义 Payload 信息
type CustomClaims struct {
	Uid    any // 唯一标识
	UserId any // userId
	jwt.StandardClaims
}

type JWT struct {
	singKey []byte // Jwt 密钥
}

// NewJWT 返回一个JWT 实例
func NewJWT() *JWT {
	return &JWT{
		singKey: []byte(config.Conf.ConfigInfo.Token.Secret),
	}
}

// GenerateToken 创建新的 Token
func (j *JWT) GenerateToken(claims *CustomClaims) (token string, err error) {
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(j.singKey)
}

// ParseToken 验证 Token
func (j *JWT) ParseToken(token string) (claims *CustomClaims, ok bool) {
	ok = false
	withClaims, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.singKey, nil
	})
	if withClaims != nil {
		claims, ok = withClaims.Claims.(*CustomClaims)
	}
	return
}
