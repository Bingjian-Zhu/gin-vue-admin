package controller

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/bingjian-zhu/gin-vue-admin/common/codes"
	"github.com/bingjian-zhu/gin-vue-admin/service"
)

//Tag 注入ITagService
type Tag struct {
	Service service.ITagService `inject:""`
}

//GetTags 获取多个文章标签
func (t *Tag) GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state
	}

	code := codes.SUCCESS
	data = t.Service.GetTags(GetPage(c), maps)
	RespData(c, http.StatusOK, code, &data)
}

//AddTag 新增文章标签
func (t *Tag) AddTag(c *gin.Context) {
	name := c.Query("name")
	state, _ := strconv.Atoi(c.DefaultQuery("state", "0"))
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := codes.InvalidParams
	if !valid.HasErrors() {
		if !t.Service.ExistTagByName(name) {
			code = codes.SUCCESS
			t.Service.AddTag(name, state, createdBy)
		} else {
			code = codes.ErrExistTag
		}
	}
	RespOk(c, http.StatusOK, code)
}

//EditTag 修改文章标签
func (t *Tag) EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := codes.InvalidParams
	if !valid.HasErrors() {
		code = codes.SUCCESS
		if t.Service.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			t.Service.EditTag(id, data)
		} else {
			code = codes.ErrNotExistTag
		}
	}
	RespOk(c, http.StatusOK, code)
}

//DeleteTag 删除文章标签
func (t *Tag) DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := codes.InvalidParams
	if !valid.HasErrors() {
		code = codes.SUCCESS
		if t.Service.ExistTagByID(id) {
			t.Service.DeleteTag(id)
		} else {
			code = codes.ErrNotExistTag
		}
	}
	RespOk(c, http.StatusOK, code)
}
