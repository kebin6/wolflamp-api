package player

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/wolflamp-api/internal/logic/player"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
)

// swagger:route post /player/update_player player UpdatePlayer
//
// updatePlayer
//
// updatePlayer
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdatePlayerReq
//
// Responses:
//  200: BaseMsgResp

func UpdatePlayerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePlayerReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := player.NewUpdatePlayerLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePlayer(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
