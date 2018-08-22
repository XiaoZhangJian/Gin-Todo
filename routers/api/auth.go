package api

import (
	"Gin-Todo/models"
	"Gin-Todo/pkg/app"
	"Gin-Todo/pkg/e"
	"Gin-Todo/pkg/util"
	"Gin-Todo/service/auth_service"
	"net/http"

	"github.com/Unknwon/com"
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

func EditUser(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("nick_name")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(name, "name").Message("昵称不能为空")
	code := e.INVALID_PARAMS

	data := make(map[string]interface{})

	if !valid.HasErrors() {
		if models.ExistUserByID(id) {
			data["nick_name"] = name
			models.EditUser(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_USER
		}
	} else {
		for _, err := range valid.Errors {
			data["error"] = err.Message
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// func AuthGitHub(c *gin.Context) {
// 	auth.Config.CookieSecret = []byte("7H9xiimk2QdTdYI7rDddfJeV")
// 	auth.Config.LoginSuccessRedirect = "/mainpage"
// 	auth.Config.CookieSecure = false

// 	githubHandler := auth.Github(setting.GitHubClientKey, setting.GitHubSecretKey, "user")
// 	githubHandler.ServeHTTP(c.Writer, c.Request)

// }

// func GetInfo(c *gin.Context) {
// 	auth.Config.CookieSecret = []byte("7H9xiimk2QdTdYI7rDddfJeV")
// 	auth.Config.LoginSuccessRedirect = "/mainpage"
// 	auth.Config.CookieSecure = false

// 	user, err := auth.GetUserCookie(c.Request)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"msg":  "ok",
// 		"user": user,
// 	})
// 	// fmt.Println("用户信息：", user.Picture(), user.Id(), user.Name())

// }
