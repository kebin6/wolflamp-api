package banner

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBannerLogic {
	return &FindBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *FindBannerLogic) FindBanner(req *types.FindBannerReq) (resp *types.FindBannerResp, err error) {

	if !l.svcCtx.Config.WolfLampRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.WolfLampRpc.FindBanner(l.ctx, &wolflamp.FindBannerReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.FindBannerResp{
		Data: types.BannerInfo{
			Status:    data.Status,
			Endpoint:  data.Endpoint,
			Module:    data.Module,
			FileType:  data.FileType,
			JumpUrl:   data.JumpUrl,
			Path:      data.Path,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			File: &types.FileInfo{
				Id:       data.File.Id,
				Name:     data.File.Name,
				FileType: data.File.FileType,
			},
		},
	}, nil

}
