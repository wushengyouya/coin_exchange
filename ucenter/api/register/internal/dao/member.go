package dao

import (
	"context"

	"github.com/wushengyouya/coin_exchange/coin-common/msdb"
	"github.com/wushengyouya/coin_exchange/coin-common/msdb/gorms"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/model"
	"gorm.io/gorm"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Seesion(ctx)
	err = session.Model(&model.Member{}).Where("mobile_phone = ?", phone).Limit(1).Take(&mem).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}
func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.DB),
	}
}
