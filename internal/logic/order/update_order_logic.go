package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.UpdateOrderReq) (resp *types.BaseMsgResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.UpdateOrder(l.ctx,
		&wolflamp.UpdateOrderReq{
			Id:          req.Id,
			Status:      req.Status,
			Num:         req.Num,
			ToAddress:   req.ToAddress,
			FromAddress: req.FromAddress,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil

}
