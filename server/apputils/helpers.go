package apputils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/J-Obog/pomodoro/data"
	"github.com/dgrijalva/jwt-go"
)


func CreateAuthToken(expHrs int, sub uint64, tokenType string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = fmt.Sprint(sub)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Duration(expHrs) * time.Hour).Unix()
	claims["tkn"] = tokenType

	tokenStr, e := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if e != nil {
		return ""
	}
	
	return tokenStr
}

func VerifyJWTToken(authToken string) (*jwt.Token, error) {
	token, e := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		tokenType := token.Claims.(jwt.MapClaims)["tkn"]
		tokenSub := token.Claims.(jwt.MapClaims)["sub"]

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		
		if tokenType != "access" {
			return nil, errors.New("Invalid token type")
		}

		if _, e := data.RS.Get(context.Background(), fmt.Sprintf("token.%s.%s", tokenSub, authToken)).Result(); e == nil {
			return nil, errors.New("Token has been blacklisted")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if e != nil {
		return nil, e
	}
	
	return token, nil
}	

func GetTokenJTI(r *http.Request) (uint64) {
	jwtToken := r.Context().Value("jwt").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	jtiStr := claims["sub"].(string)

	jti, _ := strconv.ParseUint(jtiStr, 10, 64)
	return jti
}

func GetTokenRaw(r *http.Request) (string) {
	jwtToken := r.Context().Value("jwt").(*jwt.Token)
	return jwtToken.Raw
}