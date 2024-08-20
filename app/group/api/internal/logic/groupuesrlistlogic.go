package logic

import (
	"context"

	"goChat/app/group/api/internal/svc"
	"goChat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUesrListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUesrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUesrListLogic {
	return &GroupUesrListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUesrListLogic) GroupUesrList(req *types.GroupUserListRequest) (resp *types.GroupUserListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
