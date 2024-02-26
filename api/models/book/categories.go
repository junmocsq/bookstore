package book

type Category struct {
	ID   int32  `json:"id"`
	Name string `gorm:"type:string;size:20;not null;default:'';comment:分类名" json:"name"`
	Pid  int32  `gorm:"type:int;not null;default:0;comment:父级ID" json:"pid"`
	Idx  int32  `gorm:"type:int;not null;default:0;comment:排序" json:"idx"`
}

func (c *Category) TableName() string {
	return "b_categories"
}
