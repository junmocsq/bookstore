package services

import (
	"github.com/junmocsq/bookstore/api/models/book"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/sirupsen/logrus"
)

type SSection struct {
	book.Section
	Chapter        book.Chapter `json:"chapter"`
	StatusStr      string       `json:"status_str"`
	PublishTimeStr string       `json:"publish_time_str"`
	UpdatedAtStr   string       `json:"updated_at_str"`
	CreatedAtStr   string       `json:"created_at_str"`
	Content        string       `json:"content"`
}

type sSection struct{}

func NewSection() *sSection {
	return &sSection{}
}

func (s *sSection) Add(bid int32, title, content string, idx int32) (int32, error) {
	return 0, nil
}

func (s *sSection) GetSectionByBidAndSid(bid, sid int32) *SSection {
	chapters := book.NewChapter().GetChaptersByBid(bid)
	section := book.NewSection().GetById(sid, bid)

	return s.formatSection(section, chapters, true)
}

func (s *sSection) GetSectionsByBid(bid int32, isContent bool) []SSection {
	chapters := book.NewChapter().GetChaptersByBid(bid)
	sections := book.NewSection().GetSectionsByBid(bid)
	logrus.Error(sections)
	var sArr []SSection
	for _, section := range sections {
		sArr = append(sArr, *s.formatSection(&section, chapters, false))
	}
	return sArr
}

func (s *sSection) formatSection(section *book.Section, chapters []book.Chapter, isContent bool) *SSection {
	var ss SSection
	for _, c := range chapters {
		if c.ID == section.Cid {
			ss.Chapter = c
			break
		}
	}
	ss.Section = *section
	ss.StatusStr = tools.Mapping("section_status", section.Status)
	ss.PublishTimeStr = tools.Time2Read(section.PublishTime)
	ss.UpdatedAtStr = tools.PublishTime2Read(section.UpdatedAt)
	ss.CreatedAtStr = tools.PublishTime2Read(section.CreatedAt)
	if isContent {
		var content = book.NewContent().GetById(section.ContentId)
		if content != nil {
			ss.Content = content.Content
		}
	}
	return &ss
}
