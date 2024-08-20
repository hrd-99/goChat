package logic

import (
	"context"

	"goChat/app/group/rpc/internal/svc"
	"goChat/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserListLogic {
	return &GroupUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupUserListLogic) GroupUserList(in *__.GroupUserListRequest) (*__.GroupUserListResponse, error) {
	// todo: add your logic here and delete this line

	return &__.GroupUserListResponse{}, nil
}
