package apiv1

import "github.com/gin-gonic/gin"

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

// 邮箱或账号登录
func (lo *LoginController) LoginWithEmailOrAccount(c *gin.Context) {

}

// 手机登录
func (lo *LoginController) LoginWithPhone(c *gin.Context) {

}

// 手机注册
func (lo *LoginController) SignUpWithPhone(c *gin.Context) {

}

// 邮箱注册
func (lo *LoginController) SignUpWithEmail(c *gin.Context) {

}
