package models

//User 用户授权信息
type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// UserClaim 用户身份结构体
type UserClaim struct {
	UserName   string
	UserClaims []Claims
}
