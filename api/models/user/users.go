package user

import (
	"errors"
	"fmt"

	"github.com/junmocsq/bookstore/api/models/common"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID         int32
	Nickname   string `gorm:"type:string;size:20;not null;default:'';comment:昵称" json:"nickname"`
	Account    string `gorm:"uniqueIndex;type:string;size:20;not null;default:'';comment:登录名称纯英文数字下划线" json:"account"`
	Passwd     string `gorm:"type:string;size:64;not null;default:'';comment:sha256" json:"passwd"`
	Salt       string `gorm:"type:string;size:6;not null;default:'';comment:随机6位加盐密码" json:"salt"`
	Profile    string `gorm:"type:string;size:100;not null;default:'';comment:头像地址" json:"profile"`
	Phone      string `gorm:"uniqueIndex:idx_phone;type:string;size:20;not null;default:'';comment:电话号码" json:"phone"`
	NationCode string `gorm:"uniqueIndex:idx_phone;type:string;size:5;not null;default:'86';comment:电话国家编号" json:"nation_code"`
	Email      string `gorm:"type:string;size:100;not null;default:'';comment:邮箱地址" json:"email"`
	Gender     int8   `gorm:"type:tinyint;not null;default:0;comment:性别 0 未知 1 male 2 female" json:"gender"`
	WhatsUp    string `gorm:"type:string;size:50;not null;default:'';comment:个性签名" json:"whats_up"`
	Bananas    int32  `gorm:"type:int;not null;default:0;comment:充值货币" json:"bananas"`
	Apples     int32  `gorm:"type:int;not null;default:0;comment:登录等虚拟货币" json:"apples"`
	LastSignIn int64  `gorm:"type:bigint;not null;default:0;comment:最后登录时间" json:"last_sign_in"`
	CreatedAt  int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (u *User) TableName() string {
	return "u_users"
}

func (u *User) Tag() string {
	return "u_users"
}
func (u *User) Key(id int32) string {
	return fmt.Sprintf("%d", id)
}

func NewUser() *User {
	return &User{}
}

// User表增删改查
func (u *User) add(nickname, account, passwd, salt, profile, phone, nationCode, email string, gender int8) (int32, error) {
	u.Nickname = nickname
	u.Account = account
	u.Passwd = passwd
	u.Salt = salt
	u.Profile = profile
	u.Phone = phone
	u.NationCode = nationCode
	u.Email = email
	u.Gender = gender
	var db = common.GetDB()
	stmt := db.DryRun().Create(u).Statement
	n, err := db.SetTag(u.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(u)
	if err != nil {
		logrus.WithField("model", "user_add").Error(err)
		return 0, errors.New("添加失败")
	}
	return int32(n), nil
}

// 电话注册
func (u *User) PhoneSignUp(phone, nationCode string) (int32, error) {
	salt := tools.CreateRandomString(6)
	account := tools.CreateRandomString(15)
	for {
		if u.GetByAccount(account) == nil {
			break
		}
		logrus.Warn(u.GetByAccount(account))
		account = tools.CreateRandomString(15)
	}
	passwd := tools.CreateRandomString(15)
	passwd = tools.CreatePasswd(salt, passwd)
	if u.GetByPhone(phone, nationCode) != nil {
		logrus.WithField("model", "user_phone_sign_up").Error("已注册")
		return 0, errors.New("已注册")
	}
	return u.add(account, account, passwd, salt, "", phone, nationCode, "", 0)
}

// 邮箱注册
func (u *User) EmailSignUp(email, passwd string) (int32, error) {
	if u.GetByEmail(email) != nil {
		logrus.WithField("model", "user_email_sign_up").Error("已注册")
		return 0, errors.New("已注册")
	}
	salt := tools.CreateRandomString(6)
	account := tools.CreateRandomString(15)
	for {
		if u.GetByAccount(account) == nil {
			break
		}
		account = tools.CreateRandomString(15)
	}
	passwd = tools.CreatePasswd(salt, passwd)
	return u.add(account, account, passwd, salt, "", "", "86", email, 0)
}

func (u *User) Update() int32 {

	return 0
}

func (u *User) GetById(id int32) *User {
	var db = common.GetDB()
	var user User
	stmt := db.DryRun().Where("id =?", id).First(&user).Statement
	err := db.SetTag(u.Tag()).SetKey(u.Key(id)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user)
	if err != nil {
		logrus.WithField("model", "user_GetById").Error(err)
		return nil
	}
	return &user
}

func (u *User) GetByAccount(account string) *User {
	var db = common.GetDB()
	var user User
	stmt := db.DryRun().Where("account =?", account).Find(&user).Statement
	err := db.SetTag(u.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user)
	if err != nil {
		logrus.WithField("model", "user_GetByAccount").Error(err)
		return nil
	}
	return &user
}

// 获取用户信息
func (u *User) GetByPhone(phone, nationCode string) *User {
	var db = common.GetDB()
	var user User
	stmt := db.DryRun().Where("phone =? and nation_code =?", phone, nationCode).First(&user).Statement
	err := db.SetTag(u.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user)
	if err != nil {
		logrus.WithField("model", "user_GetByPhone").Error(err)
		return nil
	}
	return &user
}

func (u *User) GetByEmail(email string) *User {
	var db = common.GetDB()
	var user User
	stmt := db.DryRun().Where("email =?", email).First(&user).Statement
	err := db.SetTag(u.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user)
	if err != nil {
		logrus.WithField("model", "user_GetByAccount").Error(err)
		return nil
	}
	return &user
}
