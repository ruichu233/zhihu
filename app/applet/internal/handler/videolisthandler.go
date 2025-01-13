package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zhihu/app/applet/internal/logic"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
)

func VideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		userIdStr := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		l := logic.NewVideoListLogic(r.Context(), svcCtx)
		resp, err := l.VideoList(&req, userId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			data := map[string]interface{}{
				"data": resp,
			}
			httpx.OkJsonCtx(r.Context(), w, data)
		}
	}
}
