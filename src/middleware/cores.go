package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			// 接收客户端发送的origin (重要!)
			c.Header("Access-Control-Allow-Origin", origin)
			// 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, HEAD, PUT")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,content-type,Authorization,Content-Length,X-CSRF-Token,Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// 允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
			// c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,Access-Control-Allow-Origin")
		}

		// 允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "options ok")
			return
		}
		c.Next()
	}
}
