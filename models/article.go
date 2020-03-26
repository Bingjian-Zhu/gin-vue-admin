package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 文章结构体
type Article struct {
	ID            int       `gorm:"primary_key" json:"id"`
	CreatedOn     time.Time `json:"created_on"`
	ModifiedOn    time.Time `json:"modified_on"`
	State         int       `json:"state"`
	TagID         int       `json:"tag_id"`
	Title         string    `json:"title"`
	Desc          string    `json:"desc"`
	Content       string    `json:"content"`
	CoverImageURL string    `json:"cover_image_url"`
	CreatedBy     string    `json:"created_by"`
	ModifiedBy    string    `json:"modified_by"`
	Tag           Tag       `json:"tag"`
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
