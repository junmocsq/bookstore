package book

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	arr := []string{"热血", "散文", "青春", "健康", "美食", "文学", "科技", "历史", "小说", "教育",
		"心理", "管理", "经济", "法律", "政治", "军事", "哲学", "艺术", "设计", "摄影", "音乐", "舞蹈", "戏剧"}
	for _, v := range arr {
		NewTag().Add(v)
	}
}

func TestUpdate(t *testing.T) {
	Convey("测试更新标签", t, func() {
		NewTag().UpdateName(1, "热血2")
		r := NewTag().UpdateName(1, "热血1")
		So(r, ShouldEqual, 1)

		NewTag().UpdateStatus(1, 2)
		tt := NewTag().GetById(1)
		So(tt.Status, ShouldEqual, 2)

		NewTag().UpdateStatus(1, 1)
		tt = NewTag().GetById(1)
		So(tt.Status, ShouldEqual, 1)
	})

}
