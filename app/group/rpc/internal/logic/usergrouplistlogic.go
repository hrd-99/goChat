package logic

import (
	"context"

	"goChat/app/group/rpc/internal/svc"
	"goChat/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupListLogic {
	return &UserGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGroupListLogic) UserGroupList(in *__.UserGroupListRequest) (*__.UserGroupListResponse, error) {
	// todo: add your logic here and delete this line

	return &__.UserGroupListResponse{}, nil
}
