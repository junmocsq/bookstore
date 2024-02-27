package user

type Comment struct {
	ID        int32  `json:"id"`
	Uid       int32  `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Bid       int32  `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Sid       int32  `gorm:"type:int;not null;default:0;comment:章节id" json:"sid"`
	ReplyId   int32  `gorm:"type:int;not null;default:0;comment:回复id" json:"reply_id"`
	TopId     int32  `gorm:"type:int;not null;default:0;comment:顶级id" json:"top_id"`
	Content   string `gorm:"type:varchar(500);not null;default:'';comment:评论内容" json:"content"`
	Likes     int32  `gorm:"type:int;not null;default:0;comment:点赞数" json:"likes"`
	Children  int32  `gorm:"type:int;not null;default:0;comment:子评论数" json:"children"`
	CreatedAt int64  `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (c *Comment) TableName() string {
	return "u_comments"
}
