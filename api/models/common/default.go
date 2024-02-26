package common

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var sqlDB *sql.DB
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("work:123456@tcp(192.168.3.103:3306)/bookstore"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err = db.DB()
	if err != nil {
		panic("failed to get sqlDB")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}

func GetDB() *gorm.DB {
	return db
}
