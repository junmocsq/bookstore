package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/junmocsq/bookstore/api/services"
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
// @Success	200	{object}	services.SBook	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/book/{id} [get]
func (b *BookController) Book(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response(400, "id错误"))
		return
	}
	res, err := services.NewBook().Get(int32(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, response(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response(200, res))
}

func (b *BookController) BookList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// @Summary	获取小说分类
// @tags		book
// @Produce	json
// @Success	200	{object}	[]services.SCategory	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/book/categories [get]
func (b *BookController) Categories(c *gin.Context) {
	c.JSON(http.StatusOK, response(200, services.NewCategory().FormatCategory(1)))
}

// @Summary	获取小说分类
// @tags		book
// @Produce	json
// @Success	200	{object}	[]services.STag	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/book/tags [get]
func (b *BookController) Tags(c *gin.Context) {
	c.JSON(http.StatusOK, response(200, services.NewTag().GetTags(1)))
}
