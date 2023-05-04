// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	test "cloud-native-zero/app/order/cmd/api/internal/handler/test"
	"cloud-native-zero/app/order/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/test",
				Handler: test.TestHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
