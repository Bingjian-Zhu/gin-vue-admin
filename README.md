# gin-vue
该项目是gin+vue的前后端分离项目，使用gorm访问MySQL

使用依赖注入的方式对项目进行解耦

使用jwt，对API接口进行权限控制。[教程](https://bingjian-zhu.github.io/2019/09/03/gin-jwt%E5%AF%B9API%E8%BF%9B%E8%A1%8C%E6%9D%83%E9%99%90%E6%8E%A7%E5%88%B6/)

在token过期后的半个小时内，用户再次操作会自动刷新token

### 下载项目
go get github.com/bingjian-zhu/gin-vue-admin/cmd

### go后台程序运行方式

1.在MySQL中运行文件夹/docs中的mysql.sql脚本

2.在gin-vue-admin/cmd目录下运行`go run main.go`

### vue前端运行方式

请看文件夹/vue-admin中的README.md