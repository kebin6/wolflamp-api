package overview

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/wolflamp-api/internal/logic/overview"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
)

// swagger:route post /overview overview GetOverview
//
// 获取总揽数据
//
// 获取总揽数据
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetOverviewReq
//
// Responses:
//  200: GetOverviewResp

func GetOverviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOverviewReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := overview.NewGetOverviewLogic(r.Context(), svcCtx)
		resp, err := l.GetOverview(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
