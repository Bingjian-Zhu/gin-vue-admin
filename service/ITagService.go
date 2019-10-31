package service

//ITagService TagService接口定义
type ITagService interface {
	//ExistTagByID 根据ID判断标签是否存在
	ExistTagByID(id int) bool
	//GetTags 分页返回标签信息
	GetTags(page int, maps interface{}) (data map[string]interface{})
	//ExistTagByName 根据名称判断标签是否存在
	ExistTagByName(name string) bool
	//AddTag 添加标签
	AddTag(name string, state int, createdBy string) bool
	//DeleteTag 删除标签
	DeleteTag(id int) bool
	//EditTag 编辑标签
	EditTag(id int, data interface{}) bool
}
