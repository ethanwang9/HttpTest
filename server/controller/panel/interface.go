package panel

import (
	"HttpTest-Server/global"
	"HttpTest-Server/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetInterfaceDir 获取 - 接口管理文件数
func GetInterfaceDir(ctx *gin.Context) {
	result, err := model.SystemApp.Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库获取系统相关信息失败",
			Data:    nil,
		})
		return
	}

	// 解码返回
	var tree []global.MenuTree
	_ = json.Unmarshal([]byte(result.InterfaceMenu), &tree)

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data:    tree,
	})
}

// SetInterfaceDir 设置 - 接口管理文件数
func SetInterfaceDir(ctx *gin.Context) {
	// 获取参数
	tree := ctx.PostForm("tree")

	// 更新参数
	d, err := model.SystemApp.Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库获取系统相关信息失败",
			Data:    nil,
		})
		return
	}
	d.InterfaceMenu = tree
	err = d.Update()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库设置系统相关信息失败",
			Data:    nil,
		})
		return
	}

	// 返回参数
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data:    nil,
	})
}

// SetInterfaceInfo 设置 - 创建接口
func SetInterfaceInfo(ctx *gin.Context) {
	err := model.InterfaceDataApp.New(model.InterfaceData{
		UID:    ctx.PostForm("uid"),
		Method: "get",
		Name:   ctx.PostForm("name"),
		Test:   []byte(`{"num":1,"time":1,"sleep":0}`),
		Url:    "https://httpbin.org/ip",
	}).Set()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库设置接口数据失败",
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

// GetInterfaceInfo 获取 - 接口信息
func GetInterfaceInfo(ctx *gin.Context) {
	uid := ctx.Query("uid")
	data, err := model.InterfaceDataApp.New(model.InterfaceData{UID: uid}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库获取接口数据失败",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data: gin.H{
			"uid":    data.UID,
			"method": data.Method,
			"url":    data.Url,
			"name":   data.Name,
			"header": data.Header,
			"test":   data.Test,
			"data":   data.Data,
		},
	})
}

// UpdateInterfaceInfo 更新 - 接口
func UpdateInterfaceInfo(ctx *gin.Context) {
	interfaceData := model.InterfaceDataApp.New(model.InterfaceData{})
	if err := ctx.ShouldBindJSON(&interfaceData); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSys,
			Message: "数据解析失败",
			Data:    nil,
		})
		return
	}

	if err := interfaceData.Update(); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "保存数据到数据库失败",
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

// DelInterfaceInfo 删除 - 接口
func DelInterfaceInfo(ctx *gin.Context) {
	interfaceData := model.InterfaceDataApp.New(model.InterfaceData{})
	if err := ctx.ShouldBindJSON(&interfaceData); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSys,
			Message: "解析提交数据失败",
			Data:    nil,
		})
		return
	}

	if err := interfaceData.Delete(); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "删除数据失败",
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
