package banner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/wolflamp-api/internal/logic/banner"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
)

// swagger:route post /banner/delete_banner banner DeleteBanner
//
// deleteBanner
//
// deleteBanner
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DeleteBannerReq
//
// Responses:
//  200: BaseMsgResp

func DeleteBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteBannerReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := banner.NewDeleteBannerLogic(r.Context(), svcCtx)
		resp, err := l.DeleteBanner(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
