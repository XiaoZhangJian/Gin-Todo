# Gin-Todo 清单列表

### 第三方包

- https://ini.unknwon.io/


<!-- 
http://chuquanl.com/golang-cron%E7%AE%80%E4%BB%8B%E5%8F%8A%E4%BD%BFcron%E6%94%AF%E6%8C%81%E5%B8%A6%E5%8F%82%E6%95%B0%E4%BB%BB%E5%8A%A1%E8%B0%83%E7%94%A8/

https://www.zhaotool.com/convert/cron.html

https://github.com/EDDYCJY/blog/blob/09a2e466d088b2bef6ae35700efb7df3ccd07c44/golang/gin/2018-02-16-Gin%E5%AE%9E%E8%B7%B5-%E8%BF%9E%E8%BD%BD%E4%BA%8C-%E6%90%AD%E5%BB%BABlogAPIs-01.md

https://github.com/tokillabeast/golang-todo/blob/eea49c09fd2400b45c790fd0feebe604cde89391/backend/models/todo.go -->


```
 // 注册 POST
 http://localhost:8000/register?email=zhangsan@qq.com&password=123456&nick_name=赵日天

 // 添加Todo POST
 http://localhost:8000/api/v1/todos?user_id=jQFjVt6PtAp&title=学习Golang&desc=入门3&

 // 获取用户全部Todo GET
 http://localhost:8000/api/v1/todos?user_id=jQFjVt6PtAp

 // 获取单个Todo 内容 GET
 http://localhost:8000/api/v1/todos/cb58ddf7-cb86-434d-8ffd-c87cde11b5ba

 // 指定id修改Todo PUT
 http://localhost:8000/api/v1/todos/cb58ddf7-cb86-434d-8ffd-c87cde11b5ba?title=修改1&desc=修改内容1&complete=true

 // 指定Id删除 Todo DELETE
 http://localhost:8000/api/v1/todos/cb58ddf7-cb86-434d-8ffd-c87cde11b5ba

```