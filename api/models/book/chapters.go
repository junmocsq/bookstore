package book

import (
	"errors"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Chapter struct {
	ID      int32  `json:"id"`
	Bid     int32  `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Title   string `gorm:"type:string;size:50;not null;default:'';comment:大章名" json:"title"`
	Summary string `gorm:"type:string;size:50;not null;default:'';comment:介绍" json:"summary"`
}

func (c *Chapter) TableName() string {
	return "b_chapters"
}

func (c *Chapter) Tag(bid int32) string {
	return "u_chapters"
}

func NewChapter() *Chapter {
	return &Chapter{}
}

func (c *Chapter) Add(bid int32, title, summary string) (int32, error) {
	if title == "" {
		return 0, errors.New("大章名不能为空")

	}
	cs := c.GetChaptersByBid(bid)
	for _, v := range cs {
		if v.Title == title {
			return v.ID, nil
		}
	}
	var db = common.GetDB()
	var chapter = Chapter{Bid: bid, Title: title, Summary: summary}
	stmt := db.DryRun().Create(&chapter).Statement
	_, err := db.SetTag(c.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&chapter)
	if err != nil {
		logrus.WithField("model", "chapter_Add").Error(err)
		return 0, err
	}
	return chapter.ID, nil
}

func (c *Chapter) Update(id, bid int32, title, summary string) int32 {
	var db = common.GetDB()
	var chapter = Chapter{ID: id, Title: title, Summary: summary}
	stmt := db.DryRun().Model(&chapter).Updates(Chapter{
		Title:   title,
		Summary: summary,
	}).Statement
	n, err := db.SetTag(c.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "chapter_update").Error(err)
		return 0
	}
	return int32(n)
}

func (c *Chapter) GetChaptersByBid(bid int32) []Chapter {
	var chapters []Chapter
	var db = common.GetDB()

	stmt := db.DryRun().Where("bid = ?", bid).Find(&chapters).Statement
	err := db.SetTag(c.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&chapters)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "chapter_GetByBid").Error(err)
		}
		return nil
	}
	return chapters
}

func (c *Chapter) DeleteById(id, bid int32) error {
	sections := NewSection().GetSectionsByBid(bid)
	for _, v := range sections {
		if v.Cid == id {
			return errors.New("大章被使用，不能删除")
		}
	}
	var db = common.GetDB()
	stmt := db.DryRun().Where("id = ?", id).Delete(&Chapter{}).Statement
	_, err := db.SetTag(c.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "chapter_DeleteById").Error(err)
		return err
	}
	return nil
}
