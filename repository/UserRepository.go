package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/common/logger"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//UserRepository 注入IDb
type UserRepository struct {
	Source datasource.IDb `inject:""`
	Log    logger.ILogger `inject:""`
	Base   BaseRepository `inject:"inline"`
}

//CheckUser 身份验证
func (a *UserRepository) CheckUser(username string, password string) bool {
	var user models.User
	where := models.User{Username: username, Password: password}
	err := a.Base.First(&where, &user)
	if err != nil {
		a.Log.Errorf("用户名或密码错误", err)
		return false
	}
	return true
}

//GetUserAvatar 获取用户头像
func (a *UserRepository) GetUserAvatar(username string) string {
	var user models.User
	where := models.User{Username: username}
	sel := "avatar"
	err := a.Base.First(&where, &user, sel)
	if err != nil {
		a.Log.Error(err)
		return ""
	}
	return user.Avatar
}

//GetRoles 获取用户角色
func (a *UserRepository) GetRoles(username string) []string {
	var roles []string
	var user models.User
	where := models.User{Username: username}
	sel := "id"
	err := a.Base.First(&where, &user, sel)
	if err != nil {
		a.Log.Error(err)
		return roles
	}

	var arrRole []models.Role
	a.Source.DB().Select("value").Where(models.Role{UserID: user.ID}).Find(&arrRole)
	for _, role := range arrRole {
		roles = append(roles, role.Value)
	}
	return roles
}

//GetUsers 获取用户信息
func (a *UserRepository) GetUsers(PageNum int, PageSize int, total *uint64, maps interface{}) *[]models.User {
	var users []models.User
	err := a.Base.GetPages(&models.User{}, &users, PageNum, PageSize, total, maps)
	if err != nil {
		a.Log.Errorf("获取用户信息失败", err)
	}
	return &users
}

//AddUser 新建用户
func (a *UserRepository) AddUser(user *models.User) bool {
	err := a.Base.Create(user)
	if err != nil {
		a.Log.Errorf("新建用户失败", err)
		return false
	}
	return true
}

//ExistUserByName 判断用户名是否已存在
func (a *UserRepository) ExistUserByName(username string) bool {
	var user models.User
	where := models.User{Username: username}
	sel := "id"
	err := a.Base.First(&where, &user, sel)
	if err != nil {
		a.Log.Error(err)
		return false
	}
	return true
}

//UpdateUser 更新用户
func (a *UserRepository) UpdateUser(modUser *models.User) bool {
	var user models.User
	err := a.Base.FirstByID(&user, modUser.ID)
	if err != nil {
		a.Log.Error(err)
	}
	user.Username = modUser.Username
	user.Password = modUser.Password
	user.ModifiedBy = modUser.ModifiedBy
	modErr := a.Base.Save(&user)
	if modErr != nil {
		a.Log.Errorf("更新用户失败", modErr)
		return false
	}
	return true
}

//DeleteUser 删除用户
func (a *UserRepository) DeleteUser(id int) bool {
	var user models.User
	err := a.Base.FirstByID(&user, id)
	if err != nil || user.Username == "admin" {
		a.Log.Errorf("删除用户失败:不能删除admin账号")
		return false
	}
	err = a.Base.DeleteByID(&user, id)
	if err != nil {
		a.Log.Errorf("删除用户失败", err)
		return false
	}
	return true
}
