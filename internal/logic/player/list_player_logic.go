package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPlayerLogic {
	return &ListPlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ListPlayerLogic) ListPlayer(req *types.ListPlayerReq) (resp *types.ListPlayerResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.ListPlayer(l.ctx,
		&wolflamp.ListPlayerReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Status:      req.Status,
			InviteCode:  req.InviteCode,
			InvitedCode: req.InvitedCode,
			Email:       req.Email,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ListPlayerResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	findLogic := NewFindPlayerLogic(l.ctx, l.svcCtx)
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, findLogic.Po2Vo(v))
	}
	return resp, nil

}
