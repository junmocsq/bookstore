package services

import "github.com/junmocsq/bookstore/api/models/book"

type sCategory struct{}

func NewCategory() *sCategory {
	return &sCategory{}
}

func (c *sCategory) Add(name string, pid int32, idx int32) (int32, error) {
	return book.NewCategory().Add(name, pid, idx)
}

func (c *sCategory) Update(id int32, name string, pid int32, idx int32, status int8) int32 {
	return book.NewCategory().Update(id, name, pid, idx, status)
}

func (c *sCategory) Mappings() map[int32]string {
	categories := book.NewCategory().GetAll()
	m := map[int32]string{}
	for _, v := range categories {
		m[v.ID] = v.Name
	}
	return m
}

type SCategory struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Childs []*SCategory
}

func (c *sCategory) FormatCategory(status int8) []*SCategory {
	categories := book.NewCategory().GetAll()
	m := map[int32]*SCategory{}
	for _, v := range categories {
		var temp = SCategory{
			ID:   v.ID,
			Name: v.Name,
		}
		m[v.ID] = &temp
	}
	var res []*SCategory
	for _, v := range categories {
		if status > 0 && v.Status != status {
			continue
		}
		if v.Pid == 0 {
			res = append(res, m[v.ID])
		} else {
			m[v.Pid].Childs = append(m[v.Pid].Childs, m[v.ID])
		}
	}
	return res
}
