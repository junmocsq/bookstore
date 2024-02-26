package book

type Chapter struct {
	ID      int32  `json:"id"`
	Bid     int32  `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Title   string `gorm:"type:string;size:50;not null;default:'';comment:大章名" json:"title"`
	Summary string `gorm:"type:string;size:50;not null;default:'';comment:介绍" json:"summary"`
}

func (c *Chapter) TableName() string {
	return "u_chapters"
}
