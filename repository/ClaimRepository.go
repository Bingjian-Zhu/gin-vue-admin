package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//ClaimRepository 注入IDb
type ClaimRepository struct {
	Source datasource.IDb `inject:""`
}

//GetUserClaims 获取用户身份信息
func (c *ClaimRepository) GetUserClaims(userName string) (claims []models.Claims) {
	var user models.User
	c.Source.DB().Where("username = ?", userName).First(&user)
	c.Source.DB().Where("user_id = ?", user.ID).Find(&claims)
	return
}
