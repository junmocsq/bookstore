package common

import (
	"path"
	"runtime"
	"time"

	"github.com/go-ini/ini"
	"github.com/junmocsq/jlib/dbcache"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	cfgs, err := ini.Load(path.Dir(filename) + "/../../conf/conf.ini")
	if err != nil {
		panic(err)
	}
	dsn := cfgs.Section("mysql").Key("dsn").Value()
	dbname := cfgs.Section("mysql").Key("dbname").Value()

	// dsn := "work:123456@tcp(192.168.3.103:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	dbcache.RegisterDb(dsn, dbname, true)
	// dbcache.Debug(true)

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

	redisHost := cfgs.Section("redis").Key("host").Value()
	redisPort := cfgs.Section("redis").Key("port").Value()

	dbcache.RedisCacheInit(redisHost, redisPort, "")
	dbcache.RegisterCache(dbcache.NewRedisCache())

}

func GetDB(dbname ...string) *dbcache.Dao {
	return dbcache.NewDb(dbname...)
}
