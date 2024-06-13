package setting

import (
	"context"
	"encoding/json"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSettingLogic {
	return &FindSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *FindSettingLogic) FindSetting(req *types.FindSettingReq) (resp *types.FindSettingResp, err error) {

	if !l.svcCtx.Config.WolfLampRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.WolfLampRpc.FindSetting(l.ctx, &wolflamp.FindSettingReq{Module: "platform_setting"})
	if err != nil {
		return nil, err
	}

	var platformSetting types.SettingInfo
	err = json.Unmarshal([]byte(data.JsonString), &platformSetting)
	if err != nil {
		return nil, err
	}

	return &types.FindSettingResp{
		Data: platformSetting,
	}, nil

}
