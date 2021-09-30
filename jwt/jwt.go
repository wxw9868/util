package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims Payload 结构体
type CustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid uint `json:"uid"`
	//Mobile string `json:"mobile"`
	Admin bool `json:"admin"`
}

// CreateToken 生成 token SecretKey 是一个 const 常量
func CreateToken(secretKey string, issuer string, uid uint, isAdmin bool) (tokenString string, err error) {
	claims := &CustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*7,
			Issuer:    issuer,
		},
		uid,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secretKey))
	return
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// ParseToken 解析 token
func ParseToken(tokenString string, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
