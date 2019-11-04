package models

//User 用户授权信息
type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Deleted  int    `json:"deteled"`
	State    int    `json:"state"`
}

// UserRole 用户身份结构体
type UserRole struct {
	UserName  string
	UserRoles []Role
}
