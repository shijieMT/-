package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"notify/notify-api/internal/logic"
	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"
)

func updateNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateNoticeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateNoticeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateNotice(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
