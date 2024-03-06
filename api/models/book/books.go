package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Book struct {
	ID              int32  `json:"id"`
	Aid             int32  `gorm:"type:int;not null;default:0;comment:作者id" json:"aid"`
	Title           string `gorm:"type:string;size:50;not null;default:'';comment:书名" json:"title"`
	Summary         string `gorm:"type:string;size:100;not null;default:'';comment:简介" json:"summary"`
	Cover           string `gorm:"type:string;size:100;not null;default:'';comment:封面" json:"cover"`
	Status          int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 上架 2 审查 3 下架" json:"status"`
	Process         int8   `gorm:"type:tinyint;not null;default:1;comment:进度 1 连载 2 完结 3 停更" json:"process"`
	IsPay           int8   `gorm:"type:tinyint;not null;default:2;comment:是否付费 1 是 2 否" json:"is_pay"`
	CategoryId      int32  `gorm:"type:int;not null;default:0;comment:分类id" json:"category_id"`
	TagIds          string `gorm:"type:json;comment:标签id" json:"tag_ids"`
	LastSectionId   int32  `gorm:"type:int;not null;default:0;comment:最后章节id" json:"last_section_id"`
	LastSectionTime int64  `gorm:"type:bigint;not null;default:0;comment:最后章节时间" json:"last_section_time"`
	SectionNum      int32  `gorm:"type:int;not null;default:0;comment:章节数量" json:"section_num"`
	Favorites       int32  `gorm:"type:int;not null;default:0;comment:收藏数" json:"favorites"`
	Likes           int32  `gorm:"type:int;not null;default:0;comment:点赞数" json:"likes"`
	Comments        int32  `gorm:"type:int;not null;default:0;comment:评论数" json:"comments"`
	Apples          int32  `gorm:"type:int;not null;default:0;comment:虚拟币" json:"apples"`
	Clicks          int32  `gorm:"type:int;not null;default:0;comment:点击数" json:"clicks"`
	Hot             int32  `gorm:"type:int;not null;default:0;comment:热度 随时间变化" json:"hot"`
	Popular         int64  `gorm:"type:int;not null;default:0;comment:人气 累加" json:"popular"`
	UpdatedAt       int64  `json:"updated_at"`
	CreatedAt       int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (b *Book) TableName() string {
	return "b_books"
}

func (b *Book) Tag(id int32) string {
	return fmt.Sprintf("b_books_%d", id)
}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) Add(aid int32, title, summary, cover string, categoryId int32, tagIds []int32) (int32, error) {
	tagIdsStr, err := json.Marshal(tagIds)
	if err != nil {
		logrus.WithField("model", "book_Add-json").Error(err)
		return 0, err
	}
	if b.checkAuthorTitle(aid, title) != nil {
		return 0, errors.New(tools.GetMsg(10002))
	}
	var book = Book{
		Aid:        aid,
		Title:      title,
		Summary:    summary,
		Cover:      cover,
		CategoryId: categoryId,
		TagIds:     string(tagIdsStr),
	}
	var db = common.GetDB()
	stmt := db.DryRun().
		Select("Aid", "Title", "Summary", "Cover", "CategoryId", "TagIds", "CreatedAt").
		Create(&book).Statement
	_, err = db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&book)
	if err != nil {
		logrus.WithField("model", "book_Add").Error(err)
		return 0, err
	}
	return book.ID, nil
}

func (b *Book) GetById(id int32, nocache ...bool) *Book {
	var book Book
	var db = common.GetDB()
	stmt := db.DryRun().Where("id =?", id).Find(&book).Statement
	var err error
	if len(nocache) > 0 {
		err = db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&book)
	} else {
		err = db.SetTag(b.Tag(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&book)
	}
	if err != nil {
		logrus.WithField("model", "book_GetById").Error(err)
		return nil
	}
	return &book
}

func (b *Book) checkAuthorTitle(aid int32, title string) *Book {
	var book Book
	var db = common.GetDB()
	stmt := db.DryRun().Where("aid =? and title =?", aid, title).Find(&book).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&book)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "book_checkAuthorTitle").Error(err)
		}
		return nil
	}
	return &book
}
func (b *Book) GetByAid(aid int32, page, size int) []*Book {
	var db = common.GetDB()
	var books []*Book
	stmt := db.DryRun().Where("aid =?", aid).Find(&books).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&books)
	if err != nil {
		logrus.WithField("model", "book_GetByAid").Error(err)
		return nil
	}
	return books
}

func (b *Book) Search(title string, status int8,
	aid, categoryId int32, tagsIds []int32, page, size int) []*Book {
	var db = common.GetDB()
	var books []*Book
	gormDB := db.DryRun()
	if title != "" {
		gormDB = gormDB.Where("title like ?", "%"+title+"%")
	}
	if status != 0 {
		gormDB = gormDB.Where("status =?", status)
	}
	if aid != 0 {
		gormDB = gormDB.Where("aid =?", aid)
	}
	if categoryId != 0 {
		gormDB = gormDB.Where("category_id =?", categoryId)
	}
	for _, tagId := range tagsIds {
		gormDB = gormDB.Where("json_contains(tag_ids,?)", tagId)
	}
	stmt := gormDB.Offset((page - 1) * size).Limit(size).Find(&books).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&books)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "book_Search").Error(err)
		}
		return nil
	}
	return books
}

