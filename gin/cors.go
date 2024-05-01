package gin

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// Cors Gin 跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		wg := errgroup.Group{}
		wg.Go(func() error {
			method := c.Request.Method
			origin := c.Request.Header.Get("Origin")
			if origin != "" {
				c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
				c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
				c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
				c.Header("Access-Control-Allow-Credentials", "true")
			}
			if method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
			c.Next()
			return nil
		})
		err := wg.Wait()
		if err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
