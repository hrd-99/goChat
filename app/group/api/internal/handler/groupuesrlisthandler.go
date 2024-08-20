package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goChat/app/group/api/internal/logic"
	"goChat/app/group/api/internal/svc"
	"goChat/app/group/api/internal/types"
)

func GroupUesrListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupUserListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGroupUesrListLogic(r.Context(), svcCtx)
		resp, err := l.GroupUesrList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