func (b *Book) UpdateBook(id int32, title, summary,
	cover string, categoryId int32, tagIds []int32) (int32, error) {
	if title != "" {
		abook := b.GetById(id)
		existBook := b.checkAuthorTitle(abook.Aid, title)
		if existBook != nil && existBook.ID != id {
			return 0, errors.New(tools.GetMsg(10002))
		}
	}
	tagIdsStr, err := json.Marshal(tagIds)
	if err != nil {
		logrus.WithField("model", "book_UpdateBook-json").Error(err)
		return 0, err
	}
	var db = common.GetDB()
	var book = Book{
		Title:      title,
		Summary:    summary,
		Cover:      cover,
		CategoryId: categoryId,
		TagIds:     string(tagIdsStr),
	}
	stmt := db.DryRun().Model(&book).Where("id =?", id).Updates(book).Statement
	n, err := db.SetTag(b.Tag(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "book_UpdateBook").Error(err)
		return 0, err
	}
	return int32(n), nil
}

func (b *Book) UpdateStatus(id int32, status int8) (int32, error) {
	var db = common.GetDB()
	stmt := db.DryRun().Model(&Book{}).Where("id =?", id).Update("status", status).Statement
	n, err := db.SetTag(b.Tag(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "book_UpdateStatus").Error(err)
		return 0, err
	}
	return int32(n), nil
}

func (b *Book) UpdateSection(id, sectionId int32, sectionTime int64) (int32, error) {
	var db = common.GetDB()
	stmt := db.DryRun().Model(&Book{}).Where("id =?", id).
		Updates(map[string]interface{}{"last_section_id": sectionId, "last_section_time": sectionTime}).Statement
	n, err := db.SetTag(b.Tag(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "book_UpdateSection").Error(err)
		return 0, err
	}
	return int32(n), nil
}

func (b *Book) incrementNum(id, num int32, field string) (int32, error) {
	if !slices.Contains([]string{"popular", "favorites", "likes",
		"comments", "apples", "clicks", "hot"}, field) {
		return 0, errors.New(tools.GetMsg(10004))
	}
	var db = common.GetDB()
	stmt := db.DryRun().Model(&Book{}).Where("id =?", id).
		Update(field, gorm.Expr(field+" + ?", num)).Statement
	n, err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "book_incrementNum").Error(err)
		return 0, err
	}
	return int32(n), nil
}

func (b *Book) IncrementPopular(id int32, num int64) (int32, error) {
	return b.incrementNum(id, int32(num), "popular")
}
func (b *Book) IncrementFavorites(id int32, num int32) (int32, error) {
	return b.incrementNum(id, num, "favorites")
}
func (b *Book) IncrementLikes(id int32, num int32) (int32, error) {
	return b.incrementNum(id, num, "likes")
}
func (b *Book) IncrementComments(id int32, num int32) (int32, error) {
	return b.incrementNum(id, num, "comments")
}
func (b *Book) IncrementApples(id int32, num int32) (int32, error) {
	return b.incrementNum(id, num, "apples")
}
func (b *Book) IncrementClicks(id int32, num int32) (int32, error) {
	return b.incrementNum(id, num, "clicks")
}
func (b *Book) IncrementHot(id int32, num int32) (int32, error) {
	return b.incrementNum(id, num, "hot")
}

func (b *Book) updateNum(id, num int32, field string) (int32, error) {
	if !slices.Contains([]string{"popular", "favorites", "likes",
		"comments", "apples", "clicks", "hot", "section_num"}, field) {
		return 0, errors.New(tools.GetMsg(10004))
	}
	var db = common.GetDB()
	stmt := db.DryRun().Model(&Book{}).Where("id =?", id).
		Update(field, num).Statement
	n, err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.WithField("model", "book_UpdateNum").Error(err)
		return 0, err
	}
	return int32(n), nil
}
func (b *Book) UpdateSectionNum(id, num int32) (int32, error) {
	return b.updateNum(id, num, "section_num")
}

func (b *Book) GetAllByField(page, size int, field string) []Book {
	var db = common.GetDB()
	var books []Book
	stmt := db.DryRun().Select("id", field).Offset((page - 1) * size).Limit(size).Find(&books).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&books)
	if err != nil {
		logrus.WithField("model", "book_GetAllByField").Error(err)
		return nil
	}
	return books
}

func (b *Book) GetBookByName(name string) *Book {
	var book Book
	var db = common.GetDB()
	stmt := db.DryRun().Where("title =?", name).Find(&book).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&book)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("model", "book_GetBookByName").Error(err)
		}
		return nil
	}
	return &book
}
