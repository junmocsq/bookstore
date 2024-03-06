package book

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestChapter(t *testing.T) {
	Convey("测试大章操作", t, func() {
		c := NewChapter()
		Convey("测试大章添加", func() {
			cs := c.GetChaptersByBid(1)
			for _, v := range cs {
				err := c.DeleteById(v.ID, v.Bid)
				So(err, ShouldBeNil)
			}
			_, err := c.Add(1, "第一部分 山水", "一章简介")
			_, err = c.Add(1, "第二部分 江河", "二章简介")
			_, err = c.Add(1, "第三部分 元婴", "三章简介")
			_, err = c.Add(1, "第四部分 三花聚顶", "四章简介")
			So(err, ShouldBeNil)
		})

		Convey("测试大章获取", func() {
			chapter := c.GetChaptersByBid(1)
			So(len(chapter), ShouldBeGreaterThan, 0)
			chapter1 := chapter[0]
			So(chapter1.Title, ShouldEqual, "第一部分 山水")
			n := c.Update(chapter1.ID, 1, "第一部分 山水1", "一章简介1")
			So(n, ShouldEqual, 1)
			chapter2 := c.GetChaptersByBid(1)
			chapter1 = chapter2[0]
			So(chapter1.Title, ShouldEqual, "第一部分 山水1")
			So(chapter1.Summary, ShouldEqual, "一章简介1")

			n = c.Update(chapter1.ID, 1, "第一部分 山水", "一章简介")
			So(n, ShouldEqual, 1)
			chapter3 := c.GetChaptersByBid(1)
			chapter1 = chapter3[0]
			So(chapter1.Title, ShouldEqual, "第一部分 山水")
			So(chapter1.Summary, ShouldEqual, "一章简介")

			n = c.Update(chapter1.ID, 1, "", "")
			So(n, ShouldEqual, 0)
		})
	})
}
