package logic

import (
	"context"

	"github.com/tangbo/twatt/liyue/im/gateway/internal/svc"
	"github.com/tangbo/twatt/liyue/im/gateway/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GreetLogic) Greet(in *pb.GreetReq) (*pb.GreetResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GreetResp{}, nil
}
