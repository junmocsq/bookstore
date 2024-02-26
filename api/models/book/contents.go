package book

type Content struct {
	ID      int32  `json:"id"`
	Content string `gorm:"type:text;comment:内容" json:"content"`
}

func (c *Content) TableName() string {
	return "b_contents"
}
