package book

import (
	"fmt"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/sirupsen/logrus"
)

type Book struct {
	ID              int32  `json:"id"`
	Aid             int32  `gorm:"type:int;not null;default:0;comment:作者id" json:"aid"`
	Title           string `gorm:"type:string;size:50;not null;default:'';comment:书名" json:"title"`
	Summary         string `gorm:"type:string;size:100;not null;default:'';comment:简介" json:"summary"`
	Cover           string `gorm:"type:string;size:100;not null;default:'';comment:封面" json:"cover"`
	Status          int8   `gorm:"type:tinyint;not null;default:1;comment:状态 1 上架 2 审查 3 下架" json:"status"`
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

func (b *Book) Add() (int32, error) {

	return 0, nil
}

func (b *Book) GetById(id int32) *Book {
	var book Book
	var db = common.GetDB()
	stmt := db.DryRun().Where("id =?", id).Find(&book).Statement
	err := db.SetTag(b.Tag(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&book)
	if err != nil {
		logrus.WithField("model", "book_GetById").Error(err)
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

func (b *Book) Search(title string, status int8, categoryId int32, page, size int) ([]*Book, error) {
	return nil, nil
}
