package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/bingjian-zhu/gin-vue-admin/repository"
)

// ClaimService IClaimRepository
type ClaimService struct {
	Repository repository.IClaimRepository `inject:""`
}

//GetUserClaims 分页返回Articles获取用户身份信息
func (c *ClaimService) GetUserClaims(userName string) (claims []models.Claims) {
	return c.Repository.GetUserClaims(userName)
}
