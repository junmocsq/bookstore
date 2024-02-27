package models

import (
	"github.com/junmocsq/bookstore/api/models/book"
	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/junmocsq/bookstore/api/models/user"
)

func init() {
	common.GetDB().DB().AutoMigrate(&user.BookFavorite{}, &user.BookLike{},
		&user.BookRead{}, &user.Bookmark{}, &user.CommentLike{}, &user.Comment{},
		&user.ReadTime{}, &user.RechargeRecord{}, &user.SectionRead{},
		&user.SectionSubscription{}, &user.ShelfBook{}, &user.Shelf{}, &user.Transaction{}, &user.User{})

	common.GetDB().DB().AutoMigrate(&book.Book{}, &book.Category{}, &book.Chapter{}, &book.Content{}, &book.Section{}, &book.Tag{})
}
