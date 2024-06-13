package banner

import (
	"context"
	"github.com/kebin6/wolflamp-api/internal/svc"
	"github.com/kebin6/wolflamp-api/internal/types"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBannerLogic {
	return &UpdateBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateBannerLogic) UpdateBanner(req *types.UpdateBannerReq) (resp *types.BaseMsgResp, err error) {

	data, err := l.svcCtx.WolfLampRpc.UpdateBanner(l.ctx,
		&wolflamp.UpdateBannerReq{
			Id:       req.Id,
			Status:   req.Status,
			Endpoint: req.Endpoint,
			Module:   req.Module,
			FileType: req.FileType,
			JumpUrl:  req.JumpUrl,
			Path:     req.Path,
			FileId:   req.FileId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil

}
