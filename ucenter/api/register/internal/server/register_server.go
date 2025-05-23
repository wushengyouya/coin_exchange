// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: register.proto

package server

import (
	"context"

	"github.com/wushengyouya/coin_exchange/grpc_common/ucenter/types/register"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/logic"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/svc"
)

type RegisterServer struct {
	svcCtx *svc.ServiceContext
	register.UnimplementedRegisterServer
}

func NewRegisterServer(svcCtx *svc.ServiceContext) *RegisterServer {
	return &RegisterServer{
		svcCtx: svcCtx,
	}
}

func (s *RegisterServer) RegisterByPhone(ctx context.Context, in *register.RegReq) (*register.RegRes, error) {
	l := logic.NewRegisterByPhoneLogic(ctx, s.svcCtx)
	return l.RegisterByPhone(in)
}

func (s *RegisterServer) SendCode(ctx context.Context, in *register.CodeReq) (*register.NoRes, error) {
	l := logic.NewSendCodeLogic(ctx, s.svcCtx)
	return l.SendCode(in)
}
