package router

import (
	"github.com/gin-gonic/gin"
)

func Test(group *gin.RouterGroup, path string) {
	r := group.Group(path)
	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "success")
		})
	}
}
