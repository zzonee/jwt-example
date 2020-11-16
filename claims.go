package jwt_example

import "github.com/dgrijalva/jwt-go"

// 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
	UserId int32
}
