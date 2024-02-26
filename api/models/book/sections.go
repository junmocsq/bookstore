package book

type Section struct {
	ID          int32  `json:"id"`
	Bid         int32  `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Cid         int32  `gorm:"type:int;not null;default:0;comment:大章id" json:"cid"`
	Title       string `gorm:"type:string;size:50;not null;default:'';comment:章节名" json:"title"`
	ContentId   int32  `gorm:"type:int;not null;default:0;comment:内容id" json:"content_id"`
	Bananas     int32  `gorm:"type:int;not null;default:0;comment:充值货币" json:"bananas"`
	Status      int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 上架 2 审查" json:"status"`
	PublishTime int64  `gorm:"type:bigint;not null;default:0;comment:发布时间" json:"publish_time"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (s *Section) TableName() string {
	return "b_sections"
}
