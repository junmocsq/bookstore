package book

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Section struct {
	ID          int32  `json:"id"`
	Bid         int32  `gorm:"type:int;not null;default:0;comment:书籍id" json:"bid"`
	Cid         int32  `gorm:"type:int;not null;default:0;comment:大章id" json:"cid"`
	Title       string `gorm:"type:string;size:50;not null;default:'';comment:章节名" json:"title"`
	ContentId   int32  `gorm:"type:int;not null;default:0;comment:内容id" json:"content_id"`
	Bananas     int32  `gorm:"type:int;not null;default:0;comment:付费货币" json:"bananas"`
	Status      int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 发布 2 定时发布 3 审查 4 下架" json:"status"`
	PublishTime int64  `gorm:"type:bigint;not null;default:0;comment:发布时间" json:"publish_time"`
	Idx         int32  `gorm:"type:int;not null;default:0;comment:章节排序" json:"idx"`
	Wordnum     int32  `gorm:"type:int;not null;default:0;comment:字数" json:"wordnum"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (s *Section) TableName() string {
	return "b_sections"
}

func (s *Section) Tag(bid int32) string {
	return fmt.Sprintf("b_sections_%d", bid)
}

func NewSection() *Section {
	return &Section{}
}

func (s *Section) Add(bid, cid int32, title string, content string, status int8, puhlistTime int64) (int32, error) {
	var db = common.GetDB()
	db.Begin()
	var cnt = Content{Content: content}
	stmt := db.DryRun().Create(&cnt).Statement
	_, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&cnt)
	if err != nil {
		logrus.WithField("model", "section_Add-content").Error(err)
		db.DB().Rollback()
		return 0, err
	}

	var section = Section{Bid: bid, Cid: cid, Title: title, Status: status,
		PublishTime: puhlistTime, ContentId: cnt.ID, Idx: s.GetMaxIdxByBid(bid) + 1, Wordnum: int32(utf8.RuneCountInString(content))}
	stmt = db.DryRun().Create(&section).Statement
	_, err = db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&section)
	if err != nil {
		logrus.WithField("model", "section_Add-section").Error(err)
		db.DB().Rollback()
		return 0, err
	}
	db.DB().Commit()

	// 修改最新章节
	newSection := s.GetMaxSectionByBid(bid)
	if newSection != nil {
		NewBook().UpdateSection(bid, newSection.ID, newSection.CreatedAt)
	}
	return section.ID, nil
}

func (s *Section) GetSectionsByBid(bid int32) []Section {
	var sections []Section
	var db = common.GetDB()
	var stmt = db.DryRun().Where("bid = ? ", bid).Find(&sections).Order("idx asc").Statement
	err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&sections)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "section_GetSectionsByBid").Error(err)
		}
		return nil
	}
	return sections
}

// GetMaxSectionByBid 获取最后章节
func (s *Section) GetMaxSectionByBid(bid int32) *Section {
	var section Section
	var db = common.GetDB()
	var stmt = db.DryRun().Where("bid=?", bid).Order("idx desc").Limit(1).Find(&section).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&section)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "section_GetMaxSectionByBid").Error(err)
		}
		return nil
	}
	return &section
}

// GetMaxIdxByBid 获取最大的idx
func (s *Section) GetMaxIdxByBid(bid int32) int32 {
	section := s.GetMaxSectionByBid(bid)
	if section == nil {
		return 0
	}
	return section.Idx
}

func (s *Section) GetById(id, bid int32) *Section {
	var section Section
	var db = common.GetDB()
	var stmt = db.DryRun().Where("id =?", id).Find(&section).Statement
	err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&section)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "section_GetById").Error(err)
		}
		return nil
	}
	return &section
}

func (s *Section) UpdateContent(id, bid int32, content string) error {
	sec := s.GetById(id, bid)
	if sec == nil {
		return errors.New("section not found")
	}
	var db = common.GetDB()
	var cnt = Content{ID: sec.ContentId, Content: content}
	stmt := db.DryRun().Model(&cnt).Updates(Content{Content: content}).Statement
	_, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "section_UpdateContent").Error(err)
		return err
	}
	return nil
}

func (s *Section) UpdatePublishTime(id, bid int32, publishTime int64) int32 {
	now := time.Now().Unix()
	if now > publishTime {
		logrus.WithField("model", "section_UpdatePublishTime").Error("发布时间不能早于当前时间")
		return 0
	}
	old := s.GetById(id, bid)
	if old == nil {
		return 0
	}
	if old.Status != 2 {
		logrus.WithField("model", "section_UpdatePublishTime-status").Error("章节状态不为定时状态，不能修改发布时间")
		return 0
	}
	var db = common.GetDB()
	var section = Section{ID: id, PublishTime: publishTime}
	stmt := db.DryRun().Model(&section).Updates(Section{
		PublishTime: publishTime,
	}).Statement
	n, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "section_UpdatePublishTime").Error(err)
		return 0
	}
	return int32(n)
}

func (s *Section) UpdateBananas(id, bid, bananas int32) int32 {
	var db = common.GetDB()
	var section = Section{ID: id, Bananas: bananas}
	stmt := db.DryRun().Model(&section).Updates(Section{
		Bananas: bananas,
	}).Statement
	n, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "section_UpdateBananas").Error(err)
		return 0
	}
	return int32(n)
}

func (s *Section) UpdateStatus(id, bid int32, status int8) int32 {
	var db = common.GetDB()
	var section = Section{ID: id, Status: status}
	stmt := db.DryRun().Model(&section).Updates(Section{
		Status: status,
	}).Statement
	n, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "section_UpdateStatus").Error(err)
		return 0
	}
	return int32(n)
}

func (s *Section) UpdateTitle(id, bid int32, title string) int32 {
	var db = common.GetDB()
	var section = Section{ID: id, Title: title}
	stmt := db.DryRun().Model(&section).Updates(Section{
		Title: title,
	}).Statement
	n, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "section_UpdateTitle").Error(err)
		return 0
	}
	return int32(n)
}

func (s *Section) UpdateCid(id, bid, cid int32) int32 {
	var db = common.GetDB()
	var section = Section{ID: id, Cid: cid}
	stmt := db.DryRun().Model(&section).Updates(map[string]interface{}{
		"cid": cid,
	}).Statement
	n, err := db.SetTag(s.Tag(bid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "section_UpdateCid").Error(err)
		return 0
	}
	return int32(n)
}
