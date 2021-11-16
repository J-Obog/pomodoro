package apputils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Content-Type, Authorization, cache-control")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}


func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("<" + r.Method + "> " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}


func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(os.Getenv("JWT_HEADER"))
		
		if token != "" {
			jwtToken, e := VerifyJWTToken(token)
			
			if e != nil {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(map[string]interface{}{"message": e.Error()})
				return
			}
			
			ctx := context.WithValue(r.Context(), "jwt", jwtToken)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Authorization header missing"})
	})
}