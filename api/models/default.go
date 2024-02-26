package models

import (
	"github.com/junmocsq/bookstore/api/models/book"
	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/junmocsq/bookstore/api/models/user"
)

func init() {
	common.GetDB().AutoMigrate(&user.User{})

	common.GetDB().AutoMigrate(&book.Bookmark{}, &book.Book{}, &book.Category{}, &book.Chapter{}, &book.Content{}, &book.Section{}, &book.Tag{})
}
