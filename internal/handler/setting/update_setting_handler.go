package setting

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/wolflamp-api/internal/logic/setting"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
)

// swagger:route post /setting/update_setting setting UpdateSetting
//
// updateSetting
//
// updateSetting
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PlatformSetting
//
// Responses:
//  200: BaseMsgResp

func UpdateSettingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSettingReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := setting.NewUpdateSettingLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSetting(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
