package logic

import (
	"context"

	"goChat/app/group/rpc/internal/svc"
	"goChat/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendLogic) AddFriend(in *__.AddFriendRequest) (*__.AddFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &__.AddFriendResponse{}, nil
}
