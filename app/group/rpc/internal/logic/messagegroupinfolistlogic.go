package logic

import (
	"context"

	"goChat/app/group/rpc/internal/svc"
	"goChat/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageGroupInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageGroupInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageGroupInfoListLogic {
	return &MessageGroupInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageGroupInfoListLogic) MessageGroupInfoList(in *__.MessageGroupInfoListRequest) (*__.MessageGroupInfoListResponse, error) {
	// todo: add your logic here and delete this line

	return &__.MessageGroupInfoListResponse{}, nil
}
