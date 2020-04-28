package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User 用户授权信息
type User struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Username   string    `json:"username" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Avatar     string    `json:"avatar"`
	UserType   int       `json:"user_type"`
	Deleted    int       `json:"deteled"`
	State      int       `json:"state"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
}

//BeforeCreate CreatedOn赋值
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now())
	return nil
}

//BeforeUpdate ModifiedOn赋值
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now())
	return nil
}

// UserRole 用户身份结构体
type UserRole struct {
	UserName  string
	UserRoles []*Role
}
