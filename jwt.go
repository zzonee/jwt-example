package jwt_example

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

var signedKey = []byte("secret")

func CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signedKey)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(*jwt.Token) (i interface{}, e error) {
		return signedKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.Request.Header.Get("Authorization")
		ss := strings.SplitN(t, " ", 2)
		if len(ss) !=  2 {
			_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("can't get header:Authorization"))
			return
		}
		claims, err := ParseToken(ss[1])
		if err != nil {
			_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("token invalid:%s", err.Error()))
			return
		}
		log.Printf("extract claims:%v", claims)
		ctx.Next()
	}
}

func IssueToken(ctx *gin.Context) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 10), // 有效期10s
		},
		UserId: 10010,
	}
	token, err := CreateToken(claims)
	if err != nil {
		ctx.String(http.StatusOK, "issue token failed. err:%s", err.Error())
		return
	}
	ctx.String(http.StatusOK, token)
}