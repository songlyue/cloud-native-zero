package test

import (
	"cloud-native-zero/app/testApi/cmd/api/internal/logic/test"
	"cloud-native-zero/app/testApi/cmd/api/internal/svc"
	"cloud-native-zero/app/testApi/cmd/api/internal/types"
	"cloud-native-zero/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DbtestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DbRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := test.NewDbtestLogic(r.Context(), svcCtx)
		resp, err := l.Dbtest(&req)
		result.HttpResult(r, w, resp, err)
	}
}
