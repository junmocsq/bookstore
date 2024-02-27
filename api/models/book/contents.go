package book

import (
	"fmt"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
)

type Content struct {
	ID      int32  `json:"id"`
	Content string `gorm:"type:text;comment:内容" json:"content"`
}

func (c *Content) TableName() string {
	return "b_contents"
}

func (c *Content) Tag() string {
	return "b_contents"
}

func (c *Content) Key(id int32) string {
	return fmt.Sprintf("_%d", id)
}

func NewContent() *Content {
	return &Content{}
}

func (c *Content) Add(content string) (int32, error) {
	var db = common.GetDB()

	var cnt = Content{Content: content}
	stmt := db.DryRun().Create(&cnt).Statement
	_, err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&cnt)
	if err != nil {
		logrus.WithField("model", "content_Add").Error(err)
		return 0, err
	}
	return cnt.ID, nil
}

func (c *Content) GetById(id int32) *Content {
	var db = common.GetDB()
	var cnt Content
	stmt := db.DryRun().Where("id = ?", id).First(&cnt).Statement
	err := db.SetTag(c.Tag()).SetKey(c.Key(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&cnt)
	if err != nil {
		logrus.WithField("model", "content_GetById").Error(err)
		return nil
	}
	return &cnt
}

func (c *Content) Update(id int32, content string) int32 {
	var db = common.GetDB()
	var cnt = Content{ID: id, Content: content}
	stmt := db.DryRun().Model(&cnt).Updates(map[string]interface{}{"content": content}).Statement
	n, err := db.SetTag(c.Tag()).SetKey(c.Key(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "content_update").Error(err)
		return 0
	}
	return int32(n)
}
