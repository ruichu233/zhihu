package logic

import (
	"context"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeNumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeNumLogic {
	return &LikeNumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeNumLogic) LikeNum(req *types.LikeNumRequest) (resp *types.LikeNumResponse, err error) {
	getPostLikeCountResponse, err := l.svcCtx.LikeRPC.GetPostLikeCount(l.ctx, &like.GetPostLikeCountRequest{
		BizId: req.BizId,
		ObjId: req.ObjId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.LikeNumResponse{
		LikeNum: getPostLikeCountResponse.Count,
	}
	return
}
