package api

import (
	"Gin-Todo/models"
	"Gin-Todo/pkg/app"
	"Gin-Todo/pkg/e"
	"Gin-Todo/pkg/util"
	"Gin-Todo/service/auth_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type LoginResult struct {
	User  *models.User
	Token string
}

func Register(c *gin.Context) {

	appG := app.Gin{c}
	name := c.Query("nick_name")
	email := c.Query("email")
	pwd := c.Query("password")
	valid := validation.Validation{}
	valid.Required(name, "nick_name").Message("用户名不能为空")
	valid.Email(email, "email").Message("邮件格式不正确")
	valid.Required(pwd, "password").Message("密码输入不正确")

	if !valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	authService := auth_service.User{
		NickName: name,
		Email:    email,
		Pwd:      pwd,
	}
	if authService.ExistByEmail() {
		appG.Response(http.StatusOK, e.ERROR_REG_EMAIL_EXIST_FAIL, nil)
		return
	}

	err := authService.Create()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_REG_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "注册成功")
}

func Login(c *gin.Context) {
	appG := app.Gin{c}

	email := c.Query("email")
	pwd := c.Query("password")

	valid := validation.Validation{}
	valid.Email(email, "email").Message("email 格式不正确")
	valid.Required(pwd, "password").Message("密码不能为空")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	authService := auth_service.User{
		Email: email,
		Pwd:   pwd,
	}
	user, err := authService.Find()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_LOGIN, nil)
		return
	}

	token, err := util.GenerateToken(user)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	loginResult := LoginResult{
		User:  user,
		Token: token,
	}

	appG.Response(http.StatusOK, e.SUCCESS, loginResult)
}
