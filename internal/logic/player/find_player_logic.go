package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPlayerLogic {
	return &FindPlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *FindPlayerLogic) FindPlayer(req *types.FindPlayerReq) (resp *types.FindPlayerResp, err error) {

	if !l.svcCtx.Config.WolfLampRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.WolfLampRpc.FindPlayer(l.ctx, &wolflamp.FindPlayerReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.FindPlayerResp{
		Data: types.PlayerInfo{
			Id:                  data.Id,
			CreatedAt:           data.CreatedAt,
			UpdatedAt:           data.UpdatedAt,
			Status:              data.Status,
			Rank:                data.Rank,
			Amount:              data.Amount,
			InvitedCode:         data.InvitedCode,
			Name:                data.Name,
			InvitedNum:          data.InvitedNum,
			TotalIncome:         data.TotalIncome,
			ProfitAndLoss:       data.ProfitAndLoss,
			Recent100WinPercent: data.Recent_100WinPercent,
			InviteCode:          data.InviteCode,
		},
	}, nil

}
