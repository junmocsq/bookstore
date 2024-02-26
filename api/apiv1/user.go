package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junmocsq/bookstore/api/models/user"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	BaseController
}

type User struct {
	ID         int32  `json:"id"`
	Nickname   string `json:"nickname"`
	Account    string `json:"account"`
	Profile    string `json:"profile"`
	Phone      string `json:"phone"`
	NationCode string `json:"nation_code"`
	Email      string `json:"email"`
	Gender     int8   `json:"gender"`
	GenderStr  string `json:"gender_str"`
	WhatsUp    string `json:"whats_up"`
	Bananas    int32  `json:"bananas"`
	Apples     int32  `json:"apples"`
	LastSignIn string `json:"last_sign_in"`
	CreatedAt  string `json:"created_at"`
}

func formatUser(u *user.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		ID:         u.ID,
		Nickname:   u.Nickname,
		Account:    u.Account,
		Profile:    u.Profile,
		Phone:      u.Phone,
		NationCode: u.NationCode,
		Email:      u.Email,
		Gender:     u.Gender,
		GenderStr:  tools.Mapping("gender", u.Gender),
		WhatsUp:    u.WhatsUp,
		Bananas:    u.Bananas,
		Apples:     u.Apples,
		LastSignIn: tools.Time2Read(u.LastSignIn),
		CreatedAt:  tools.Time2Read(u.CreatedAt),
	}
}

func NewUserController() *UserController {
	return &UserController{}
}

// @Summary	获取用户信息
// @tags		user
// @Produce	json
// @Success	200	{object}	User	"成功"
// @Failure	400	{object}	string	"请求错误"
// @Failure	500	{object}	string	"内部错误"
// @Router		/apiv1/user [get]
func (u *UserController) User(c *gin.Context) {
	uid, _ := c.Get("uid")
	c.JSON(http.StatusOK, gin.H{
		"UID": formatUser(user.NewUser().GetById(int32(uid.(int)))),
	})
}

type UserUpdateRequest struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
	Gender   int8   `json:"gender"`
	WhatsUp  string `json:"whats_up"`
}

// @Summary	修改数据
// @tags		user
// @Produce	json
// @Accept		json
// @Param		body	body		UserUpdateRequest	true	"修改数据"
// @Success	200		{object}	User				"成功"
// @Failure	400		{object}	string				"请求错误"
// @Failure	500		{object}	string				"内部错误"
// @Router		/apiv1/user/update [post]
func (u *UserController) Update(c *gin.Context) {
	var req UserUpdateRequest
	if c.ShouldBind(&req) != nil {
		logrus.WithField("apiv1_ctrl", "user_update").Error("参数错误")
	}
	c.JSON(http.StatusOK, req)
}

// 密码修改
func (u *UserController) UpdatePasswd(c *gin.Context) {

}

// 绑定电话
func (u *UserController) BindPhone(c *gin.Context) {

}

// 绑定邮箱
func (u *UserController) BindEmail(c *gin.Context) {

}
