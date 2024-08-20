package logic

import (
	"context"

	"goChat/app/msg/rpc/internal/svc"
	"goChat/app/msg/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLogic {
	return &PullLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PullLogic) Pull(in *__.PullRequest) (*__.PullResponse, error) {
	// todo: add your logic here and delete this line

	return &__.PullResponse{}, nil
}
