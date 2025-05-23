package repo

import (
	"context"

	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/model"
)

type MemberRepo interface {
	FindByPhone(ctx context.Context, phone string) (*model.Member, error)
}
