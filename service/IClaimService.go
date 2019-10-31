package service

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IClaimService ClaimService接口定义
type IClaimService interface {
	//GetUserClaims 分页返回Articles获取用户身份信息
	GetUserClaims(userName string) (claims []models.Claims)
}
