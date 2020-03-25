package jwt

import (
	"encoding/json"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/bingjian-zhu/gin-vue-admin/common/codes"
	"github.com/bingjian-zhu/gin-vue-admin/common/helper"
	"github.com/bingjian-zhu/gin-vue-admin/common/setting"
	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/bingjian-zhu/gin-vue-admin/service"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

// JWT 注入IService
type JWT struct {
	UserService service.IUserService `inject:""`
	RoleService service.IRoleService `inject:""`
}

//JwtAuthorizator 定义身份授权事件类型
type JwtAuthorizator func(data interface{}, c *gin.Context) bool

//app 程序配置
var app = setting.Config.APP

//GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 5,
		MaxRefresh:  time.Hour,
		IdentityKey: app.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserRole); ok {
				//get roles from username
				v.UserRoles = j.RoleService.GetUserRoles(v.UserName)
				jsonRole, _ := json.Marshal(v.UserRoles)
				//maps the claims in the JWT
				return jwt.MapClaims{
					"userName":  v.UserName,
					"userRoles": helper.B2S(jsonRole),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			roles := jwt.ExtractClaims(c)
			//extracts identity from roles
			jsonRole := roles["userRoles"].(string)
			var userRoles []models.Role
			json.Unmarshal(helper.S2B(jsonRole), &userRoles)
			//Set the identity
			return &models.UserRole{
				UserName:  roles["userName"].(string),
				UserRoles: &userRoles,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals models.User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			if j.UserService.CheckUser(userID, password) {
				return &models.UserRole{
					UserName: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

//AdminAuthorizator role is admin can access
func AdminAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.UserRole); ok {
		for _, itemRole := range *v.UserRoles {
			if itemRole.Value == "admin" {
				return true
			}
		}
	}

	return false
}

//TestAuthorizator username is test can access
func TestAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.UserRole); ok && v.UserName == "test" {
		return true
	}

	return false
}

//AllUserAuthorizator all users access
func AllUserAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

//NoRouteHandler 404 handler
func NoRouteHandler(c *gin.Context) {
	code := codes.PageNotFound
	c.JSON(404, gin.H{"code": code, "message": codes.GetMsg(code)})
}
