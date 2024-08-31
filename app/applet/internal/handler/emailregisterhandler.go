package handler

import (
	"net/http"
	"zhihu/app/applet/internal/logic"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmailRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEmailRegisterLogic(r.Context(), svcCtx)
		resp, err := l.EmailRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
