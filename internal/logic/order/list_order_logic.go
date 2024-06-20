package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/orderenum"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ListOrderLogic) ListOrder(req *types.ListOrderReq) (resp *types.ListOrderResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.ListOrder(l.ctx,
		&wolflamp.ListOrderReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Status:   req.Status,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ListOrderResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.OrderInfo{
				Id:          v.Id,
				Code:        v.Code,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Status:      v.Status,
				StatusDesc:  orderenum.NewOrderStatus(v.Status).Desc(),
				Type:        v.Type,
				Num:         v.Num,
				ToAddress:   v.ToAddress,
				FromAddress: v.FromAddress,
			})
	}
	return resp, nil

}
