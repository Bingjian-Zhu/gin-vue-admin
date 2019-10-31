# gin-vue
该项目是gin+vue的前后端分离项目，使用gorm访问MySQL

使用jwt，对API接口进行权限控制。[教程](https://bingjian-zhu.github.io/2019/09/03/gin-jwt%E5%AF%B9API%E8%BF%9B%E8%A1%8C%E6%9D%83%E9%99%90%E6%8E%A7%E5%88%B6/)

在token过期后的半个小时内，用户再次操作会自动刷新token

### go后台程序运行方式

1.在MySQL中新建blog数据库，运行文件夹/docs/sql中的mysql.sql脚本

2.在文件夹/conf中修改配置文件api.ini中的数据库连接配置

3.在gin-vue目录下运行`go run main.go`

### vue运行方式请看文件夹/vue-admin中的README.md

喜欢请star

联系方式（QQ）：939496716
