package controller

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/bingjian-zhu/gin-vue-admin/common/codes"
	"github.com/bingjian-zhu/gin-vue-admin/common/logger"
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/bingjian-zhu/gin-vue-admin/service"
)

// Article 注入IArticleService
type Article struct {
	Log     logger.ILogger          `inject:""`
	Service service.IArticleService `inject:""`
}

//GetArticle 获取单个文章
func (a *Article) GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Article
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = a.Service.GetArticle(id)
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

//GetTables 获取多个文章
func (a *Article) GetTables(c *gin.Context) {
	code := codes.SUCCESS
	page, pagesize := GetPage(c)
	data := a.Service.GetTables(page, pagesize)
	RespData(c, http.StatusOK, code, data)
}

//AddArticle 新增文章
func (a *Article) AddArticle(c *gin.Context) {
	article := models.Article{}
	code := codes.InvalidParams
	err := c.Bind(&article)
	if err == nil {
		article.ModifiedOn = article.CreatedOn
		valid := validation.Validation{}
		valid.Min(article.TagID, 0, "tag_id").Message("标签ID必须不小于0")
		valid.Required(article.Title, "title").Message("标题不能为空")
		valid.Required(article.Desc, "desc").Message("简述不能为空")
		valid.Required(article.Content, "content").Message("内容不能为空")
		valid.Required(article.CreatedBy, "created_by").Message("创建人不能为空")
		valid.Range(article.State, 0, 1, "state").Message("状态只允许0或1")
		if !valid.HasErrors() {
			if a.Service.AddArticle(&article) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//GetArticles 获取文章信息
func (a *Article) GetArticles(c *gin.Context) {
	res := make(map[string]interface{}, 2)
	var total uint64
	code := codes.SUCCESS
	page, pagesize := GetPage(c)
	articles := a.Service.GetArticles(page, pagesize, &total, "")
	res["list"] = &articles
	res["total"] = total
	RespData(c, http.StatusOK, code, &res)
}

// //EditArticle 修改文章
// func (a *Article) EditArticle(c *gin.Context) {
// 	model := models.Article{}
// 	code := codes.InvalidParams
// 	err := c.Bind(&model)
// 	if err == nil {
// 		valid := validation.Validation{}
// 		valid.Min(model.ID, 1, "id").Message("ID必须大于0")
// 		valid.MaxSize(model.Title, 100, "title").Message("标题最长为100字符")
// 		valid.MaxSize(model.Desc, 255, "desc").Message("简述最长为255字符")
// 		valid.MaxSize(model.Content, 65535, "content").Message("内容最长为65535字符")
// 		valid.Required(model.ModifiedBy, "modified_by").Message("修改人不能为空")
// 		valid.MaxSize(model.ModifiedBy, 100, "modified_by").Message("修改人最长为100字符")
// 		valid.Range(model.State, 0, 1, "state").Message("状态只允许0或1")
// 		if !valid.HasErrors() {
// 			if a.ArticleService.ExistArticleByID(model.ID) {
// 				if a.TagService.ExistTagByID(model.TagID) {
// 					a.ArticleService.EditArticle(model)
// 					code = codes.SUCCESS
// 				} else {
// 					code = codes.ErrNotExistTag
// 				}
// 			} else {
// 				code = codes.ErrNotExistArticle
// 			}
// 		} else {
// 			for _, err := range valid.Errors {
// 				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
// 			}
// 		}
// 	}
// 	res.RespOk(c, http.StatusOK, code)
// }

// //DeleteArticle 删除文章
// func (a *Article) DeleteArticle(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	valid := validation.Validation{}
// 	valid.Min(id, 1, "id").Message("ID必须大于0")

// 	code := codes.InvalidParams
// 	if !valid.HasErrors() {
// 		if a.Service.ExistArticleByID(id) {
// 			a.Service.DeleteArticle(id)
// 			code = codes.SUCCESS
// 		} else {
// 			code = codes.ErrNotExistArticle
// 		}
// 	} else {
// 		for _, err := range valid.Errors {
// 			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
// 		}
// 	}
// 	res.RespOk(c, http.StatusOK, code)
// }
