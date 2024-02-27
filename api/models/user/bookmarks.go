package user

import (
	"fmt"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Bookmark struct {
	ID        int32  `json:"id"`
	Bid       int32  `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Uid       int32  `gorm:"type:int;not null;default:0;comment:用户id" json:"uid"`
	Sid       int32  `gorm:"type:int;not null;default:0;comment:章节id" json:"sid"`
	Idx       int32  `gorm:"type:int;not null;default:0;comment:读到哪" json:"idx"`
	Summary   string `gorm:"type:string;size:100;not null;default:'';comment:介绍" json:"summary"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (b *Bookmark) TableName() string {
	return "u_bookmarks"
}

func (b *Bookmark) Tag(uid int32) string {
	return fmt.Sprintf("u_bookmarks_%d", uid)
}

func NewBookmark() *Bookmark {
	return &Bookmark{}
}

func (b *Bookmark) Add(uid, bid, sid, idx int32, summary string) int32 {
	var bookmark = Bookmark{Bid: bid, Uid: uid, Sid: sid, Idx: idx, Summary: summary}
	var db = common.GetDB()
	stmt := db.DryRun().Create(&bookmark).Statement
	_, err := db.SetTag(b.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&bookmark)
	if err != nil {
		logrus.WithField("model", "user_add").Error(err)
		return 0
	}
	return int32(bookmark.ID)
}

func (b *Bookmark) GetByBidAndSid(uid, bid, sid int32) []*Bookmark {
	var bookmarks []*Bookmark
	var db = common.GetDB()
	var stmt *gorm.Statement
	if sid > 0 {
		stmt = db.DryRun().Where("uid = ? and bid = ? and sid = ?", uid, bid, sid).Find(&bookmarks).Statement
	} else {
		stmt = db.DryRun().Where("uid = ? and bid = ?", uid, bid).Find(&bookmarks).Statement
	}
	err := db.SetTag(b.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&bookmarks)
	if err != nil {
		logrus.WithField("model", "bookmark_GetByBidAndSid").Error(err)
		return nil
	}
	return bookmarks
}

func (b *Bookmark) DeleteById(id, uid int32) int32 {
	var db = common.GetDB()
	stmt := db.DryRun().Delete(&Bookmark{}, id).Statement
	n, err := db.SetTag(b.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "bookmark_DeleteById").Error(err)
		return 0
	}
	return int32(n)
}
