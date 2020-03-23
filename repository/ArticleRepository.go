package repository

import (
	"fmt"

	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/common/logger"
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//ArticleRepository 注入IDb
type ArticleRepository struct {
	Source datasource.IDb `inject:""`
	Log    logger.ILogger `inject:""`
	Base   BaseRepository `inject:"inline"`
}

//GetTables 分页返回Articles
func (a *ArticleRepository) GetTables(PageNum int, PageSize int, maps map[string]interface{}) *[]models.Article {
	var articles []models.Article
	var total uint64
	err := a.Base.GetPages(&models.Article{}, &articles, PageNum, PageSize, &total, maps)
	if err != nil {
		fmt.Println(err)
	}
	return &articles
}

//GetArticle 根据id获取Article
func (a *ArticleRepository) GetArticle(id int) (article models.Article) {
	a.Source.DB().Where("id = ?", id).First(&article)
	a.Source.DB().Where("id = ?", article.TagID).First(&article.Tag)
	//db.Model(&article).Related(&article.Tag)
	return
}

//AddArticle 新增Article
func (a *ArticleRepository) AddArticle(article *models.Article) bool {
	a.Source.DB().Create(article)
	return true
}

//EditArticle 编辑Article
func (a *ArticleRepository) EditArticle(article models.Article) bool {
	a.Source.DB().Model(&models.Article{}).Where("id = ?", article.ID).Update(article)
	return true
}

//DeleteArticle 删除Article
func (a *ArticleRepository) DeleteArticle(id int) bool {
	a.Source.DB().Where("id = ?", id).Delete(&models.Article{})
	return true
}

//ExistArticleByID 根据ID判断Article是否存在
func (a *ArticleRepository) ExistArticleByID(id int) bool {
	var article models.Article
	a.Source.DB().Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

//GetArticleTotal 获取Article总数
func (a *ArticleRepository) GetArticleTotal(maps map[string]interface{}) (count int) {
	a.Source.DB().Model(&models.Article{}).Where(maps).Count(&count)
	return
}

//GetArticles 获取文章
func (a *ArticleRepository) GetArticles(PageNum int, PageSize int, total *uint64, maps interface{}) *[]models.Article {
	var articles []models.Article
	err := a.Base.GetPages(&models.Article{}, &articles, PageNum, PageSize, total, maps)
	if err != nil {
		a.Log.Errorf("获取文章信息失败", err)
	}
	return &articles
}
