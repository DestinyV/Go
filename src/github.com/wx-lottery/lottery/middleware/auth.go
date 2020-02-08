package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"lottery/errcode"
	"lottery/models"
	"net/http"
	"net/url"
	"path"
)

func AuthHandler(excludes ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		uri, err := url.ParseRequestURI(c.Request.RequestURI)
		if err != nil {
			c.JSON(http.StatusBadRequest, errcode.ErrorInvalidParam.WithError(err))
			c.Abort()
			return
		}
		matched := false
		for _, pattern := range excludes {
			if ok, _ := path.Match(pattern, uri.Path); ok {
				matched = true
				break
			}
		}
		if !matched {
			session := sessions.Default(c)
			u := session.Get(models.UserSessionKey)
			if u == nil {
				c.JSON(http.StatusUnauthorized, errcode.ErrorUnauthorized)
				c.Abort()
			}
		}
	}
}
