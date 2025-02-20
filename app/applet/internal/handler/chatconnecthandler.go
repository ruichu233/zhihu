package handler

import (
	"net/http"
	"strconv"

	"zhihu/app/applet/internal/logic"
	"zhihu/app/applet/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewChatConnectLogic(r.Context(), svcCtx, w, r, userId)
		err = l.ChatConnect()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
