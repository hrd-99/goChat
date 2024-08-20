package logic

import (
	"context"

	"goChat/app/group/api/internal/svc"
	"goChat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageGroupInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageGroupInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageGroupInfoListLogic {
	return &MessageGroupInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageGroupInfoListLogic) MessageGroupInfoList(req *types.MessageGroupInfoListRequest) (resp *types.MessageGroupInfoListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
