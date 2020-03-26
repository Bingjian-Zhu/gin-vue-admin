package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Tag 标签结构体
type Tag struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Name       string    `json:"name"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
	State      int       `json:"state"`
}

//BeforeCreate CreatedOn赋值
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
