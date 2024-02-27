package user

type RechargeRecord struct {
	ID           int32  `json:"id"`
	Uid          int32  `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Channel      int8   `gorm:"type:tinyint;not null;default:0;comment:1 系统充值 2 微信 3 支付宝 4 苹果" json:"channel"`
	Amount       int32  `gorm:"type:int;not null;default:0;comment:充值金额 单位分" json:"amount"`
	Status       int8   `gorm:"type:tinyint;not null;default:0;comment:0 未支付 1 已支付" json:"status"`
	OrderId      string `gorm:"type:varchar(64);not null;default:'';comment:订单号" json:"order_id"`
	ThirdOrderId string `gorm:"type:varchar(64);not null;default:'';comment:第三方订单号" json:"third_order_id"`
	UpdatedAt    int64  `gorm:"type:int;not null;default:0;comment:更新时间" json:"updated_at"`
	CreatedAt    int64  `gorm:"type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (rr *RechargeRecord) TableName() string {
	return "u_recharge_records"
}
