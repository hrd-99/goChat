package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goChat/app/msg/api/internal/logic"
	"goChat/app/msg/api/internal/svc"
	"goChat/app/msg/api/internal/types"
)

func uploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
