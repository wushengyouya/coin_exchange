package model

import "database/sql"

type Member struct {
	Id                         int64         `gorm:"column:id"`
	AliNo                      string        `gorm:"column:ali_no"`
	QrCodeUrl                  string        `gorm:"column:qr_code_url"`
	AppealSuccessTimes         int64         `gorm:"column:appeal_success_times"`
	AppealTimes                int64         `gorm:"column:appeal_times"`
	ApplicationTime            int64         `gorm:"column:application_time"`
	Avatar                     string        `gorm:"column:avatar"`
	Bank                       string        `gorm:"column:bank"`
	Branch                     string        `gorm:"column:branch"`
	CardNo                     string        `gorm:"column:card_no"`
	CertifiedBusinessApplyTime int64         `gorm:"column:certified_business_apply_time"`
	CertifiedBusinessCheckTime sql.NullInt64 `gorm:"column:certified_business_check_time"`
	CertifiedBusinessStatus    sql.NullInt64 `gorm:"column:certified_business_status"`
	ChannelId                  int64         `gorm:"column:channel_id"`
	Email                      string        `gorm:"column:email"`
	FirstLevel                 int64         `gorm:"column:first_level"`
	GoogleDate                 int64         `gorm:"column:google_date"`
	GooleKey                   string        `gorm:"column:goole_key"`
	GoogleState                int64         `gorm:"column:google_state"`
	IdNumber                   string        `gorm:"column:id_number"`
	InviterId                  int64         `gorm:"column:inviter_id"`
	IsChannel                  int64         `gorm:"column:is_channel"`
	JyPassword                 string        `gorm:"column:jy_password"`
	LastLoginTime              int64         `gorm:"column:last_login_time"`
	City                       string        `gorm:"column:city"`
	Country                    string        `gorm:"column:country"`
	District                   string        `gorm:"column:district"`
	Province                   string        `gorm:"column:province"`
	LoginCount                 int64         `gorm:"column:login_count"`
	LoginLock                  int64         `gorm:"column:login_lock"`
	Margin                     string        `gorm:"column:margin"`
	MemberLevel                int64         `gorm:"column:member_level"`
	MobilePhone                string        `gorm:"column:mobile_phone"`
	Password                   string        `gorm:"column:password"`
	PromotionCode              string        `gorm:"column:promotion_code"`
	PublishAdvertise           int64         `gorm:"column:publish_advertise"`
	RealName                   string        `gorm:"column:real_name"`
	RealNameStatus             int64         `gorm:"column:real_name_status"`
	RegistrationTime           int64         `gorm:"column:registration_time"`
	Salt                       string        `gorm:"column:salt"`
	SecondLevel                int64         `gorm:"column:second_level"`
	SignInAbility              int64         `gorm:"column:sign_in_ability"`
	Status                     int64         `gorm:"column:status"`
	ThirdLevel                 int64         `gorm:"column:third_level"`
	Token                      string        `gorm:"column:token"`
	TokenExpireTime            int64         `gorm:"column:token_expire_time"`
	TransactionStatus          int64         `gorm:"column:transaction_status"`
	TransactionTime            int64         `gorm:"column:transaction_time"`
	Transactions               int64         `gorm:"column:transactions"`
	Username                   string        `gorm:"column:username"`
	QrWeCodeUrl                string        `gorm:"column:qr_we_code_url"`
	Wechat                     string        `gorm:"column:wechat"`
	Local                      string        `gorm:"column:local"`
	Integration                int64         `gorm:"column:integration"`
	MemberGradeId              int64         `gorm:"column:member_grade_id"`  // 等级id
	KysStatus                  int64         `gorm:"column:kys_status"`       // kyc等级
	GeneralizeTotal            int64         `gorm:"column:generalize_total"` // 注册赠送积分
	InviterParentId            int64         `gorm:"column:inviter_parent_id"`
	SuperPartner               string        `gorm:"column:super_partner"`
	KickFee                    float64       `gorm:"column:kick_fee"`
	Power                      float64       `gorm:"column:power"`      // 个人矿机算力（每日维护）
	TeamLevel                  int64         `gorm:"column:team_level"` // 团队人数（每日维护）
	TeamPower                  float64       `gorm:"column:team_power"` // 团队框架算力（每日维护）
	MemberLevelId              int64         `gorm:"column:member_level_id"`
}

func (*Member) TableName() string {
	return "member"
}

const (
	GENERAL = iota
	REALNAME
	IDENTIFICATION
)

const (
	NORMALPARTER = "0"
	SUPERPARTER  = "1"
	PSUPERPARTER = "2"
)

const (
	NORMAL = iota
	ILLEGAL
)

func (m *Member) FillSuperPartner(partner string) {
	if partner == "" {
		m.SuperPartner = NORMALPARTER
		m.Status = NORMAL
	} else {
		m.SuperPartner = partner
		m.Status = ILLEGAL
	}
}
