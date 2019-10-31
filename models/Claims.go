package models

//Claims 身份信息结构体
type Claims struct {
	ID     int    `gorm:"primary_key" json:"claim_id"`
	UserID int    `json:"user_id"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}
