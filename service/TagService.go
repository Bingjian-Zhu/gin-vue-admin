package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/common/setting"
	"github.com/bingjian-zhu/gin-vue-admin/repository"
)

// TagService 注入ITagRepo
type TagService struct {
	Repository repository.ITagRepository `inject:""`
}

//ExistTagByID 根据ID判断标签是否存在
func (t *TagService) ExistTagByID(id int) bool {
	return t.Repository.ExistTagByID(id)
}

//GetTags 分页返回标签信息
func (t *TagService) GetTags(page int, maps interface{}) (data map[string]interface{}) {
	data["lists"] = t.Repository.GetTags(page, setting.Config.APP.Pagesize, maps)
	data["total"] = t.Repository.GetTagTotal(maps)
	return
}

//ExistTagByName 根据名称判断标签是否存在
func (t *TagService) ExistTagByName(name string) bool {
	return t.Repository.ExistTagByName(name)
}

//AddTag 添加标签
func (t *TagService) AddTag(name string, state int, createdBy string) bool {
	return t.Repository.AddTag(name, state, createdBy)
}

//EditTag 编辑标签
func (t *TagService) EditTag(id int, data interface{}) bool {
	return t.Repository.EditTag(id, data)
}

//DeleteTag 删除标签
func (t *TagService) DeleteTag(id int) bool {
	return t.Repository.DeleteTag(id)
}
