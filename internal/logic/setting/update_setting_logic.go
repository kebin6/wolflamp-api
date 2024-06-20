package setting

import (
	"context"
	"encoding/json"
	"github.com/kebin6/wolflamp-rpc/common/enum"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSettingLogic {
	return &UpdateSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateSettingLogic) UpdateSetting(req *types.UpdateSettingReq) (resp *types.BaseMsgResp, err error) {

	jsonString, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	data, err := l.svcCtx.WolfLampRpc.UpdateSetting(l.ctx,
		&wolflamp.UpdateSettingReq{
			Module:     enum.PlatformSetting.Val(),
			JsonString: string(jsonString),
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil

}
