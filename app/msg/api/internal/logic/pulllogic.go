package logic

import (
	"context"

	"goChat/app/msg/api/internal/svc"
	"goChat/app/msg/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLogic {
	return &PullLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullLogic) Pull(req *types.PullRequest) (resp *types.PullResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
