package overview

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOverviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOverviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOverviewLogic {
	return &GetOverviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetOverviewLogic) GetOverview(req *types.GetOverviewReq) (resp *types.GetOverviewResp, err error) {

	if !l.svcCtx.Config.WolfLampRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	result, err := l.svcCtx.WolfLampRpc.GetOverview(l.ctx, &wolflamp.GetOverviewReq{StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		return nil, err
	}

	return &types.GetOverviewResp{
		Data: types.OverviewInfo{
			TodayParticipateCount: result.Data.TodayParticipateCount,
			TodayNewPlayerCount:   result.Data.TodayNewPlayerCount,
			TodayRoundCount:       result.Data.TodayRoundCount,
			TodayEatCount:         result.Data.TodayEatCount,
			TodayPlatformProfit:   result.Data.TodayPlatformProfit,
			TotalPlatformProfit:   result.Data.TotalPlatformProfit,
			TotalPlayerCount:      result.Data.TotalPlayerCount,
		},
	}, nil

}
