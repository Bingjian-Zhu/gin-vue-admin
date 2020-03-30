package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/logger"
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/jinzhu/gorm"
)

//UserRepository 注入IDb
type UserRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

//CheckUser 身份验证
func (a *UserRepository) CheckUser(where interface{}) bool {
	var user models.User
	if err := a.Base.First(where, &user); err != nil {
		a.Log.Errorf("用户名或密码错误", err)
		return false
	}
	return true
}

//GetUserAvatar 获取用户头像
func (a *UserRepository) GetUserAvatar(sel *string, where interface{}) *string {
	var user models.User
	if err := a.Base.First(where, &user, *sel); err != nil {
		a.Log.Errorf("获取用户头像失败", err)
	}
	return &user.Avatar
}

//GetUserID 获取用户ID
func (a *UserRepository) GetUserID(sel *string, where interface{}) int {
	var user models.User
	if err := a.Base.First(where, &user, *sel); err != nil {
		a.Log.Errorf("获取用户ID失败", err)
		return -1
	}
	return user.ID
}

//GetUsers 获取用户信息
func (a *UserRepository) GetUsers(PageNum int, PageSize int, total *uint64, where interface{}) *[]models.User {
	var users []models.User
	if err := a.Base.GetPages(&models.User{}, &users, PageNum, PageSize, total, where); err != nil {
		a.Log.Errorf("获取用户信息失败", err)
	}
	return &users
}

//AddUser 新建用户
func (a *UserRepository) AddUser(user *models.User) bool {
	if err := a.Base.Create(user); err != nil {
		a.Log.Errorf("新建用户失败", err)
		return false
	}
	return true
}

//ExistUserByName 判断用户名是否已存在
func (a *UserRepository) ExistUserByName(where interface{}) bool {
	var user models.User
	sel := "id"
	err := a.Base.First(&where, &user, sel)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		a.Log.Error(err)
		return false
	}
	return true
}

//UpdateUser 更新用户
func (a *UserRepository) UpdateUser(user *models.User, role *models.Role) bool {
	//使用事务同时更新用户数据和角色数据
	tx := a.Base.GetTransaction()
	if err := tx.Save(user).Error; err != nil {
		a.Log.Errorf("更新用户失败", err)
		tx.Rollback()
		return false
	}
	if err := tx.Save(&role).Error; err != nil {
		a.Log.Errorf("更新用户角色失败", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//DeleteUser 删除用户同时删除用户的角色
func (a *UserRepository) DeleteUser(id int) bool {
	//采用事务同时删除用户和相应的用户角色
	var (
		userWhere = models.User{ID: id}
		user      models.User
		roleWhere = models.Role{UserID: id}
		role      models.Role
	)
	tx := a.Base.GetTransaction()
	tx.Where(&roleWhere).Delete(&role)
	if err := tx.Where(&userWhere).Delete(&user).Error; err != nil {
		a.Log.Errorf("删除用户失败", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//GetUserByID 获取用户
func (a *UserRepository) GetUserByID(id int) *models.User {
	var user models.User
	if err := a.Base.FirstByID(&user, id); err != nil {
		a.Log.Error(err)
	}
	return &user
}
