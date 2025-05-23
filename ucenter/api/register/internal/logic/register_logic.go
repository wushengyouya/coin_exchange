package logic

import (
	"context"
	"errors"

	"github.com/wushengyouya/coin_exchange/grpc_common/ucenter/types/register"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/domain"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDoamin *domain.CaptchaDomain
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDoamin: domain.NewCaptchaDomain(),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line
	// 1.先检验人机是否通过
	isVerify := l.CaptchaDoamin.Verify(
		in.Captcha.Server, l.svcCtx.Config.Captcha.Vid, l.svcCtx.Config.Captcha.Key, in.Captcha.Token, 2, in.Ip,
	)
	if !isVerify {
		return nil, errors.New("人机校验不通过")
	}
	logx.Info("人机校验通过")
	// 2.校验验证码
	redisValue := ""
	err := l.svcCtx.Cache.GetCtx(context.Background(), RegisterCacheKey+in.Phone, &redisValue)
	if err != nil {
		return nil, errors.New("验证码错误")
	}
	if in.Code != redisValue {
		return nil, errors.New("验证码输入错误")
	}

	// 3.验证码通过进行注册
	logx.Info("ucenter register by phone call ....")
	return &register.RegRes{}, nil
}
