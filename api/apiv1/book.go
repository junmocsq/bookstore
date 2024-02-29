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

// @Summary	获取图书详情
// @tags		book
// @Produce	json
// @Param		id	path		int		true	"图书id"
// @Success	200	{object}	User	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/book/{id} [get]
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
