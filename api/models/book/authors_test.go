package book

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAuthor(t *testing.T) {
	var authors = []struct {
		name         string
		introduction string
		profile      string
		country      string
	}{
		{"加西亚.马尔克斯", "1982年诺贝尔文学奖获得者", "", "哥伦比亚"},
		{"村上春树", "日本著名作家", "", "日本"},
		{"鲁迅", "中国现代文学的奠基人", "", "中国"},
		{"莫言", "2012年诺贝尔文学奖获得者", "", "中国"},
		{"余华", "中国现代文学的奠基人", "", "中国"},
		{"王小波", "中国现代文学的奠基人", "", "中国"},
		{"钱钟书", "中国现代文学的奠基人", "", "中国"},
		{"孙浩辉", "历史作家", "", "中国"},
		{"二月河", "历史作家", "", "中国"},
	}
	Convey("测试作者操作", t, func() {
		a := NewAuthor()
		Convey("测试作者添加", func() {
			for _, author := range authors {
				if a.checkName(author.name) != nil {
					continue
				}
				_, err := a.Add(author.name, author.introduction, author.profile, author.country)
				So(err, ShouldBeNil)
			}
		})

		Convey("测试作者获取", func() {
			all := a.Search("", 1, 10)
			So(len(all), ShouldBeGreaterThan, 0)
			author := a.checkName("加西亚.马尔克斯")
			So(author.Name, ShouldEqual, "加西亚.马尔克斯")
			author1 := a.GetById(author.ID)
			So(author1.Name, ShouldEqual, "加西亚.马尔克斯")
			_, err := a.UpdateAuthor(author.ID, "加西亚.马尔克斯1", "", "", "")
			So(err, ShouldBeNil)
			author2 := a.GetById(author.ID)
			So(author2.Name, ShouldEqual, "加西亚.马尔克斯1")
			_, err = a.UpdateAuthor(author.ID, "加西亚.马尔克斯", "", "", "")
			So(err, ShouldBeNil)
			author3 := a.GetById(author.ID)
			So(author3.Name, ShouldEqual, "加西亚.马尔克斯")
			ids := []int32{}
			for _, author := range all {
				ids = append(ids, author.ID)
			}
			authors := a.GetByIds(ids)
			So(len(authors), ShouldEqual, len(ids))
		})
	})
}
