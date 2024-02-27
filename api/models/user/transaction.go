package user

type Transaction struct {
	ID           int32  `json:"id"`
	Uid          int32  `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Category     int8   `gorm:"type:tinyint;not null;default:0;comment:流水分类 1 充值 2 订阅 3 登录 4 签到" json:"category"`
	IncomeStatus int8   `gorm:"type:tinyint;not null;default:1;comment:收入状态 1 收入 2 支出" json:"income_status"`
	Bananas      int32  `gorm:"type:int;not null;default:0;comment:货币变化数" json:"bananas"`
	Apples       int32  `gorm:"type:int;not null;default:0;comment:虚拟币变化数" json:"apples"`
	Mark         string `gorm:"type:varchar(500);not null;default:'';comment:备注" json:"mark"`
	CreatedAt    int64  `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (Transaction) TableName() string {
	return "u_transactions"
}
