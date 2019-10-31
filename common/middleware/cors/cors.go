package cors

import (
	"github.com/gin-gonic/gin"
)

// CorsHandler 跨域请求处理
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
		c.Header("Access-Control-Allow-Credentials", "true")      // 跨域请求是否需要带cookie信息 默认设置为true
		//  header的类型
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Max-Age", "21600") //可以缓存预检请求结果的时间（以秒为单位）
		c.Set("content-type", "application/json")   // 设置返回格式是json
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
			////放行所有OPTIONS方法，本项目直接返回204
			//c.JSON(200, "Options Request!")
		}

		c.Next()
	}
}
