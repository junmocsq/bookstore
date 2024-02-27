package user

type Shelf struct {
	ID        int32  `json:"id"`
	Uid       int32  `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Name      string `gorm:"type:varchar(30);not null;default:'';comment:书架名称" json:"name"`
	Summary   string `gorm:"type:varchar(100);not null;default:'';comment:书架简介" json:"summary"`
	Cover     string `gorm:"type:varchar(100);not null;default:'';comment:书架封面" json:"cover"`
	Count     int32  `gorm:"type:int;not null;default:0;comment:书籍数量" json:"count"`
	Idx       int32  `gorm:"type:int;not null;default:0;comment:排序" json:"idx"`
	CreatedAt int64  `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (s *Shelf) TableName() string {
	return "u_shelves"
}
