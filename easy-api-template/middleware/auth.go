package middleware

import (
	"github.com/gin-gonic/gin"
	"easy-api-template/serializer"
	"net/http"
	"easy-api-template/util"
)

// 登录认证拦截器
func MiddlewareAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := util.GetUserCookie(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code:  http.StatusUnauthorized,
				Data:  nil,
				Msg:   "认证失败",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
