package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IUserRepository User接口定义
type IUserRepository interface {
	//CheckUser 身份验证
	CheckUser(username string, password string) bool
	//GetUserAvatar 获取用户头像
	GetUserAvatar(username string) string
	//GetRoles 获取用户角色
	GetRoles(username string) []string
	//GetUsers 获取用户信息
	GetUsers(PageNum int, PageSize int, total *uint64, maps interface{}) *[]models.User
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(username string) bool
	//UpdateUser 更新用户
	UpdateUser(user *models.User) bool
	//DeleteUser 更新用户
	DeleteUser(id int) bool
}
