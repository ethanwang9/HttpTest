package sys

import (
	"HttpTest-Server/global"
	"HttpTest-Server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBasic(ctx *gin.Context) {
	result, err := model.SystemApp.Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库获取系统相关信息失败",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data: gin.H{
			"git": result.GitUrl,
			"doc": result.DocUrl,
		},
	})
}
