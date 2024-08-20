package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goChat/app/user/api/internal/logic"
	"goChat/app/user/api/internal/svc"
	"goChat/app/user/api/internal/types"
)

func PersonalInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PersonalInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPersonalInfoLogic(r.Context(), svcCtx)
		resp, err := l.PersonalInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
