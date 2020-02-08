package routers

import (
	"github.com/gin-gonic/gin"
	"lottery/errcode"
	"lottery/models"
	"lottery/wechat"
	"net/http"
)

type HandlerFunc func(c *gin.Context) error

func wrapper(handlerFunc HandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		err := handlerFunc(c)
		if err != nil {
			var e *errcode.ServiceError
			if h, ok := err.(*errcode.ServiceError); ok {
				e = h
			} else {
				e = errcode.ErrorServerErr
			}
			if e.Code != errcode.OK {
				c.JSON(e.Status, e)
				c.Abort()
			} else {
				c.JSON(e.Status, e)
			}
		}
	}
}

func InitRouters(engine *gin.Engine) {
	engine.HandleMethodNotAllowed = true
	engine.LoadHTMLGlob("./templates/*.html")
	engine.GET("/lottery/auth", wechat.WxOAuth2RedirectHandler)
	engine.GET("/lottery/login", wechat.WxLogin)
	engine.GET("/lottery/user", wrapper(models.GetUserInfo))
	engine.POST("/lottery/user/checkin", wrapper(models.UserCheckin))
	engine.GET("/lottery/user/list", wrapper(models.GetUserList))
	engine.GET("/lottery/luckyNum", wrapper(models.GetLuckyNumByOpenID))
	engine.GET("/lottery/next", wrapper(models.GetNextRewardHandler))
	engine.POST("/lottery/roll", wrapper(models.LotteryDraw))
	engine.NoMethod(NoMethodHandler)
	engine.NoRoute(NotFoundHandler)
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, errcode.ErrorNotFound)
}

func NoMethodHandler(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, errcode.ErrorNoMethod)
}
