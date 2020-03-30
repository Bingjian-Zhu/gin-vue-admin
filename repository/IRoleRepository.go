package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IRoleRepository Role接口定义
type IRoleRepository interface {
	//GetUserRoles 分页返回Articles获取用户身份信息
	GetUserRoles(where interface{}) *[]models.Role
	//GetRoles 获取用户角色
	GetRoles(sel *string, where interface{}) *[]string
	//AddRole 添加用户角色
	AddRole(role *models.Role) bool
	//GetRole 获取角色
	GetRole(where interface{}) *models.Role
}
