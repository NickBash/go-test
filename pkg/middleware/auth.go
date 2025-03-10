package middleware

import (
	"context"
	"http/test/configs"
	"http/test/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
	return
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		isValid, data := jwt.NewJwt(config.Auth.Secret).Verify(token)

		if !isValid {
			writeUnauthed(w)
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
