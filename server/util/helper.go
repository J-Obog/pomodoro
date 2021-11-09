package apputil

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateAuthToken(exp int64, sub uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        fmt.Sprint(sub),
		ExpiresAt: exp,
		IssuedAt:  time.Now().Unix(),
	})

	tokenStr, e := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if e != nil {
		return ""
	}
	return tokenStr
}