package book

import (
	"errors"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Author struct {
	ID           int32  `json:"id"`
	Name         string `gorm:"type:varchar(50);not null;default:'';comment:作者名" json:"name"`
	Introduction string `gorm:"type:varchar(300);not null;default:'';comment:作者简介" json:"introduction"`
	Profile      string `gorm:"type:varchar(100);not null;default:'';comment:作者头像" json:"profile"`
	Country      string `gorm:"type:varchar(50);not null;default:'';comment:国家" json:"country"`
	CreatedAt    int64  `gorm:"autoCreateTime;type:int;not null;default:0;comment:创建时间" json:"created_at"`
}

func (a *Author) TableName() string {
	return "b_authors"
}

func NewAuthor() *Author {
	return &Author{}
}

func (a *Author) Tag() string {
	return "b_authors"
}

func (a *Author) Add(name, introduction, profile, country string) (int32, error) {
	if a.checkName(name) != nil {
		return 0, errors.New(tools.GetMsg(10003))
	}
	var author = Author{
		Name:         name,
		Introduction: introduction,
		Profile:      profile,
		Country:      country,
	}
	var db = common.GetDB()
	stmt := db.DryRun().
		Select("Name", "Introduction", "Profile", "Country", "CreatedAt").
		Create(&author).Statement
	_, err := db.SetTag(a.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&author)
	if err != nil {
		logrus.WithField("model", "author_Add").Error(err)
		return 0, err
	}
	return author.ID, nil
}

func (a *Author) checkName(name string) *Author {
	var author Author
	var db = common.GetDB()
	stmt := db.DryRun().Where("name =?", name).Find(&author).Statement
	err := db.SetTag(a.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&author)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "author_checkName").Error(err)
		}
		return nil
	}
	return &author
}

func (a *Author) GetById(id int32) *Author {
	var author Author
	var db = common.GetDB()
	stmt := db.DryRun().Where("id =?", id).Find(&author).Statement
	err := db.SetTag(a.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&author)
	if err != nil {
		logrus.WithField("model", "author_GetById").Error(err)
		return nil
	}
	return &author
}

func (a *Author) GetByIds(ids []int32) []*Author {
	var authors []*Author
	var db = common.GetDB()
	stmt := db.DryRun().Where("id in ?", ids).Find(&authors).Statement
	err := db.SetTag(a.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&authors)
	if err != nil {
		logrus.WithField("model", "author_GetByIds").Error(err)
		return nil
	}
	return authors
}

func (a *Author) UpdateAuthor(id int32, name, introduction, profile, country string) (int32, error) {
	existAuthor := a.checkName(name)
	if existAuthor != nil && existAuthor.ID != id {
		return 0, errors.New(tools.GetMsg(10003))
	}

	var db = common.GetDB()
	var author = Author{
		Name:         name,
		Introduction: introduction,
		Profile:      profile,
		Country:      country,
	}
	stmt := db.DryRun().Model(&author).Where("id =?", id).Updates(author).Statement
	n, err := db.SetTag(a.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "author_UpdateAuthor").Error(err)
		return 0, err
	}
	return int32(n), nil
}

func (a *Author) Search(name string, page, size int) []*Author {
	var db = common.GetDB()
	var authors []*Author
	var stmt *gorm.Statement
	if name != "" {
		stmt = db.DryRun().Where("name like ?", "%"+name+"%").
			Find(&authors).Limit(size).Offset((page - 1) * size).Order("id desc").Statement
	} else {
		stmt = db.DryRun().Find(&authors).
			Limit(size).Offset((page - 1) * size).Order("id desc").Statement
	}
	err := db.SetTag(a.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&authors)
	if err != nil {
		logrus.WithField("model", "author_Search").Error(err)
		return nil
	}
	return authors

}
