# gin-vue-admin
* [演示地址](http://zbj-home.picp.io)：入门级服务器，轻点~~

* 该项目是gin+vue的前后端分离项目，使用gorm访问MySQL

* 项目结构进行分层，使用依赖注入的方式对项目进行解耦---[Gin实现依赖注入教程](https://bingjian-zhu.github.io/2019/11/06/Gin%E5%AE%9E%E7%8E%B0%E4%BE%9D%E8%B5%96%E6%B3%A8%E5%85%A5/)

* 使用jwt，对API接口进行权限控制---[gin-jwt对API进行权限控制教程](https://bingjian-zhu.github.io/2019/09/03/gin-jwt%E5%AF%B9API%E8%BF%9B%E8%A1%8C%E6%9D%83%E9%99%90%E6%8E%A7%E5%88%B6/)

* 使用[go-playground/validator](https://github.com/go-playground/validator)开源库简化gin的请求校验---[gin请求数据校验教程](https://bingjian-zhu.github.io/2020/04/28/gin%E8%AF%B7%E6%B1%82%E6%95%B0%E6%8D%AE%E6%A0%A1%E9%AA%8C/)

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
