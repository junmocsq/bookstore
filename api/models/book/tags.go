package book

type Tag struct {
	ID     int32  `json:"id"`
	Name   string `gorm:"type:string;size:20;not null;default:'';comment:分类名" json:"name"`
	Status int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 上架 2 下架" json:"status"`
}

func (t *Tag) TableName() string {
	return "b_tags"
}
