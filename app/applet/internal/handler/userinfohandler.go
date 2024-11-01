package handler

import (
	"net/http"
	"zhihu/app/applet/internal/logic"
	"zhihu/app/applet/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserId int64 `json:"user_id"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(req.UserId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}