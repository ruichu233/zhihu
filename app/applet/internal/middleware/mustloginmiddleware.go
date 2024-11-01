package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zhihu/pkg/xcode"
)

type MustLoginMiddleware struct {
}

func NewMustLoginMiddleware() *MustLoginMiddleware {
	return &MustLoginMiddleware{}
}

func (m *MustLoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("user_id") == "" {
			httpx.ErrorCtx(r.Context(), w, xcode.NoLogin)
			return
		}
		next(w, r)
	}
}
