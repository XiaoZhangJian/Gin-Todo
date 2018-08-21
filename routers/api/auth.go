package api

import (
	"Gin-Todo/models"
	"Gin-Todo/pkg/app"
	"Gin-Todo/pkg/e"
	"Gin-Todo/pkg/util"
	"fmt"
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func RegisterNewUser(c *gin.Context) {

	appG := app.Gin{c}

	name := c.Query("nick_name")
	email := c.Query("email")
	pwd := c.Query("password")

	valid := validation.Validation{}
	auth := models.AuthenticateRequest{
		NickName: name,
		Email:    email,
		Password: pwd,
	}
	ok, _ := valid.Valid(&auth)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if models.ExistUserByEmail(email) {
		appG.Response(http.StatusOK, e.ERROR_REG_EMAIL_EXIST_FAIL, nil)
		return
	}
	uid, _ := models.GetUUID(email)
	user, err := models.AddUser(uid, name, email, pwd)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_REG_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

func AuthenticateUser(c *gin.Context) {
	name := c.Query("nick_name")
	email := c.Query("email")
	pwd := c.Query("password")

	valid := validation.Validation{}
	auth := models.AuthenticateRequest{
		NickName: name,
		Email:    email,
		Password: pwd,
	}
	ok, _ := valid.Valid(&auth)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		user, err := models.GetAuthenticateUser(auth)
		if err != nil {
			fmt.Errorf("fail to Authenticate User err : %v", err)
		}

		token, err := util.GenerateToken(user)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			data["user"] = user
			code = e.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			log.Println("---Error---", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// func GetUser(c *gin.Context) {
// 	id := com.StrTo(c.Param("id")).MustInt()
// 	valid := validation.Validation{}
// 	valid.Required(id, "id").Message("ID不能为空")
// }

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
