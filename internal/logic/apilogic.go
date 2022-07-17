package logic

import (
	"context"

	"github.com/logan-liang/api/internal/svc"
	"github.com/logan-liang/api/internal/types"

	"github.com/logan-liang/rpc/demo"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.TestRpc.Ping(l.ctx, &demo.Request{
		Ping: "1",
	})

	return &types.Response{
		Message: "212121",
	}, err
}
