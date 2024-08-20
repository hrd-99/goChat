package logic

import (
	"context"

	"goChat/app/user/rpc/internal/svc"
	"goChat/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *__.LoginRequest) (*__.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &__.LoginResponse{}, nil
}
