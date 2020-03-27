package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/models"
	pageModel "github.com/bingjian-zhu/gin-vue-admin/page"
	"github.com/bingjian-zhu/gin-vue-admin/page/emun"
	"github.com/bingjian-zhu/gin-vue-admin/repository"
)

// UserService 注入IUserRepository
type UserService struct {
	Repository repository.IUserRepository `inject:""`
}

//CheckUser 身份验证
func (a *UserService) CheckUser(username string, password string) bool {
	where := models.User{Username: username, Password: password}
	return a.Repository.CheckUser(&where)
}

//GetUserAvatar 获取用户头像
func (a *UserService) GetUserAvatar(username string) string {
	return a.Repository.GetUserAvatar(username)
}

//GetRoles 获取用户角色
func (a *UserService) GetRoles(username string) []string {
	return a.Repository.GetRoles(username)
}

//GetUsers 获取用户信息
func (a *UserService) GetUsers(page, pagesize int, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	users := a.Repository.GetUsers(page, pagesize, &total, maps)
	var pageUsers []pageModel.Users
	var pageUser pageModel.Users
	for _, user := range *users {
		pageUser.ID = user.ID
		pageUser.Name = user.Username
		pageUser.Password = user.Password
		pageUser.Avatar = user.Avatar
		pageUser.UserType = emun.GetUserType(user.UserType)
		pageUser.State = emun.GetStatus(user.State)
		pageUser.Deteled = emun.GetDeleted(user.Deleted)
		pageUser.CreatedOn = user.CreatedOn.Format("2006-01-02 15:04:05")
		pageUsers = append(pageUsers, pageUser)
	}
	res["list"] = &pageUsers
	res["total"] = total
	return &res
}

//AddUser 新建用户
func (a *UserService) AddUser(user *models.User) bool {
	return a.Repository.AddUser(user)
}

//ExistUserByName 判断用户名是否已存在
func (a *UserService) ExistUserByName(username string) bool {
	return a.Repository.ExistUserByName(username)
}

//UpdateUser 更新用户
func (a *UserService) UpdateUser(user *models.User) bool {
	return a.Repository.UpdateUser(user)
}

//DeleteUser 删除用户
func (a *UserService) DeleteUser(id int) bool {
	return a.Repository.DeleteUser(id)
}
