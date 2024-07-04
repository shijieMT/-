package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"notify/notify-api/internal/logic"
	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"
)

func setNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetNoticeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSetNoticeLogic(r.Context(), svcCtx)
		resp, err := l.SetNotice(&req)
		//
		_ = resp
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.WriteJson(w, 200, "请求成功")
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
