package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/bingjian-zhu/gin-vue-admin/repository"
)

// RoleService IRoleRepository
type RoleService struct {
	Repository repository.IRoleRepository `inject:""`
}

//GetUserRoles 分页返回Articles获取用户身份信息
func (c *RoleService) GetUserRoles(userName string) (roles []models.Role) {
	return c.Repository.GetUserRoles(userName)
}
