package oauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"easy-api-template/serializer"
	oauth_service "easy-api-template/service/oauth"
	"easy-api-template/util"
)

// login
func Login(c *gin.Context) {
	svc := oauth_service.LoginParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err,
				Msg:  "参数错误，请检查！",
			},
			ResCode: serializer.PARAMSERROR,
		})
		return
	}
	result := svc.Login()
	if result.ResCode != serializer.LOGINSUCCESS {
		c.JSON(result.Code, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: result.Code,
				Data: result.Data,
				Msg:  result.Msg,
			},
			ResCode: result.ResCode,
		})
		return
	}
	err := util.SetLoginCookies(c,svc.UserName,result.Data.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err,
				Msg:  "set cookie failure",
			},
			ResCode: serializer.SETCOOKIEFAILURE,
		})
		return
	}
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
	//c.Redirect(http.StatusFound,conf.GetConfig("SreStabilityOperatingPlatform::FrontendURL").String())
}

// register
func Register(c *gin.Context) {
	svc := oauth_service.RegisterParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err,
				Msg:  "参数错误，请检查！",
			},
			ResCode: serializer.PARAMSERROR,
		})
		return
	}
	result := svc.Register()
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// logout
func Logout(c *gin.Context) {
	util.Logout(c)
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "logout success",
		},
		ResCode: serializer.LOGOUTSUCCESS,
	})
}

// get user info
func GetUserInfo(c *gin.Context)  {
	u,err := util.GetUserCookie(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err,
				Msg:  "get user cookie fail",
			},
			ResCode: serializer.LOGOUTSUCCESS,
		})
	}
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: u,
			Msg:  "get user info success",
		},
		ResCode: serializer.SUCCESS,
	})
}

func CheckLogin(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}