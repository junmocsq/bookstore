package book

import (
	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
)

type Category struct {
	ID     int32  `json:"id"`
	Name   string `gorm:"type:string;size:20;not null;default:'';comment:分类名" json:"name"`
	Pid    int32  `gorm:"type:int;not null;default:0;comment:父级ID" json:"pid"`
	Idx    int32  `gorm:"type:int;not null;default:0;comment:排序" json:"idx"`
	Status int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 上架 2 下架" json:"status"`
}

func (c *Category) TableName() string {
	return "b_categories"
}

func (c *Category) Tag() string {
	return "b_categories"
}

func NewCategory() *Category {
	return &Category{}
}

func (c *Category) Add(name string, pid int32, idx int32) (int32, error) {
	var category = Category{Name: name, Pid: pid, Idx: idx}
	var db = common.GetDB()
	stmt := db.DryRun().Select("Name", "Pid", "Idx").Create(&category).Statement
	_, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&category)
	if err != nil {
		logrus.WithField("model", "category_Add").Error(err)
		return 0, err
	}
	return category.ID, nil
}

func (c *Category) Update(id int32, name string, pid int32, idx int32, status int8) int32 {
	var db = common.GetDB()
	var category = Category{ID: id, Name: name, Pid: pid, Idx: idx, Status: status}
	stmt := db.DryRun().Model(&category).Updates(&category).Statement
	n, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "category_update").Error(err)
		return 0
	}
	return int32(n)
}

func (c *Category) GetAll() []*Category {
	var db = common.GetDB()
	var categories []*Category
	stmt := db.DryRun().Find(&categories).Statement
	err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&categories)
	if err != nil {
		logrus.WithField("model", "category_GetAll").Error(err)
		return nil
	}
	return categories
}

func (c *Category) GetById(id int32) *Category {
	var db = common.GetDB()
	var category Category
	stmt := db.DryRun().Where("id = ?", id).First(&category).Statement
	err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&category)
	if err != nil {
		logrus.WithField("model", "category_GetById").Error(err)
		return nil
	}
	return &category
}
