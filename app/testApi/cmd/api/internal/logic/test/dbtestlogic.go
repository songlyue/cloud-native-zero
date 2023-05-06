package test

import (
	"context"

	"cloud-native-zero/app/testApi/cmd/api/internal/svc"
	"cloud-native-zero/app/testApi/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DbtestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDbtestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DbtestLogic {
	return &DbtestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DbtestLogic) Dbtest(req *types.DbRequest) (resp *types.DbResponse, err error) {
	one, err := l.svcCtx.UserMode.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp = new(types.DbResponse)
	resp.Message = one.Name
	return
}
