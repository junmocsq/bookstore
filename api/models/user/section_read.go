package user

type SectionRead struct {
	ID       int32 `json:"id"`
	Uid      int32 `gorm:"uniqueIndex:uid_bid_sid;type:int;not null;default:0;comment:用户id" json:"uid"`
	Bid      int32 `gorm:"uniqueIndex:uid_bid_sid;type:int;not null;default:0;comment:书籍id" json:"bid"`
	Sid      int32 `gorm:"uniqueIndex:uid_bid_sid;type:int;not null;default:0;comment:章节id" json:"sid"`
	Idx      int32 `gorm:"type:int;not null;default:0;comment:读到哪" json:"idx"`
	ReadTime int64 `gorm:"type:int;not null;default:0;comment:最后阅读时间" json:"read_time"`
}

func (sr *SectionRead) TableName() string {
	return "u_section_read_records"
}
