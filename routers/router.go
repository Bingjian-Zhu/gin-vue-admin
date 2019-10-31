package routers

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"github.com/bingjian-zhu/gin-vue-admin/common/datasource"
	"github.com/bingjian-zhu/gin-vue-admin/common/middleware/cors"

	"github.com/bingjian-zhu/gin-vue-admin/common/middleware/jwt"
	"github.com/bingjian-zhu/gin-vue-admin/common/setting"
	"github.com/bingjian-zhu/gin-vue-admin/controller"
	"github.com/bingjian-zhu/gin-vue-admin/repository"
	"github.com/bingjian-zhu/gin-vue-admin/service"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config.APP.RunMode)
	Configure(r)
	return r
}

//Configure 配置router
func Configure(r *gin.Engine) {
	//controller declare
	var user controller.User
	//var tag cv1.Tag
	var myjwt jwt.JWT
	//inject declare
	var article controller.Article
	db := datasource.Db{}
	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &article},
		&inject.Object{Value: &db},
		&inject.Object{Value: &repository.ArticleRepository{}},
		&inject.Object{Value: &service.ArticleService{}},
		&inject.Object{Value: &user},
		&inject.Object{Value: &repository.UserRepository{}},
		//&inject.Object{Value: &service.UserService{}},
		&inject.Object{Value: &repository.ClaimRepository{}},
		//&inject.Object{Value: &service.ClaimService{}},
		&inject.Object{Value: &myjwt},
		&inject.Object{Value: &repository.BaseRepository{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
	//database connect
	err = db.Connect()
	if err != nil {
		log.Fatal("db fatal:", err)
	}
	var authMiddleware = myjwt.GinJWTMiddlewareInit(jwt.AllUserAuthorizator)
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/info", user.GetUserInfo)
		userAPI.POST("/logout", user.Logout)
	}

	var adminMiddleware = myjwt.GinJWTMiddlewareInit(jwt.AdminAuthorizator)
	apiv1 := r.Group("/api/v1")
	//使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	apiv1.Use(adminMiddleware.MiddlewareFunc())
	{
		//vue获取table信息
		apiv1.GET("/table/list", article.GetArticles)

		// apiv1.GET("/articles", article.GetArticles)
		// apiv1.GET("/articles/:id", article.GetArticle)
		// apiv1.POST("/articles", article.AddArticle)
		// apiv1.PUT("/articles/:id", article.EditArticle)
		// apiv1.DELETE("/articles/:id", article.DeleteArticle)
	}
}
