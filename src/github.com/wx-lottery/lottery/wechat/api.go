package wechat

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"lottery/errcode"
	"lottery/models"
	"lottery/oa"
	"lottery/util"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: time.Second * 10}

func WxOAuth2RedirectHandler(c *gin.Context) {
	now := time.Now()
	if now.Before(time.Date(2020, 1, 17, 16, 0, 0, 0, time.Local)) {
		c.HTML(http.StatusFound, "wrong_time.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "https://open.weixin.qq.com/connect/oauth2/authorize?appid="+AppID+
		"&redirect_uri=http%3A%2F%2Fwww.vankeytech.com%2flottery%2flogin&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect")
}

func WxLogin(c *gin.Context) {
	code := c.Query("code")
	ak, err := GetAccessToken(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, errcode.ErrorInvalidParam.WithError(err))
		c.Abort()
		return
	}
	employee, notEmployee := GetOAInfo(ak.OpenID)
	if notEmployee {
		c.Redirect(http.StatusFound, "/2020/banding.html")
		c.Abort()
		return
	}
	exists := models.User{}
	if models.DB.Where(&models.User{OpenID: ak.OpenID}).First(&exists).RecordNotFound() {
		userInfo, err := GetUserInfo(ak)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errcode.ErrorServerErr.WithError(err))
			c.Abort()
			return
		}
		user := models.User{
			OpenID:     userInfo.OpenID,
			Name:       employee.Name,
			Nickname:   userInfo.Nickname,
			HeadImgUrl: userInfo.HeadImgUrl,
			GroupName:  employee.Group.GroupName,
			Telephone:  employee.Telephone,
		}
		models.DB.Create(&user)
		err = setUserSession(c, user.OpenID)
	} else {
		err = setUserSession(c, exists.OpenID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errcode.ErrorServerErr.WithError(err))
		c.Abort()
		return
	}
	c.Redirect(http.StatusFound, "/2020/signIn.html")
}

func GetAccessToken(code string) (ak AccessTokenBody, err error) {
	resp, err := httpClient.Get("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + AppID +
		"&secret=" + AppSecret + "&code=" + code + "&grant_type=authorization_code")
	if err != nil {
		return
	}
	if err = util.ConvertResponseBodyToStruct(resp.Body, &ak); err != nil {
		return
	}
	if ak.ErrCode != 0 {
		return ak, errors.New(ak.ErrMsg)
	}
	ak.ExpiresAt = time.Now().Add(time.Second * time.Duration(ak.ExpiresIn))
	return
}

func GetUserInfo(ak AccessTokenBody) (userInfo WxUserInfo, err error) {
	resp, err := httpClient.Get("https://api.weixin.qq.com/sns/userinfo?access_token=" + ak.AccessToken +
		"&openid=" + ak.OpenID + "&lang=zh_CN")
	if err != nil {
		return
	}
	err = util.ConvertResponseBodyToStruct(resp.Body, &userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		return userInfo, errors.New(userInfo.ErrMsg)
	}
	return
}

func setUserSession(c *gin.Context, openid string) error {
	session := sessions.Default(c)
	session.Set(models.UserSessionKey, openid)
	return session.Save()
}

func GetOAInfo(openid string) (employee oa.Employee, notEmployee bool) {
	notEmployee = oa.DB.Where(&oa.Employee{OpenID: openid}).First(&employee).RecordNotFound()
	employee.Group = &oa.Group{}
	oa.DB.Model(&employee).Related(&employee.Group).Find(&employee)
	return
}

type ErrorMsg struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type AccessTokenBody struct {
	ErrorMsg
	AccessToken  string    `json:"access_token"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	OpenID       string    `json:"openid"`
	ExpiresAt    time.Time `json:"-"`
}

type WxUserInfo struct {
	ErrorMsg
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"unionid"`
}
