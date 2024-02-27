package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BaseController
}

func NewBookController() *BookController {
	return &BookController{}
}

func (b *BookController) Book(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func (b *BookController) BookList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
