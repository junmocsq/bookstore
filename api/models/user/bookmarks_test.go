package user

import "testing"

func TestAdd(t *testing.T) {
	b := NewBookmark()

	b.Add(1, 1, 1, 1, "123")
	b.Add(1, 1, 1, 2, "1234")

	t.Log(b.GetByBidAndSid(1, 1, 0))
}
