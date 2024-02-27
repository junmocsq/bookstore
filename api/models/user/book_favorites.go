// 书籍收藏表
package user

type BookFavorite struct {
	ID        int64 `json:"id"`
	Uid       int32 `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Bid       int32 `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	CreatedAt int64 `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (bf *BookFavorite) TableName() string {
	return "u_book_favorites"
}
