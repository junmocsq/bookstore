// 阅读书籍时长表
package user

type ReadTime struct {
	ID       int32 `json:"id"`
	Uid      int32 `gorm:"uniqueIndex:uid_bid;type:int;not null;default:0;comment:用户id" json:"uid"`
	Bid      int32 `gorm:"uniqueIndex:uid_bid;type:int;not null;default:0;comment:书籍id" json:"bid"`
	ReadTime int64 `gorm:"type:int;not null;default:0;comment:最后阅读时间" json:"read_time"`
}

func (rt *ReadTime) TableName() string {
	return "u_read_times"
}
