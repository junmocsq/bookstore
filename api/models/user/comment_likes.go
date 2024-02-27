package user

type CommentLike struct {
	ID        int64 `json:"id"`
	Uid       int32 `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	CommentId int32 `gorm:"type:int;not null;default:0;comment:评论id" json:"comment_id"`
	CreatedAt int64 `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (cl *CommentLike) TableName() string {
	return "u_comment_likes"
}
