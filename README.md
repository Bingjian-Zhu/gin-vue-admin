# gin-vue-admin
* [演示地址](http://zbj-home.picp.io)：入门级服务器，请轻点~~

* 以动态路由的方式实现不同的角色加载不同的菜单

> 账号：admin  密码：111111
> 账号：test   密码：111111

* 该项目是gin+vue的前后端分离项目，使用gorm访问MySQL

* 项目结构进行分层，使用依赖注入的方式对项目进行解耦---[Gin实现依赖注入教程](https://www.cnblogs.com/FireworksEasyCool/p/11805148.html)

* 使用jwt，对API接口进行权限控制---[gin-jwt对API进行权限控制教程](https://www.cnblogs.com/FireworksEasyCool/p/11455834.html)

* 使用[go-playground/validator](https://github.com/go-playground/validator)开源库简化gin的请求校验---[gin请求数据校验教程](https://www.cnblogs.com/FireworksEasyCool/p/12794311.html)

* 用Docker上云---[多阶段构建Golang程序Docker镜像教程](https://www.cnblogs.com/FireworksEasyCool/p/12838875.html)

* 在token过期后的半个小时内，用户再次操作会自动刷新token

### 项目结构

<pre><code>
├── cmd  程序入口
├── common 通用模块代码
├── config 配置文件
├── controller API控制器
├── docs 数据库文件
├── models 数据表实体
├── page 页面数据返回实体
├── repository 数据访问层
├── router 路由
├── service 业务逻辑层
├── vue-admin Vue前端页面代码
</code></pre>

### 下载安装项目
`go get -x github.com/bingjian-zhu/gin-vue-admin/cmd`

### go后台程序运行方式

1.在MySQL中运行文件夹/docs中的mysql.sql脚本

2.在gin-vue-admin/cmd目录下运行`go run main.go`

### vue前端运行方式

请看文件夹/vue-admin中的README.md
