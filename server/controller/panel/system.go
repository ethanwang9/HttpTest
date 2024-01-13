package panel

import (
	"HttpTest-Server/api"
	"HttpTest-Server/global"
	"HttpTest-Server/model"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"net/http"
)

// SetSystemProxy 创建/更新 - 系统代理
func SetSystemProxy(ctx *gin.Context) {
	list := ctx.PostForm("proxy")

	err := model.SystemApp.New(model.System{ProxyList: datatypes.JSON(list)}).Update()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "设置代理池失败",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data:    nil,
	})
}

// GetSystemProxy 获取 - 系统代理
func GetSystemProxy(ctx *gin.Context) {
	s, err := model.SystemApp.New(model.System{}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "获取系统代理池列表失败",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data:    s.ProxyList,
	})
}

// GetSystemUpdate 获取 - 获取系统版本信息
func GetSystemUpdate(ctx *gin.Context) {
	// 获取参数
	v := ctx.Query("v")
	h := ctx.Query("h")

	// 获取更新数据
	result, err := api.New.Ht.Update.Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSys,
			Message: "系统获取更新信息失败",
			Data:    nil,
		})
		return
	}

	// 判断参数是否合法
	if len(v) == 0 || len(h) == 0 {
		t := "未知版本号"
		if len(v) != 0 {
			t = v
		}
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusSuccess,
			Message: "success",
			Data: gin.H{
				"status": global.UpdateStatusHashClientNotTrue,
				"msg":    global.UpdateStatusHashClientNotTrueMsg,
				"log":    result[0].Log,
				"version": gin.H{
					"release": result[0].Version,
					"client":  t,
					"server":  global.SysVersion,
				},
			},
		})
		return
	}

	// 安全校验
	//versionRelease := strings.Split(result[0].Version, ".")
	//versionClient := strings.Split(v, ".")
	//versionServer := strings.Split(global.SysVersion, ".")
	// 1. 校验客户端版本号与服务端版本号是否一致
	if v != global.SysVersion {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusSuccess,
			Message: "success",
			Data: gin.H{
				"status": global.UpdateStatusVersionNotTrue,
				"msg":    global.UpdateStatusVersionNotTrueMsg,
				"log":    result[0].Log,
				"version": gin.H{
					"release": result[0].Version,
					"client":  v,
					"server":  global.SysVersion,
				},
			},
		})
		return
	}
	// 2. 校验软件版本号
	//fmt.Println(versionRelease, versionClient, versionServer)

	// 3. 校验 client hash
	// 4. 校验 server hash

	// 更新校验

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data: gin.H{
			"status": global.UpdateStatusSuccess,
			"msg":    "",
			"log":    result[0].Log,
			"version": gin.H{
				"release": result[0].Version,
				"client":  v,
				"server":  global.SysVersion,
			},
		},
	})
}
