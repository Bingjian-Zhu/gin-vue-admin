package models

import (
	"time"
)

//Model 常用属性定义
type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}
