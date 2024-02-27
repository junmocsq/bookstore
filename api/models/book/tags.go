package book

import (
	"errors"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
)

type Tag struct {
	ID     int32  `json:"id"`
	Name   string `gorm:"uniqueIndex;type:string;size:20;not null;default:'';comment:分类名" json:"name"`
	Status int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 上架 2 下架" json:"status"`
}

func (t *Tag) TableName() string {
	return "b_tags"
}

func (t *Tag) Tag() string {
	return "b_tags"
}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Add(name string) (int32, error) {
	var db = common.GetDB()
	exsitsTag := t.GetByName(name)
	if exsitsTag != nil {
		return 0, errors.New("已存在标签")
	}
	var tag = Tag{Name: name}
	stmt := db.DryRun().Select("Name").Create(&tag).Statement
	_, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&tag)
	if err != nil {
		logrus.WithField("model", "tag_Add").Error(err)
		return 0, err
	}
	return tag.ID, nil
}

func (t *Tag) update(id int32, name string, status int8) int32 {
	var db = common.GetDB()
	var tag = Tag{ID: id, Name: name, Status: status}

	stmt := db.DryRun().Model(&tag).Updates(&tag).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "tag_update").Error(err)
		return 0
	}
	return int32(n)
}

func (t *Tag) UpdateName(id int32, name string) int32 {
	tag := t.GetByName(name)
	if tag != nil {
		return 0
	}
	return t.update(id, name, 0)
}

func (t *Tag) UpdateStatus(id int32, status int8) int32 {
	return t.update(id, "", status)
}

func (t *Tag) GetById(id int32) *Tag {
	var db = common.GetDB()
	var tag Tag
	stmt := db.DryRun().Where("id = ?", id).First(&tag).Statement
	err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&tag)
	if err != nil {
		logrus.WithField("model", "tag_GetById").Error(err)
		return nil
	}
	return &tag
}

func (t *Tag) GetByName(name string) *Tag {
	var db = common.GetDB()
	var tag Tag
	stmt := db.DryRun().Where("name = ?", name).First(&tag).Statement
	err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&tag)
	if err != nil {
		logrus.WithField("model", "tag_GetByNamed").Error(err)
		return nil
	}
	return &tag
}

func (t *Tag) GetAll() []*Tag {
	var db = common.GetDB()
	var tags []*Tag
	stmt := db.DryRun().Find(&tags).Statement
	err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&tags)
	if err != nil {
		logrus.WithField("model", "tag_GetAll").Error(err)
		return nil
	}
	return tags
}
