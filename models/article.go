package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 文章结构体
type Article struct {
	ID            int       `gorm:"primary_key" json:"id"`
	State         int       `json:"state" validate:"min=0,max=1"`
	TagID         int       `json:"tag_id" validate:"gt=0"`
	Title         string    `json:"title" validate:"required"`
	Desc          string    `json:"desc" validate:"required"`
	Content       string    `json:"content" validate:"required"`
	CoverImageURL string    `json:"cover_image_url"`
	CreatedBy     string    `json:"created_by" validate:"required"`
	ModifiedBy    string    `json:"modified_by"`
	Tag           Tag       `json:"tag"`
	CreatedOn     time.Time `json:"created_on"`
	ModifiedOn    time.Time `json:"modified_on"`
}

//BeforeCreate 在创建Article之前，先把创建时间赋值
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

//BeforeUpdate 在更新Article之前，先把更新时间赋值
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
