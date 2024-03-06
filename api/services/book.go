package services

import (
	"encoding/json"

	"github.com/junmocsq/bookstore/api/models/book"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/sirupsen/logrus"
)

type SBook struct {
	*book.Book
	StatusStr    string        `json:"status_str"`
	ProcessStr   string        `json:"process_str"`
	IsPayStr     string        `json:"is_pay_str"`
	CategoryStr  string        `json:"category_str"`
	TagIdsArr    []book.Tag    `json:"tag_ids_arr"`
	UpdatedAtStr string        `json:"updated_at_str"`
	CreatedAtStr string        `json:"created_at_str"`
	LastSection  *book.Section `json:"last_section"` // 最新章节
}

type sBook struct {
}

func NewBook() *sBook {
	return &sBook{}
}

func (b *sBook) Add(aid int32, title, summary, cover string, categoryId int32, tagIds []int32) (int32, error) {
	return book.NewBook().Add(aid, title, summary, cover, categoryId, tagIds)
}

func (b *sBook) Get(id int32) (*SBook, error) {
	var bk = book.NewBook().GetById(id)
	if bk == nil {
		return nil, nil
	}
	return b.formatBook(bk), nil
}

func (b *sBook) formatBook(bk *book.Book) *SBook {
	var tag_ids_str = bk.TagIds
	var tagIds []int32
	if tag_ids_str != "" {
		err := json.Unmarshal([]byte(tag_ids_str), &tagIds)
		if err != nil {
			logrus.WithField("services", "book_formatBook-JsonUnmarshal").Error(err)
			return nil
		}
	}
	return &SBook{
		Book:         bk,
		StatusStr:    tools.Mapping("book_status", bk.Status),
		ProcessStr:   tools.Mapping("book_process", bk.Process),
		IsPayStr:     tools.Mapping("book_is_pay", bk.IsPay),
		CategoryStr:  NewCategory().Mappings()[bk.CategoryId],
		TagIdsArr:    NewTag().GetTagsByIds(tagIds),
		UpdatedAtStr: tools.PublishTime2Read(bk.UpdatedAt),
		CreatedAtStr: tools.PublishTime2Read(bk.CreatedAt),
		LastSection:  book.NewSection().GetById(bk.LastSectionId, bk.ID),
	}
}
