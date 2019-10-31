package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//ITagRepository Tag接口定义
type ITagRepository interface {
	//GetTags 分页返回标签信息
	GetTags(pageNum int, pageSize int, maps interface{}) (tags []models.Tag)
	//GetTagTotal 获取标签总数
	GetTagTotal(maps interface{}) (count int)
	//ExistTagByName 根据名称判断标签是否存在
	ExistTagByName(name string) bool
	//AddTag 添加标签
	AddTag(name string, state int, createdBy string) bool
	//ExistTagByID 根据ID判断标签是否存在
	ExistTagByID(id int) bool
	//DeleteTag 删除标签
	DeleteTag(id int) bool
	//EditTag 编辑标签
	EditTag(id int, data interface{}) bool
}
