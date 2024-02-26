package tools

import (
	"testing"
	"time"
)

func TestCreateRandomString(t *testing.T) {
	s := CreateRandomString(10)
	t.Log(s)
}

func BenchmarkCreteRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateRandomString(10)
	}
}

func TestSha256(t *testing.T) {
	s := CreateRandomString(100)
	t.Log(Sha256(s))
}

func TestPublishTime2Read(t *testing.T) {
	temp := time.Unix(time.Now().Unix()-86400*10, 0)
	t.Log(Time2Read(temp.Unix(), DATE))
	t.Log(PublishTime2Read(temp.Unix()))
}
