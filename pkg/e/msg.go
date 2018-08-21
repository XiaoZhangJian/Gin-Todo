package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TASK:               "已存在该任务名称",
	ERROR_NOT_EXIST_TASK:           "该任务不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",

	ERROR_AUTH_TOKEN: "Token生成失败",
	ERROR_AUTH:       "Token错误",
	ERROR_LOGIN:      "登录失败",

	ERROR_EXIST_USER:           "已存在该用户",
	ERROR_NOT_EXIST_USER:       "该用户不存在",
	ERROR_REG_EMAIL_EXIST_FAIL: "该邮箱已注册",
	ERROR_REG_USER_FAIL:        "用户注册失败",

	ERROR_GET_TODOS_FAIL:        "获取任务列表失败",
	ERROR_ADD_TODO_FAIL:         "添加任务失败",
	ERROR_CHECK_EXIST_TODO_FAIL: "检查任务是否存在失败",
	ERROR_EDIT_TODO_FAIL:        "指定任务修改失败",
	ERROR_DELETE_TODO_FAIL:      "指定任务删除失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
