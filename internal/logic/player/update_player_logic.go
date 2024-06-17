package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePlayerLogic {
	return &UpdatePlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdatePlayerLogic) UpdatePlayer(req *types.UpdatePlayerReq) (resp *types.BaseMsgResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.UpdatePlayer(l.ctx,
		&wolflamp.UpdatePlayerReq{
			Id:               req.Id,
			Status:           &req.Status,
			Rank:             &req.Rank,
			Amount:           &req.Amount,
			Lamp:             &req.Lamb,
			DepositAddress:   &req.DepositAddress,
			SystemCommission: &req.SystemCommission,
			Email:            &req.Email,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil

}
