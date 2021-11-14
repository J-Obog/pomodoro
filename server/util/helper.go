package apputil

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


func CreateAuthToken(expHrs int, sub uint64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        fmt.Sprint(sub),
		ExpiresAt: time.Now().Add(time.Duration(expHrs) * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	})

	tokenStr, e := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if e != nil {
		return ""
	}
	
	return tokenStr
}