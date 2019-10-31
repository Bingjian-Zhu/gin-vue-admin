package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IClaimRepository Claim接口定义
type IClaimRepository interface {
	//GetUserClaims 分页返回Articles获取用户身份信息
	GetUserClaims(userName string) (claims []models.Claims)
}
