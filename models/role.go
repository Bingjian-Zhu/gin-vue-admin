package models

//Role 身份信息结构体
type Role struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Value    string `json:"value"`
}
