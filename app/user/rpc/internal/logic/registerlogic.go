package logic

import (
	"context"

	"goChat/app/user/rpc/internal/svc"
	"goChat/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *__.RegisterRequest) (*__.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &__.RegisterResponse{}, nil
}
