package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IUserRepository User接口定义
type IUserRepository interface {
	//CheckUser 身份验证
	CheckUser(where interface{}) bool
	//GetUserAvatar 获取用户头像
	GetUserAvatar(sel *string, where interface{}) *string
	//GetUserID 获取用户ID
	GetUserID(sel *string, where interface{}) int
	//GetUsers 获取用户信息
	GetUsers(PageNum int, PageSize int, total *uint64, where interface{}) *[]models.User
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(where interface{}) bool
	//UpdateUser 更新用户
	UpdateUser(user *models.User, role *models.Role) bool
	//DeleteUser 更新用户
	DeleteUser(id int) bool
	//GetUserByID 获取用户
	GetUserByID(id int) *models.User
}
