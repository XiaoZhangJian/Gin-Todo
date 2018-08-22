package api

import (
	"Gin-Todo/models"
	"Gin-Todo/pkg/app"
	"Gin-Todo/pkg/e"
	"Gin-Todo/pkg/util"
	"Gin-Todo/service/auth_service"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type LoginResult struct {
	User  *models.User
	Token string
}

// @Summary 注册接口
// @Produce json
// @Param nick_name query string false "NickName"
// @Param email query string false "email"
// @Param password query string false "Password"
// @Success 200 {string} json "{"code": 200,"data": "注册成功","msg": "ok"}"
// @Router /reg [post]
func Register(c *gin.Context) {
	appG := app.Gin{c}

	name := c.Query("nick_name")
	email := c.Query("email")
	pwd := c.Query("password")

	fmt.Println(name, email, pwd)

	valid := validation.Validation{}
	valid.Required(name, "nick_name").Message("用户名不能为空")
	valid.Required(email, "email").Message("邮件格式不正确")
	valid.Required(pwd, "password").Message("密码输入不正确")

	if valid.HasErrors() {
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

// @Summary 登录接口
// @Produce json
// @Param email query string false "email"
// @Param password query string false "Password"
// @Success 200 {string} json "{"code":200,"data":{"User":{"ID":2,"created_at":1534758059,"modified_at":1534758059,"uid":"jQFjVt6PtAp","nick_name":"赵日天","email":"zhangsan@qq.com","Password":"123456"},"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwiY3JlYXRlZF9hdCI6MTUzNDc1ODA1OSwibW9kaWZpZWRfYXQiOjE1MzQ3NTgwNTksInVpZCI6ImpRRmpWdDZQdEFwIiwibmlja19uYW1lIjoi6LW15pel5aSpIiwiZW1haWwiOiJ6aGFuZ3NhbkBxcS5jb20iLCJQYXNzd29yZCI6IjEyMzQ1NiIsImV4cCI6MTUzNDkzMjI3MywiaXNzIjoiSmFzb24ifQ.e-6bH7SLVupRg-OzKcubYxlOW18CUW3R15A1F1r1RHM"},"msg":"ok"}"
// @Router /login [get]
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
