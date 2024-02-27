package apiv1

import "github.com/gin-gonic/gin"

type SectionController struct {
	BaseController
}

func NewSectionController() *SectionController {
	return &SectionController{}
}

func (sc *SectionController) SectionList(c *gin.Context) {

}

func (sc *SectionController) Section(c *gin.Context) {

}