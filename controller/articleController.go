package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/bingjian-zhu/gin-vue-admin/common/codes"
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/bingjian-zhu/gin-vue-admin/page"
	"github.com/bingjian-zhu/gin-vue-admin/page/emun"
	"github.com/bingjian-zhu/gin-vue-admin/service"
)

// Article 注入IArticleService
type Article struct {
	Service service.IArticleService `inject:""`
	//TagService     service.ITagService     `inject:""`
}

// //GetArticle 获取单个文章
// func (a *Article) GetArticle(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	valid := validation.Validation{}
// 	valid.Min(id, 1, "id").Message("ID必须大于0")
// 	var data models.Article
// 	code := codes.InvalidParams
// 	if !valid.HasErrors() {
// 		if a.Service.ExistArticleByID(id) {
// 			data = a.Service.GetArticle(id)
// 			code = codes.SUCCESS
// 		} else {
// 			code = codes.ErrNotExistArticle
// 		}
// 	} else {
// 		for _, err := range valid.Errors {
// 			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
// 		}
// 	}
// 	res.RespData(c, http.StatusOK, code, &data)
// }

//GetTables 获取多个文章
func (a *Article) GetTables(c *gin.Context) {
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagID int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID, _ = strconv.Atoi(arg)
		maps["tag_id"] = tagID

		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}
	code := codes.InvalidParams
	var viewArticles []page.Article
	var viewArticle page.Article
	if !valid.HasErrors() {
		code = codes.SUCCESS
		page, pagesize := GetPage(c)
		articles := a.Service.GetTables(page, pagesize, maps)
		for _, article := range *articles {
			viewArticle.ID = article.ID
			viewArticle.Author = article.CreatedBy
			viewArticle.DisplayTime = article.ModifiedOn.String()
			viewArticle.Pageviews = 3474
			viewArticle.Status = emun.GetArticleStatus(article.State)
			viewArticle.Title = article.Title
			viewArticles = append(viewArticles, viewArticle)
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, &viewArticles)
}

//AddArticle 新增文章
func (a *Article) AddArticle(c *gin.Context) {
	article := models.Article{}
	code := codes.InvalidParams
	err := c.Bind(&article)
	if err == nil {
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
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
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
