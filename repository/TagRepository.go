package repository

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//TagRepository 注入IDb
type TagRepository struct {
	Source datasource.IDb `inject:""`
}

//GetTags 分页获取标签
func (t *TagRepository) GetTags(pageNum int, pageSize int, maps interface{}) (tags []models.Tag) {
	t.Source.DB().Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

//GetTagTotal 获取标签总数
func (t *TagRepository) GetTagTotal(maps interface{}) (count int) {
	t.Source.DB().Model(&models.Tag{}).Where(maps).Count(&count)

	return
}

//ExistTagByName 根据名称判断标签是否存在
func (t *TagRepository) ExistTagByName(name string) bool {
	var tag models.Tag
	t.Source.DB().Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

//AddTag 添加标签
func (t *TagRepository) AddTag(name string, state int, createdBy string) bool {
	t.Source.DB().Create(&models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

//ExistTagByID 根据ID判断标签是否存在
func (t *TagRepository) ExistTagByID(id int) bool {
	var tag models.Tag
	t.Source.DB().Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

//DeleteTag 删除标签
func (t *TagRepository) DeleteTag(id int) bool {
	t.Source.DB().Where("id = ?", id).Delete(&models.Tag{})

	return true
}

//EditTag 编辑标签
func (t *TagRepository) EditTag(id int, data interface{}) bool {
	t.Source.DB().Model(&models.Tag{}).Where("id = ?", id).Updates(data)

	return true
}
