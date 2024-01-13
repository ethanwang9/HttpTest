package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Default(group *gin.RouterGroup, path string) {
	r := group.Group(path)
	{
		r.GET("/", func(ctx *gin.Context) {
			f, _ := os.ReadFile("./template/index.html")
			ctx.Writer.Write(f)
		})
	}
}
