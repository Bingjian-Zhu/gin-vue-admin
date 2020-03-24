package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IArticleRepository Article接口定义
type IArticleRepository interface {
	//GetTables 分页返回Articles
	GetTables(PageNum int, PageSize int, maps map[string]interface{}) *[]models.Article
	//GetArticle 根据id获取Article
	GetArticle(id int) *models.Article
	//AddArticle 新增Article
	AddArticle(article *models.Article) bool
	//EditArticle 编辑Article
	EditArticle(article models.Article) bool
	//DeleteArticle 删除Article
	DeleteArticle(id int) bool
	//ExistArticleByID 根据ID判断Article是否存在
	ExistArticleByID(id int) bool
	//GetArticleTotal 获取Article总数
	GetArticleTotal(maps map[string]interface{}) (count int)
	//GetArticles 获取文章
	GetArticles(PageNum int, PageSize int, total *uint64, maps interface{}) *[]models.Article
}
