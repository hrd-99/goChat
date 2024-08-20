package logic

import (
	"context"

	"goChat/app/group/rpc/internal/svc"
	"goChat/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleFriendLogic {
	return &HandleFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandleFriendLogic) HandleFriend(in *__.HandleFriendRequest) (*__.HandleFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &__.HandleFriendResponse{}, nil
}
