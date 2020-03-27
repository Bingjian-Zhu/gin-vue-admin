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
