package v1

import (
	"Gin-Todo/pkg/app"
	"Gin-Todo/pkg/e"
	"Gin-Todo/pkg/setting"
	"Gin-Todo/pkg/util"
	"Gin-Todo/service/todo_service"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 根据ID获取Todo
// @Produce  json
// @Param id param string true "ID"
// @Success 200 {string} json "{"code": 200,"data": {"todo": {"id": "e999d5a3-8f84-4812-8e57-0d66450a9aee","user_id": "jQFjVt6PtAp","name": "学习Golang","desc": "入门1","completed": false} },"msg": "ok"}"
// @Router /api/v1/todos/{id} [get]
func GetTodo(c *gin.Context) {
	appG := app.Gin{c}
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	todoService := todo_service.Todo{
		Id: id,
	}
	if exist, err := todoService.ExistByID(); !exist || err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_TODO_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	todo, err := todoService.Find()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_TODO_FAIL, nil)
		return
	}

	data["todo"] = todo
	appG.Response(http.StatusOK, e.SUCCESS, data)

}

// 获取该用户的全部Todo
// @Produce  json
// @Param uid query string false "UserId"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id":"e999d5a3-8f84-4812-8e57-0d66450a9aee","user_id":"jQFjVt6PtAp","name":"学习Golang","desc":"入门1","completed":false},{"id":"ea85e87c-487e-4449-b9c2-8cd1e13c3a54","user_id":"jQFjVt6PtAp","name":"学习Golang","desc":"入门2","completed":false}],"total":2},"msg":"ok"}"
// @Router /api/v1/todos [get]
func GetTodos(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}
	userId := c.Query("uid")
	if userId != "" {
		valid.Required(userId, "user_id").Message("用户ID不能为空")
	}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	todoService := todo_service.Todo{
		UserID:   userId,
		PageNum:  util.GetPage(c),
		PageSize: setting.PageSize,
	}
	todos, err := todoService.FindAll()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_TODOS_FAIL, nil)
		return
	}

	total, err := todoService.Count() //models.GetTodoTotal(maps)

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_TODOS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = todos
	data["total"] = total
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// 新增任务
// @Produce json
// @Param uid query string false "UserId"
// @Param name query string false "Name"
// @Param desc query string false "Desc"
// @Success 200 {string} json "{"code":200,"data":null,"msg":"ok"}"
// @Router /api/v1/todos [post]
func AddTodo(c *gin.Context) {
	appG := app.Gin{c}
	userId := c.Query("uid")
	name := c.Query("name")
	desc := c.Query("desc")
	valid := validation.Validation{}
	valid.Required(userId, "user_id").Message("用户ID不能为空")
	valid.Required(name, "name").Message("任务名称不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	todoService := todo_service.Todo{
		UserID:    userId,
		Name:      name,
		Desc:      desc,
		Completed: false,
	}

	if err := todoService.Create(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_ADD_TODO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//修改任务
// @Produce json
// @Param id param string false "Id"
// @Param name query string false "Name"
// @Param desc query string false "Desc"
// @Param complete query string false "Completed"
// @Success 200 {string} json "{"code":200,"data":null,"msg":"ok"}"
// @Router /api/v1/todos/{id} [put]
func EditTodo(c *gin.Context) {
	appG := app.Gin{c}
	id := c.Param("id")
	name := c.Query("name")
	desc := c.Query("desc")
	b := c.Query("complete")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(name, "name").Message("任务名称不能为空")
	valid.Required(b, "b").Message("布尔值不能为空")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	complete, _ := strconv.ParseBool(b)
	todoService := todo_service.Todo{
		Id:        id,
		Name:      name,
		Desc:      desc,
		Completed: complete,
	}
	exist, err := todoService.ExistByID()
	if err != nil || !exist {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_TODO_FAIL, nil)
		return
	}
	if err := todoService.Update(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_EDIT_TODO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//删除任务
// @Produce json
// @Param id param string false "Id"
// @Success 200 {string} json "{"code":200,"data":null,"msg":"ok"}"
// @Router /api/v1/todos/{id} [delete]
func DelTodo(c *gin.Context) {
	appG := app.Gin{c}

	id := c.Param("id")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	// code := e.INVALID_PARAMS

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	todoService := todo_service.Todo{
		Id: id,
	}

	exist, err := todoService.ExistByID()
	if err != nil || !exist {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_TODO_FAIL, nil)
		return
	}

	if err := todoService.Delete(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_TODO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
