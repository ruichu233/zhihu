package middleware

import (
	"net/http"
	"zhihu/pkg/token"
)

const (
	AuthKey = "access-token"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(AuthKey)
		if accessToken == "" {
			next(w, r)
			return
		}
		payload, err := token.Parse(accessToken, "")
		if err != nil {
			// todo 过期token更新
			next(w, r)
			return
		}
		r.Header.Set("user_id", payload)
		next(w, r)
	}
}
