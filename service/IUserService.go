package service

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IUserService UserService接口定义
type IUserService interface {
	//CheckUser 身份验证
	CheckUser(username string, password string) bool
	//GetUserAvatar 获取用户头像
	GetUserAvatar(username string) *string
	//GetRoles 获取用户角色
	GetRoles(username string) *[]string
	//GetUsers 获取用户信息
	GetUsers(page, pagesize int, maps interface{}) interface{}
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(username string) bool
	//UpdateUser 更新用户
	UpdateUser(user *models.User) bool
	//DeleteUser 删除用户
	DeleteUser(id int) bool
}
