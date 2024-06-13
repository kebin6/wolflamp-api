package banner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/wolflamp-api/internal/logic/banner"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
)

// swagger:route post /banner/create_banner banner CreateBanner
//
// createBanner
//
// createBanner
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateBannerReq
//
// Responses:
//  200: BaseMsgResp

func CreateBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBannerReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := banner.NewCreateBannerLogic(r.Context(), svcCtx)
		resp, err := l.CreateBanner(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
