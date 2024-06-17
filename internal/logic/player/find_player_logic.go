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
		Data: l.Po2Vo(data),
	}, nil

}

func (l *FindPlayerLogic) Po2Vo(po *wolflamp.PlayerInfo) (vo types.PlayerInfo) {

	return types.PlayerInfo{
		Id:                  po.Id,
		CreatedAt:           po.CreatedAt,
		UpdatedAt:           po.UpdatedAt,
		Status:              po.Status,
		Rank:                po.Rank,
		Amount:              po.Amount,
		InvitedCode:         po.InvitedCode,
		InvitedNum:          po.InvitedNum,
		TotalIncome:         po.TotalIncome,
		Recent100WinPercent: po.Recent_100WinPercent,
		InviteCode:          po.InviteCode,
		Lamb:                po.Lamp,
		SystemCommission:    po.SystemCommission,
		Email:               po.Email,
		DepositAddress:      po.DepositAddress,
	}

}
