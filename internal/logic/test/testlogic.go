package test

import (
	"cloud-native-zero/app/order/cmd/api/internal/svc"
	"cloud-native-zero/app/order/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// todo
	return
}
