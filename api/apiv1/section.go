package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/junmocsq/bookstore/api/services"
)

type SectionController struct {
	BaseController
}

func NewSectionController() *SectionController {
	return &SectionController{}
}

// @Summary	获取图书目录
// @tags		book
// @Produce	json
// @Param		bid	path		int		true	"图书id"
// @Success	200	{object}	[]services.SSection	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/book/catalog/{bid} [get]
func (sc *SectionController) Catalog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response(400, "id错误"))
		return
	}
	l := services.NewSection().GetSectionsByBid(int32(id), false)
	c.JSON(http.StatusOK, response(200, l))
}

// @Summary	获取章节内容
// @tags		book
// @Produce	json
// @Param		bid	path		int		true	"图书id"
// @Param		sid	path		int		true	"小节id"
// @Success	200	{object}	[]services.SSection	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/book/section/{bid}/{sid} [get]
func (sc *SectionController) Section(c *gin.Context) {
	bid, err := strconv.Atoi(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response(400, "id错误"))
		return
	}
	sid, err := strconv.Atoi(c.Param("sid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response(400, "id错误"))
		return
	}
	l := services.NewSection().GetSectionByBidAndSid(int32(bid), int32(sid))
	c.JSON(http.StatusOK, response(200, l))
}
