package logic

import (
	"context"

	"goChat/app/msg/rpc/internal/svc"
	"goChat/app/msg/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *__.UploadRequest) (*__.UploadResponse, error) {
	// todo: add your logic here and delete this line

	return &__.UploadResponse{}, nil
}
