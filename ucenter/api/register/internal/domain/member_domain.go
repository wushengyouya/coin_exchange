package domain

import (
	"context"

	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/model"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/repo"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) *model.Member {
	// dao查询数据库库
	return nil
}

func NewMemberDomain() *MemberDomain {
	return &MemberDomain{}
}
