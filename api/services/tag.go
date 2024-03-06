package services

import "github.com/junmocsq/bookstore/api/models/book"

type sTag struct {
}

type STag struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func NewTag() *sTag {
	return &sTag{}
}

func (t *sTag) Add(name string) (int32, error) {
	return book.NewTag().Add(name)
}

func (t *sTag) Mappings() map[int32]book.Tag {
	tags := book.NewTag().GetAll()
	m := map[int32]book.Tag{}
	for _, v := range tags {
		m[v.ID] = v
	}
	return m
}

func (t *sTag) UpdateName(id int32, name string) int32 {
	return book.NewTag().UpdateName(id, name)
}

func (t *sTag) UpdateStatus(id int32, status int8) int32 {
	return book.NewTag().UpdateStatus(id, status)
}

func (t *sTag) GetTagsByIds(ids []int32) []book.Tag {
	all := t.Mappings()
	var tags []book.Tag
	for _, id := range ids {
		if temp, ok := all[id]; ok {
			tags = append(tags, temp)
		}
	}
	return tags
}

func (t *sTag) GetTags(status int8) []STag {
	stags := book.NewTag().GetAll()
	var stagsArr []STag
	for _, tag := range stags {
		if status > 0 && tag.Status != status {
			continue
		}
		stagsArr = append(stagsArr, STag{ID: tag.ID, Name: tag.Name})
	}
	return stagsArr
}
