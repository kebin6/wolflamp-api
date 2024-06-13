package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderLogic {
	return &FindOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *FindOrderLogic) FindOrder(req *types.FindOrderReq) (resp *types.FindOrderResp, err error) {

	if !l.svcCtx.Config.WolfLampRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.WolfLampRpc.FindOrder(l.ctx, &wolflamp.FindOrderReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.FindOrderResp{
		Data: types.OrderInfo{
			Id:            data.Id,
			Status:        data.Status,
			CreatedAt:     data.CreatedAt,
			UpdatedAt:     data.UpdatedAt,
			Type:          data.Type,
			Num:           data.Num,
			ToAddress:     data.ToAddress,
			FromAddress:   data.FromAddress,
			TransactionId: data.TransactionId,
		},
	}, nil

}
