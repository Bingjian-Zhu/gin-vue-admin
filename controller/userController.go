package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/bingjian-zhu/gin-vue-admin/common/codes"
	"github.com/bingjian-zhu/gin-vue-admin/page/models"
	"github.com/bingjian-zhu/gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

//User 注入IUserService
type User struct {
	Service service.IUserService `inject:""`
}

//GetUserInfo 根据token获取用户信息
func (a *User) GetUserInfo(c *gin.Context) {
	roles := jwt.ExtractClaims(c)
	userName := roles["userName"].(string)
	avatar := a.Service.GetUserAvatar(userName)
	code := codes.SUCCESS
	userRoles := a.Service.GetRoles(userName)
	data := models.User{Roles: userRoles, Introduction: "", Avatar: avatar, Name: userName}
	RespData(c, http.StatusOK, code, &data)
}

//Logout 退出登录
func (a *User) Logout(c *gin.Context) {
	RespOk(c, http.StatusOK, codes.SUCCESS)
}
