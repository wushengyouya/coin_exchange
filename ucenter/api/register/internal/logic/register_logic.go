package logic

import (
	"context"

	"github.com/wushengyouya/coin_exchange/grpc_common/ucenter/types/register"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line
	logx.Info("ucenter register by phone call ....")
	return &register.RegRes{}, nil
}
