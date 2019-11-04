package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//RoleRepository 注入IDb
type RoleRepository struct {
	Source datasource.IDb `inject:""`
}

//GetUserRoles 获取用户身份信息
func (c *RoleRepository) GetUserRoles(userName string) (roles []models.Role) {
	c.Source.DB().Where("user_name = ?", userName).Find(&roles)
	return
}
