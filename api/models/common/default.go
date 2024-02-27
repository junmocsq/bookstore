package common

import (
	"github.com/junmocsq/jlib/dbcache"
	"time"
)

func init() {
	dsn := "work:123456@tcp(192.168.3.103:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	dbcache.RegisterDb(dsn, "bookstore", true)

	sqlDB, err := GetDB().DB().DB()
	if err != nil {
		panic("failed to get sqlDB")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbcache.RedisCacheInit("192.168.3.103", "6379", "")

}

func GetDB(dbname ...string) *dbcache.Dao {
	return dbcache.NewDb(dbname...)
}
