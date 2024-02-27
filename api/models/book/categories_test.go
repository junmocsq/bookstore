package book

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCategoriesAdd(t *testing.T) {
	m := map[string][]string{
		"文学":   {"小说", "散文", "诗歌", "戏剧", "杂文", "童话", "寓言", "其他"},
		"IT科技": {"计算机", "互联网", "科普", "科学", "科技", "其他"},
		"教辅":   {"中小学教辅", "考试", "外语", "工具书", "其他"},
		"艺术":   {"美术", "摄影", "音乐", "舞蹈", "戏剧", "其他"},
		"人文社科": {"历史", "哲学", "宗教", "政治", "军事", "社会", "法律", "文化", "其他"},
		"经管":   {"经济", "管理", "商业", "金融", "投资", "营销", "其他"},
		"娱乐":   {"影视", "音乐", "游戏", "动漫", "体育", "其他"},
	}
	c := NewCategory()
	for k, v := range m {
		id, err := c.Add(k, 0, 0)
		if err != nil {
			continue
		}
		for _, _v := range v {
			r, err := c.Add(_v, id, 0)
			t.Log(r, err)
		}
	}
}

func TestFetch(t *testing.T) {
	c := NewCategory()
	Convey("测试获取分类", t, func() {
		all := c.GetAll()
		So(len(all), ShouldBeGreaterThan, 0)

		c1 := c.GetById(1)
		c.Update(c1.ID, c1.Name, c1.Pid, c1.Idx+1, 1)
		c2 := c.GetById(1)
		So(c2.Idx, ShouldEqual, c1.Idx+1)
	})

}
