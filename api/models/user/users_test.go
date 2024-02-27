package user

import (
	"testing"

	"github.com/junmocsq/jlib/dbcache"
)

func TestMain(m *testing.M) {
	dsn := "work:123456@tcp(192.168.3.103:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	dbcache.RegisterDb(dsn, "bookstore", true)
	// dbcache.RegisterCache(dbcache.NewLocalCache())
	dbcache.RedisCacheInit("192.168.3.103", "6379", "")
	dbcache.RegisterCache(dbcache.NewRedisCache())
	m.Run()
}
func TestPhoneSignUp(t *testing.T) {
	u := NewUser()
	u.PhoneSignUp("13800138000", "86")
}

func TestEmailSignUp(t *testing.T) {
	u := NewUser()
	u.EmailSignUp("abc@gmail.com", "123456")
}

func TestGet(t *testing.T) {
	u := NewUser()

	for i := 0; i < 50; i++ {
		u.GetById(1)
	}
}
