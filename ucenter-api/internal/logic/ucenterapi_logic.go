package logic

import (
	"context"
	"time"

	"github.com/wushengyouya/coin_exchange/grpc_common/ucenter/types/register"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/svc"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = l.svcCtx.UCRegisterRpc.RegisterByPhone(c, &register.RegReq{})
	if err != nil {
		logx.Error("Register err: ", err)
	}

	return &types.Response{
		"注册成功",
	}, err
}
