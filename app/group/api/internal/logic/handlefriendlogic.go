package logic

import (
	"context"

	"goChat/app/group/api/internal/svc"
	"goChat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleFriendLogic {
	return &HandleFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleFriendLogic) HandleFriend(req *types.HandleFriendRequest) (resp *types.HandleFriendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
