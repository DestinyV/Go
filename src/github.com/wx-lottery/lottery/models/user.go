package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lottery/errcode"
	"time"
)

const UserSessionKey = "CurrentUser"

type User struct {
	OpenID      string     `gorm:"primary_key;column:openid" json:"openid"`
	Name        string     `gorm:"default:NULL" json:"name"`
	Nickname    string     `gorm:"default:NULL" json:"nickname"`
	GroupName   string     `gorm:"default:NULL" json:"groupName"`
	Telephone   string     `gorm:"default:NULL" json:"telephone"`
	HeadImgUrl  string     `gorm:"default:NULL" json:"headImgUrl"`
	CheckinTime *time.Time `gorm:"default:NULL" json:"checkin_time,omitempty"`
	Master      int        `gorm:"default:0" json:"master"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfo(c *gin.Context) error {
	openid, err := GetUserFromSession(c)
	if err != nil {
		return err
	}
	user := &User{OpenID: openid}
	if err = DB.Model(user).First(user).Error; err != nil {
		return errcode.ErrorUnauthorized
	}
	return errcode.ErrorOK.WithData(user)
}

// 用户签到
func UserCheckin(c *gin.Context) error {
	var (
		luckyNum LuckyNumber
		err      error
	)
	openid, err := GetUserFromSession(c)
	if err != nil {
		return err
	}
	tx := DB.Begin()
	if luckyNum, err = GenerateLuckyNum(openid, tx); err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Model(&User{OpenID: openid}).Update("CheckinTime", time.Now()).Error; err != nil {
		tx.Rollback()
		return errcode.ErrorServerErr
	}
	tx.Commit()
	return errcode.ErrorOK.WithData(luckyNum)
}

// 获取用户列表
func GetUserList(c *gin.Context) error {
	var users []User
	if err := DB.Order("checkin_time desc").Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return errcode.ErrorServerErr
	}
	return errcode.ErrorOK.WithData(users)
}

func GetUserFromSession(c *gin.Context) (openid string, err error) {
	session := sessions.Default(c)
	t := session.Get(UserSessionKey)
	if t == nil {
		return openid, errcode.ErrorServerErr
	}
	if openid, ok := t.(string); ok {
		return openid, nil
	}
	return openid, errcode.ErrorServerErr
}
