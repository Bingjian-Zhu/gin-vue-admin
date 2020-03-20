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
	var role models.Role
	//此处不能使用事务同时创建用户和角色，因为Role表中需要UserID，而UserID需要插入用户数据后才生成，所以不能用事务，否则会报错
	if err := a.Base.Create(user); err != nil {
		a.Log.Errorf("新建用户失败", err)
		return false
	}
	//当成功插入User数据后，user为指针地址，可以获取到ID的值。省去了查数据库拿ID的值步骤
	role.UserID = user.ID
	role.UserName = user.Username
	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	if err := a.Base.Create(&role); err != nil {
		a.Log.Errorf("新建用户是创建角色失败", role)
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
		return false
	}
	user.Username = modUser.Username
	user.Password = modUser.Password
	user.ModifiedBy = modUser.ModifiedBy
	user.UserType = modUser.UserType

	var role models.Role
	where := models.Role{UserID: user.ID}
	err = a.Base.First(&where, &role)
	role.UserName = user.Username
	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	//使用事务同时更新用户数据和角色数据
	tx := a.Base.GetTransaction()
	if err := tx.Save(user).Error; err != nil {
		a.Log.Errorf("更新用户失败", err)
		tx.Rollback()
		return false
	}
	if err := tx.Save(role).Error; err != nil {
		a.Log.Errorf("更新用户角色失败", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//DeleteUser 删除用户同时删除用户的角色
func (a *UserRepository) DeleteUser(id int) bool {
	var user models.User
	err := a.Base.FirstByID(&user, id)
	if err != nil || user.Username == "admin" {
		a.Log.Errorf("删除用户失败:不能删除admin账号")
		return false
	}
	var role models.Role
	where := models.Role{UserID: id}
	//采用事务同时删除用户和相应的用户角色
	tx := a.Base.GetTransaction()
	tx.Where(&where).Delete(&role)
	err = tx.Where("id=?", id).Delete(&user).Error
	if err != nil {
		a.Log.Errorf("删除用户失败", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
