package api

import (
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Register Account godoc
// @Summary UserRegister 用户注册接口
// @Description Register account
// @Tags accounts
// @Accept  multipart/form-data
// @Produce  json
// @Param nickname query string true "昵称"
// @Param user_name query string true "登录名"
// @Param password query string true "密码"
// @Param password_confirm query string true "确认密码"
// @Success 200 {object} model.User
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// @Summary UserLogin 用户登录接口
// @Description account login
// @Tags accounts
// @Accept  multipart/form-data
// @Produce json
// @Param user_name query string true "登录名"
// @Param password query string true "密码"
// @Success 200 {object} model.User
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// @Summary UserMe 用户详情
// @Description get Users by ID
// @Accept  json
// @Produce  json
// @Param user path int true "Account ID"
// @Success 200 {object} model.User
// @Router /api/v1/user/me [get]
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// @Summary UserLogout 用户登出
// @Description logout Users by ID
// @Accept  json
// @Produce  json
// @Param user path int true "Account ID"
// @Success 200 {object} model.User
// @Router /api/v1/user/logout [get]
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
