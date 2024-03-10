package panel

import (
	"HttpTest-Server/api"
	"HttpTest-Server/global"
	"HttpTest-Server/model"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"net/http"
	"strconv"
	"strings"
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
	versionRelease := strings.Split(result[0].Version, ".")
	versionServer := strings.Split(global.SysVersion, ".")
	understandVersion := false
	for i := 0; i < 3; i++ {
		r, _ := strconv.Atoi(versionRelease[i])
		s, _ := strconv.Atoi(versionServer[i])
		if s > r {
			understandVersion = true
			break
		} else if s == r {
			continue
		} else if s < r {
			break
		}
	}
	if understandVersion {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusSuccess,
			Message: "success",
			Data: gin.H{
				"status": global.UpdateStatusVersionNotFound,
				"msg":    global.UpdateStatusVersionNotFoundMsg,
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
	// 3. 校验 hash
	if global.SysVersion == result[0].Version {
		// 校验 server hash
		if result[0].Hash.Server != global.SysHash {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.StatusSuccess,
				Message: "success",
				Data: gin.H{
					"status": global.UpdateStatusHashServerNotTrue,
					"msg":    global.UpdateStatusHashServerNotTrueMsg,
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
		// 校验 client hash
		if result[0].Hash.Client != h {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.StatusSuccess,
				Message: "success",
				Data: gin.H{
					"status": global.UpdateStatusHashClientNotTrue,
					"msg":    global.UpdateStatusHashClientNotTrueMsg,
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
	}
	// 更新校验
	hasUpdate := false
	for i := 0; i < 3; i++ {
		r, _ := strconv.Atoi(versionRelease[i])
		s, _ := strconv.Atoi(versionServer[i])
		if r > s {
			hasUpdate = true
			break
		}
	}
	if hasUpdate {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusSuccess,
			Message: "success",
			Data: gin.H{
				"status": global.UpdateStatusFoundNewVersion,
				"msg":    global.UpdateStatusFoundNewVersionMsg,
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
