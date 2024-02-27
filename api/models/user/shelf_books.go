package user

type ShelfBook struct {
	ID        int32 `json:"id"`
	Uid       int32 `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Bid       int32 `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	ShelfId   int32 `gorm:"type:int;not null;default:0;comment:书架id" json:"shelf_id"`
	Idx       int32 `gorm:"type:int;not null;default:0;comment:排序" json:"idx"`
	CreatedAt int64 `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (sb *ShelfBook) TableName() string {
	return "u_shelf_books"
}
