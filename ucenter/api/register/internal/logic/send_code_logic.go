package logic

import (
	"context"
	"errors"
	"time"

	"github.com/wushengyouya/coin_exchange/coin-common/tools"
	"github.com/wushengyouya/coin_exchange/grpc_common/ucenter/types/register"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER::"

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *register.CodeReq) (*register.NoRes, error) {
	// todo: add your logic here and delete this line
	// 收到手机号和国家标识
	// 生成验证码
	// 根据对应的国家和手机号调用对应的平台
	// 将验证码写入redis
	// 返回成功

	code := tools.Rand4Num()
	go func() {
		logx.Info("调用短信平台发送验证码成功。。。\n")
	}()
	logx.Infof("验证码为%s:%s\n", RegisterCacheKey+in.Phone, code)
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(c, RegisterCacheKey+in.Phone, code, 10*time.Minute)
	if err != nil {
		return nil, errors.New("验证码发送失败")
	}
	return &register.NoRes{}, nil
}
