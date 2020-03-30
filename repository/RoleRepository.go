package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/logger"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//RoleRepository 注入IDb
type RoleRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

//GetUserRoles 获取用户身份信息
func (a *RoleRepository) GetUserRoles(where interface{}) *[]models.Role {
	var roles []models.Role
	if err := a.Base.Find(where, &roles, ""); err != nil {
		a.Log.Errorf("获取用户身份信息错误", err)
	}
	return &roles
}

//GetRoles 获取用户角色
func (a *RoleRepository) GetRoles(sel *string, where interface{}) *[]string {
	var arrRole []string
	var roles []models.Role
	if err := a.Base.Find(where, &roles, *sel); err != nil {
		a.Log.Errorf("获取用户角色失败", err)
	}
	for _, role := range roles {
		arrRole = append(arrRole, role.Value)
	}
	return &arrRole
}

//AddRole 添加用户角色
func (a *RoleRepository) AddRole(role *models.Role) bool {
	if err := a.Base.Create(&role); err != nil {
		a.Log.Errorf("添加用户角色失败", err)
		return false
	}
	return true
}

//GetRole 获取角色
func (a *RoleRepository) GetRole(where interface{}) *models.Role {
	var role models.Role
	if err := a.Base.First(where, &role); err != nil {
		a.Log.Errorf("获取角色失败", err)
	}
	return &role
}
