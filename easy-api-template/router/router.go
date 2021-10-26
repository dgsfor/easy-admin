package router

import (
	"github.com/gin-gonic/gin"
	oauth_controller "easy-api-template/controller/oauth"
	"easy-api-template/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	atLeastLoginMiddleware := middleware.MiddlewareAuthorization()
	r := gin.Default()
	r.Use(middleware.Cors())
	// 路由
	v1 := r.Group("/api/easy-api-template/v1")
	{
		// user
		ur := v1.Group("/oauth")
		ur.POST("/login", oauth_controller.Login)
		ur.POST("/register",oauth_controller.Register)
		ur.GET("/logout", oauth_controller.Logout)
		ur.GET("/userinfo",atLeastLoginMiddleware,oauth_controller.GetUserInfo)
		ur.GET("/check_login", atLeastLoginMiddleware, oauth_controller.CheckLogin)
	}
	return r
}
