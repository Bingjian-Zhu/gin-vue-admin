package service

import (
	"fmt"

	"github.com/bingjian-zhu/gin-vue-admin/repository"
)

// UserService 注入IUserRepository
type UserService struct {
	Repository repository.IUserRepository `inject:""`
}

//CheckUser 身份验证
func (a *UserService) CheckUser(username string, password string) bool {
	fmt.Println("33333333333333")
	return a.Repository.CheckUser(username, password)
}

//GetUserAvatar 获取用户头像
func (a *UserService) GetUserAvatar(username string) string {
	return a.Repository.GetUserAvatar(username)
}

//GetRoles 获取用户角色
func (a *UserService) GetRoles(username string) []string {
	return a.Repository.GetRoles(username)
}
