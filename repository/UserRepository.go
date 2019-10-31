package repository

import (
	"fmt"

	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//UserRepository 注入IDb
type UserRepository struct {
	Source datasource.IDb `inject:""`
	Base   BaseRepository `inject:"inline"`
}

//CheckUser 身份验证
func (a *UserRepository) CheckUser(username string, password string) bool {
	var user models.User
	where := models.User{Username: username, Password: password}
	err := a.Base.First(&where, &user)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//GetUserAvatar 获取用户头像
func (a *UserRepository) GetUserAvatar(username string) string {
	var user models.User
	a.Source.DB().Select("avatar").Where(models.User{Username: username}).First(&user)
	return user.Avatar
}

//GetRoles 获取用户角色
func (a *UserRepository) GetRoles(username string) []string {
	var user models.User
	a.Source.DB().Select("id").Where(models.User{Username: username}).First(&user)
	var claims []models.Claims
	a.Source.DB().Select("value").Where(models.Claims{UserID: user.ID}).Find(&claims)
	var roles []string
	for _, claim := range claims {
		roles = append(roles, claim.Value)
	}
	return roles
}
