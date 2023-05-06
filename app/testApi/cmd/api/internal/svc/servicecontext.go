package svc

import (
	"cloud-native-zero/app/testApi/cmd/api/client/pb"
	"cloud-native-zero/app/testApi/cmd/api/client/testrpc"
	"cloud-native-zero/app/testApi/cmd/api/internal/config"
	"cloud-native-zero/app/testApi/cmd/api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserMode model.UserModel
	TestRpc  pb.TestRpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserMode: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		TestRpc:  testrpc.NewTestRpc(zrpc.MustNewClient(c.TestRpcConf)),
	}
}
