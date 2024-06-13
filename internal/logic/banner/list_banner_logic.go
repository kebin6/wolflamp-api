package banner

import (
	"context"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBannerLogic {
	return &ListBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ListBannerLogic) ListBanner(req *types.ListBannerReq) (resp *types.ListBannerResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.ListBanner(l.ctx,
		&wolflamp.ListBannerReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Status:   req.Status,
			Endpoint: req.Endpoint,
			Module:   req.Module,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ListBannerResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		var file types.FileInfo
		if v.File != nil {
			file = types.FileInfo{
				Id:       v.File.Id,
				Name:     v.File.Name,
				Path:     v.File.Path,
				FileType: v.File.FileType,
			}
		}
		resp.Data.Data = append(resp.Data.Data,
			types.BannerInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Status:    v.Status,
				Endpoint:  v.Endpoint,
				Module:    v.Module,
				FileType:  v.FileType,
				Path:      v.Path,
				JumpUrl:   v.JumpUrl,
				File:      &file,
			})
	}
	return resp, nil

}
