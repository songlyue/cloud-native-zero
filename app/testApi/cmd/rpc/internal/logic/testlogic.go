package logic

import (
	"context"

	"cloud-native-zero/app/testApi/cmd/rpc/internal/svc"
	"cloud-native-zero/app/testApi/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestLogic) Test(in *pb.TestReq) (*pb.TestResp, error) {
	return &pb.TestResp{
		Message: "萨法",
	}, nil
}
