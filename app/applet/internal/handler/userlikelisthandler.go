package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zhihu/app/applet/internal/logic"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
)

func UserLikeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLikeListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLikeListLogic(r.Context(), svcCtx)
		resp, err := l.UserLikeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
