package v1

import (
	"Gin-Todo/models"
	"Gin-Todo/pkg/e"
	"Gin-Todo/pkg/setting"
	"Gin-Todo/pkg/util"
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//获取多个任务
func GetTasks(c *gin.Context) {
	createdBy := c.Query("created_by")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if createdBy != "" {
		maps["created_by"] = createdBy
	}

	code := e.SUCCESS
	data["lists"] = models.GetTasks(util.GetPage(c), setting.PageSize, createdBy)
	data["total"] = models.GetTaskTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//新增任务
func AddTask(c *gin.Context) {

	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "1")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为255字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTaskByName(name) {
			code = e.SUCCESS
			models.AddTask(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TASK
		}
	} else {
		for _, err := range valid.Errors {
			log.Fatal(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//修改任务
func EditTask(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTaskByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTask(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TASK
		}
	} else {
		for _, err := range valid.Errors {
			log.Fatal(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除任务
func DeleteTask(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTaskByID(id) {
			models.DeleteTask(id)
		} else {
			code = e.ERROR_NOT_EXIST_TASK
		}
	} else {
		for _, err := range valid.Errors {
			log.Fatal(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
