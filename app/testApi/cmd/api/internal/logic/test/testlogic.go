package test

import (
	"cloud-native-zero/app/testApi/cmd/api/client/pb"
	"cloud-native-zero/app/testApi/cmd/api/internal/svc"
	"cloud-native-zero/app/testApi/cmd/api/internal/types"
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
	testReq := pb.TestReq{Id: 123123}
	test, err := l.svcCtx.TestRpc.Test(l.ctx, &testReq)
	if err != nil {
		return nil, err
	}
	resp = new(types.Response)

	resp.Message = test.Message
	return
}
