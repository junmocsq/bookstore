package book

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBook(t *testing.T) {
	var books = []struct {
		aid                   int32
		title, summary, cover string
		categoryId            int32
		tagIds                []int32
	}{
		{1, "百年孤独", "布恩迪亚家族七代人的命运", "", 1, []int32{1, 2, 3}},
		{1, "霍乱时期的爱情", "霍乱时期的爱情简介", "", 1, []int32{1, 2, 3}},
		{5, "活着", "或者的故事", "", 1, []int32{1, 2, 3}},
		{7, "围城", "城里的人想出去，城外的人想进来", "", 1, []int32{1, 2, 3}},
		{8, "大秦帝国", "历史", "", 1, []int32{1, 2, 3}},
		{9, "康雍乾帝王", "清朝历史", "", 1, []int32{1, 2, 3}},
	}

	b := NewBook()
	Convey("测试书籍操作", t, func() {
		So(b.checkAuthorTitle(1, "testmmcskkk"), ShouldBeNil)
		Convey("测试书籍添加", func() {
			for _, book := range books {
				if b.checkAuthorTitle(book.aid, book.title) != nil {
					continue
				}
				_, err := b.Add(book.aid, book.title, book.summary, book.cover, book.categoryId, book.tagIds)
				So(err, ShouldBeNil)
			}
		})

		Convey("测试书籍获取", func() {
			all := b.Search("", 1, 0, 0, []int32{}, 1, 10)
			So(len(all), ShouldBeGreaterThan, 0)
			book := b.checkAuthorTitle(1, "百年孤独")
			So(book.Title, ShouldEqual, "百年孤独")
			book1 := b.GetById(book.ID)
			So(book1.Title, ShouldEqual, "百年孤独")
			_, err := b.UpdateBook(book.ID, "百年孤独1", "", "", 0, []int32{})
			So(err, ShouldBeNil)
			book2 := b.GetById(book.ID)
			So(book2.Title, ShouldEqual, "百年孤独1")
			b.UpdateBook(book.ID, "百年孤独", "", "", 0, []int32{5, 6, 8})
			book3 := b.GetById(book.ID)
			So(book3.Title, ShouldEqual, "百年孤独")

			books := b.GetByAid(1, 1, 10)
			So(len(books), ShouldBeGreaterThan, 0)

			books1 := b.Search("百年", 0, 0, 0, []int32{}, 1, 10)
			So(len(books1), ShouldBeGreaterThan, 0)
		})

		Convey("测试书籍上下架", func() {
			book := b.GetById(1)
			_, err := b.UpdateStatus(book.ID, 1)
			So(err, ShouldBeNil)
			n, err := b.UpdateStatus(book.ID, 2)
			So(err, ShouldBeNil)
			So(n, ShouldEqual, 1)
			book1 := b.GetById(book.ID)
			So(book1.Status, ShouldEqual, 2)
			_, err = b.UpdateStatus(book.ID, 1)
			So(err, ShouldBeNil)
			book2 := b.GetById(book.ID)
			So(book2.Status, ShouldEqual, 1)
		})

		Convey("测试书籍数据修改", func() {
			book := b.GetById(1, true)
			_, err := b.updateNum(book.ID, 1, "likes")
			So(err, ShouldBeNil)
			n, err := b.IncrementLikes(book.ID, 10)
			So(err, ShouldBeNil)
			So(n, ShouldEqual, 1)
			book1 := b.GetById(book.ID, true)
			So(book1.Likes, ShouldEqual, 11)
			_, err = b.IncrementLikes(book.ID, -1)
			So(err, ShouldBeNil)
			book2 := b.GetById(book.ID, true)
			So(book2.Likes, ShouldEqual, 10)

			_, err = b.updateNum(book.ID, 1, "views")
			So(err, ShouldNotBeNil)
		})
	})

}

func Test(t *testing.T) {
	t.SkipNow()
	_, filename, _, _ := runtime.Caller(0)
	file := path.Dir(filename) + "/../../conf/novel-list.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	type tempBook struct {
		Books []struct {
			BookInfo struct {
				Title      string `json:"title"`
				Intro      string `json:"intro"`
				Author     string `json:"author"`
				Categories []struct {
					Title string `json:"title"`
				} `json:"categories"`
			} `json:"bookInfo" `
		} `json:"books"`
	}
	var res tempBook
	json.Unmarshal(content, &res)
	// t.Log(res)

	for _, v := range res.Books {
		title := v.BookInfo.Title
		intro := v.BookInfo.Intro
		author := v.BookInfo.Author
		var aid int32 = 0
		existAuthor := NewAuthor().checkName(author)

		if existAuthor == nil {
			aid, err = NewAuthor().Add(author, "", "", "")
			if err != nil {
				t.Fatal(err)
			}
		} else {
			aid = existAuthor.ID
		}
		t.Log(aid)
		temp := []rune(intro)
		if len(temp) > 100 {
			intro = string(temp[:100])
		}
		n, err := NewBook().Add(aid, title, intro, "", 1, []int32{})
		if err != nil {
			t.Error(err)
		}
		t.Log(aid, n)
	}

}
