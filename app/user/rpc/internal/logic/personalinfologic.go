package logic

import (
	"context"

	"goChat/app/user/rpc/internal/svc"
	"goChat/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonalInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonalInfoLogic {
	return &PersonalInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersonalInfoLogic) PersonalInfo(in *__.PersonalInfoRequest) (*__.PersonalInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &__.PersonalInfoResponse{}, nil
}
