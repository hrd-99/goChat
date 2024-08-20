package logic

import (
	"context"

	"goChat/app/user/api/internal/svc"
	"goChat/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonalInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonalInfoLogic {
	return &PersonalInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersonalInfoLogic) PersonalInfo(req *types.PersonalInfoRequest) (resp *types.PersonalInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
