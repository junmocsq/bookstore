package book

import "testing"

func TestAdd(t *testing.T) {
	arr := []string{"热血", "散文", "青春", "健康", "美食", "文学", "科技", "历史", "小说", "教育",
		"心理", "管理", "经济", "法律", "政治", "军事", "哲学", "艺术", "设计", "摄影", "音乐", "舞蹈", "戏剧"}
	for _, v := range arr {
		r, err := NewTag().Add(v)
		t.Log(r, err)
	}
}

func TestUpdate(t *testing.T) {
	r := NewTag().UpdateName(1, "热血1")
	t.Log(r)

	r = NewTag().UpdateName(1, "热血")
	t.Log(r)

	r = NewTag().UpdateStatus(1, 2)
	t.Log(r)

	r = NewTag().UpdateStatus(1, 1)
	t.Log(r)

	t.Log(NewTag().GetByName("热血"))
}
