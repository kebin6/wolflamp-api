package banner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/wolflamp-api/internal/logic/banner"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
)

// swagger:route post /banner/find_banner banner FindBanner
//
// findBanner
//
// findBanner
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: FindBannerReq
//
// Responses:
//  200: FindBannerResp

func FindBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindBannerReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := banner.NewFindBannerLogic(r.Context(), svcCtx)
		resp, err := l.FindBanner(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
