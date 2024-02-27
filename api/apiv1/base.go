package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BaseController struct {
}

func Prepare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// common logic before route handling
		logrus.WithField("prepare", "123").Info("start")
		c.Next()
		logrus.WithField("prepare", "456").Info("end")
	}
}

func CheckUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		// common logic before route handling
		// 校验是否登录
		logrus.WithField("CheckUser", "123").Info("start")
		c.Set("uid", 4)
		c.Next()
		logrus.WithField("CheckUser", "456").Info("end")
	}
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
