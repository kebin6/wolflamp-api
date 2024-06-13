package notice

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateNoticeLogic) UpdateNotice(req *types.UpdateNoticeReq) (resp *types.BaseMsgResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.UpdateSetting(l.ctx,
		&wolflamp.UpdateSettingReq{
			Module:     "platform_notice",
			JsonString: req.Notice,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil

}
