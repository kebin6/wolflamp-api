package notice

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindNoticeLogic {
	return &FindNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *FindNoticeLogic) FindNotice(req *types.FindNoticeReq) (resp *types.FindNoticeResp, err error) {

	if !l.svcCtx.Config.WolfLampRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.WolfLampRpc.FindSetting(l.ctx, &wolflamp.FindSettingReq{Module: "platform_notice"})
	if err != nil {
		return nil, err
	}

	return &types.FindNoticeResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  i18n.Success,
			Data: data.JsonString,
		},
	}, nil

}
