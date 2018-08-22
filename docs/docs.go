// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-08-22 17:27:29.5541743 +0800 CST m=+0.030391849

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/api/v1/todos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取该用户的全部Todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserId",
                        "name": "uid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"lists\":[{\"id\":\"e999d5a3-8f84-4812-8e57-0d66450a9aee\",\"user_id\":\"jQFjVt6PtAp\",\"name\":\"学习Golang\",\"desc\":\"入门1\",\"completed\":false},{\"id\":\"ea85e87c-487e-4449-b9c2-8cd1e13c3a54\",\"user_id\":\"jQFjVt6PtAp\",\"name\":\"学习Golang\",\"desc\":\"入门2\",\"completed\":false}],\"total\":2},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserId",
                        "name": "uid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Desc",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":null,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/todos/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据ID获取Todo",
                "parameters": [
                    {},
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200,\"data\": {\"todo\": {\"id\": \"e999d5a3-8f84-4812-8e57-0d66450a9aee\",\"user_id\": \"jQFjVt6PtAp\",\"name\": \"学习Golang\",\"desc\": \"入门1\",\"completed\": false} },\"msg\": \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改任务",
                "parameters": [
                    {},
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Desc",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Completed",
                        "name": "complete",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":null,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除任务",
                "parameters": [
                    {},
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":null,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"User\":{\"ID\":2,\"created_at\":1534758059,\"modified_at\":1534758059,\"uid\":\"jQFjVt6PtAp\",\"nick_name\":\"赵日天\",\"email\":\"zhangsan@qq.com\",\"Password\":\"123456\"},\"Token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwiY3JlYXRlZF9hdCI6MTUzNDc1ODA1OSwibW9kaWZpZWRfYXQiOjE1MzQ3NTgwNTksInVpZCI6ImpRRmpWdDZQdEFwIiwibmlja19uYW1lIjoi6LW15pel5aSpIiwiZW1haWwiOiJ6aGFuZ3NhbkBxcS5jb20iLCJQYXNzd29yZCI6IjEyMzQ1NiIsImV4cCI6MTUzNDkzMjI3MywiaXNzIjoiSmFzb24ifQ.e-6bH7SLVupRg-OzKcubYxlOW18CUW3R15A1F1r1RHM\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reg": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NickName",
                        "name": "nick_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200,\"data\": \"注册成功\",\"msg\": \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
