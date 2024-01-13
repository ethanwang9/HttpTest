package panel

import (
	"HttpTest-Server/global"
	"HttpTest-Server/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTestDir 获取 - 自动化测试文件树
func GetTestDir(ctx *gin.Context) {
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
	json.Unmarshal([]byte(result.TestMenu), &tree)

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data:    tree,
	})
}

// SetTestDir 设置 - 自动化测试文件树
func SetTestDir(ctx *gin.Context) {
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
	d.TestMenu = tree
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

// SetTestInfo 创建 - 场景
func SetTestInfo(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	name := ctx.PostForm("name")

	if err := model.TestDataApp.New(model.TestData{
		UID:  uid,
		Name: name,
		Test: []byte(`{"num":1,"time":1,"sleep":0}`),
	}).Set(); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库设置场景相关信息失败",
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

// DelTestInfo 删除 - 场景
func DelTestInfo(ctx *gin.Context) {
	testData := model.TestDataApp.New(model.TestData{})
	if err := ctx.ShouldBindJSON(&testData); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSys,
			Message: "解析提交数据失败",
			Data:    nil,
		})
		return
	}

	if err := testData.Delete(); err != nil {
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

// GetTestInfo 获取 - 场景
func GetTestInfo(ctx *gin.Context) {
	uid := ctx.Query("uid")
	data, err := model.TestDataApp.New(model.TestData{UID: uid}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "数据库获取场景数据失败",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data: gin.H{
			"uid":  data.UID,
			"name": data.Name,
			"list": data.InterfaceList,
			"test": data.Test,
		},
	})
}

// UpdateTestInfo 更新 - 场景
func UpdateTestInfo(ctx *gin.Context) {
	testData := model.TestDataApp.New(model.TestData{})
	if err := ctx.ShouldBindJSON(&testData); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSys,
			Message: "数据解析失败",
			Data:    nil,
		})
		return
	}

	if err := testData.Update(); err != nil {
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

// GetTestReport 获取 - 测试报告
func GetTestReport(ctx *gin.Context) {
	uid := ctx.Query("uid")

	report := model.ReportApp.New(model.Report{TestUid: uid})
	result, err := report.Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "从数据库获取数据失败",
			Data:    nil,
		})
		return
	}

	data := make([]gin.H, 0)
	for _, r := range result {
		data = append(data, gin.H{
			"uid":        r.UID,
			"status":     r.Status,
			"total":      r.Total,
			"success":    r.Success,
			"fail":       r.Fail,
			"speed_time": r.SpeedTime,
			"created_at": r.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data:    data,
	})
}

// DelTestReport 获取 - 测试报告
func DelTestReport(ctx *gin.Context) {
	report := model.ReportApp.New(model.Report{})
	if err := ctx.ShouldBindJSON(&report); err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSys,
			Message: "解析提交数据失败",
			Data:    nil,
		})
		return
	}

	if err := report.Delete(); err != nil {
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

// GetTestReportDetails 获取 - 测试报告详细信息
func GetTestReportDetails(ctx *gin.Context) {
	uid := ctx.Query("uid")
	size, _ := strconv.Atoi(ctx.Query("size"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	// 获取测试报告详细数据
	result, err := model.ReportDataApp.New(model.ReportData{ReportUid: uid}).Get(size, page)
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "获取数据库信息错误",
			Data:    nil,
		})
		return
	}

	// 获取测试报告总数
	count, err := model.ReportDataApp.New(model.ReportData{ReportUid: uid}).Count()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "获取数据库信息错误",
			Data:    nil,
		})
		return
	}

	// 获取测试报告信息
	report, err := model.ReportApp.New(model.Report{UID: uid}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "获取数据库信息错误",
			Data:    nil,
		})
		return
	}

	content := make([]gin.H, 0)
	for _, data := range result {
		content = append(content, gin.H{
			"uid":         data.UID,
			"status":      data.Status,
			"method":      data.Method,
			"url":         data.Url,
			"name":        data.Name,
			"status_code": data.StatusCode,
			"speed_time":  data.SpeedTime,
		})
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data: gin.H{
			"count": count,
			"report": gin.H{
				"status":     report[0].Status,
				"total":      report[0].Total,
				"success":    report[0].Success,
				"fail":       report[0].Fail,
				"speed_time": report[0].SpeedTime,
			},
			"details": content,
		},
	})
}

// GetTestReportRequest 获取 - 测试接口请求信息
func GetTestReportRequest(ctx *gin.Context) {
	uid := ctx.Query("uid")

	result, err := model.ReportDataApp.New(model.ReportData{UID: uid}).Get2()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.StatusErrorSQL,
			Message: "获取数据库信息失败",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.StatusSuccess,
		Message: "success",
		Data: gin.H{
			"header": result.Header,
			"body":   result.Body,
			"cookie": result.Cookie,
		},
	})
}
