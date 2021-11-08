package apputil

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	rcache "github.com/J-Obog/pomodoro/cache"
	"github.com/dgrijalva/jwt-go"
)

// middleware for handling CORS configs
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Content-Type, Authorization, cache-control")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		// handle OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// middleware for logging requests
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("<" + r.Method + "> " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

//middleware for verifying jwt token
func JwtRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_token := r.Header.Get(os.Getenv("JWT_HEADER"))
		
		if auth_token == "" {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Authorization header missing"})
			return
		}

		if token, e := jwt.Parse(auth_token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		}); e != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Authorization token has expired"})
		} else {
			jti := token.Claims.(jwt.MapClaims)["jti"]
			if _, e := rcache.RS.Get(rcache.CTX, fmt.Sprintf("token-%s", jti)).Result(); e == nil {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid authorization token"})
			} else {
				ctx := context.WithValue(r.Context(), "jti", jti)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}

	})
}
