package user

type SectionSubscription struct {
	ID        int32 `json:"id"`
	Uid       int32 `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Bid       int32 `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Sid       int32 `gorm:"type:int;not null;default:0;comment:章节id" json:"sid"`
	Bananas   int32 `gorm:"type:int;not null;default:0;comment:消耗金币数" json:"bananas"`
	CreatedAt int64 `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (ss *SectionSubscription) TableName() string {
	return "u_section_subscriptions"
}
