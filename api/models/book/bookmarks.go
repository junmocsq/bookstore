package book

type Bookmark struct {
	ID        int32 `json:"id"`
	Bid       int32 `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Sid       int32 `gorm:"type:int;not null;default:0;comment:章节id" json:"sid"`
	Idx       int32 `gorm:"type:int;not null;default:0;comment:读到哪" json:"idx"`
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
}

func (b *Bookmark) TableName() string {
	return "b_bookmarks"
}
