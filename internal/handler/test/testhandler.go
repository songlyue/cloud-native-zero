package test

import (
	"cloud-native-zero/app/order/cmd/api/internal/logic/test"
	"cloud-native-zero/app/order/cmd/api/internal/svc"
	"cloud-native-zero/app/order/cmd/api/internal/types"
	"cloud-native-zero/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := test.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		result.HttpResult(r, w, resp, err)
	}
}
