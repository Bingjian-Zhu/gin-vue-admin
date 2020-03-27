package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/bingjian-zhu/gin-vue-admin/page"
	"github.com/bingjian-zhu/gin-vue-admin/page/emun"
	"github.com/bingjian-zhu/gin-vue-admin/repository"
)

// ArticleService 注入IArticleRepo
type ArticleService struct {
	Repository repository.IArticleRepository `inject:""`
}

//GetArticle 根据id获取Article
func (a *ArticleService) GetArticle(id int) *models.Article {
	where := models.Article{ID: id}
	return a.Repository.GetArticle(&where)
}

//GetTables 分页返回文章
func (a *ArticleService) GetTables(pageNum, pagesize int) *[]page.Article {
	var (
		pageArticles []page.Article
		pageArticle  page.Article
	)
	where := models.Article{}
	articles := a.Repository.GetTables(pageNum, pagesize, &where)
	for _, article := range *articles {
		pageArticle.ID = article.ID
		pageArticle.Author = article.CreatedBy
		pageArticle.DisplayTime = article.ModifiedOn.String()
		pageArticle.Pageviews = 3474
		pageArticle.Status = emun.GetArticleStatus(article.State)
		pageArticle.Title = article.Title
		pageArticles = append(pageArticles, pageArticle)
	}
	return &pageArticles
}

//AddArticle 新增Article
func (a *ArticleService) AddArticle(article *models.Article) bool {
	return a.Repository.AddArticle(article)
}

//GetArticles 获取文章信息
func (a *ArticleService) GetArticles(PageNum int, PageSize int, total *uint64, where interface{}) *[]models.Article {
	return a.Repository.GetArticles(PageNum, PageSize, total, where)
}

// //ExistArticleByID 根据ID判断Article是否存在
// func (a *ArticleService) ExistArticleByID(id int) bool {
// 	return a.Repository.ExistArticleByID(id)
// }

// //EditArticle 编辑Article
// func (a *ArticleService) EditArticle(article models.Article) bool {
// 	return a.Repository.EditArticle(article)
// }

// //DeleteArticle 删除Article
// func (a *ArticleService) DeleteArticle(id int) bool {
// 	return a.Repository.DeleteArticle(id)
// }
