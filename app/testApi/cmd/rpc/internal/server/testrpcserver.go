// Code generated by goctl. DO NOT EDIT.
// Source: test.proto

package server

import (
	"context"

	"cloud-native-zero/app/testApi/cmd/rpc/internal/logic"
	"cloud-native-zero/app/testApi/cmd/rpc/internal/svc"
	"cloud-native-zero/app/testApi/cmd/rpc/pb"
)

type TestRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTestRpcServer
}

func NewTestRpcServer(svcCtx *svc.ServiceContext) *TestRpcServer {
	return &TestRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *TestRpcServer) Test(ctx context.Context, in *pb.TestReq) (*pb.TestResp, error) {
	l := logic.NewTestLogic(ctx, s.svcCtx)
	return l.Test(in)
}
