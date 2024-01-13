package global

// 系统相关

const (
	// SysPort 端口号
	SysPort = 8888
	// SysMode 软件模式
	// debug | release
	SysMode = "debug"
	// SysLink 更新地址
	SysLink = "https://raw.cachefly.998111.xyz/ethanwang9/HttpTest/main/version.json"
	// SysVersion 软件版本号
	SysVersion = "1.0.0"
	// SysHash 系统版本Hash
	SysHash = "3329C449DB724EE1B5E838EA50D5DFC5"
)

// 状态码

const (
	// StatusSuccess 成功
	StatusSuccess = 0
	// StatusErrorSys 系统错误
	StatusErrorSys = 100
	// StatusErrorSQL 数据库错误
	StatusErrorSQL = 200
	// StatusErrorUserDo 用户错误
	StatusErrorUserDo = 300
)

// 数据库

const (
	// DbPathFile 数据库文件位置
	DbPathFile = "./db.db"
)

// 日志

const (
	// LogPath 日志目录
	LogPath = "./log"
	// LogInConsole 日志是否显示在控制台
	LogInConsole = true
	// LogPrefix 日志前缀
	LogPrefix = ""
	// LogFormat 日志显示方式
	// json | console
	LogFormat = "json"
	// LogStacktraceKey 日志栈名
	LogStacktraceKey = "stacktrace"
	// LogEncodeLevel 编码级
	// LowercaseLevelEncoder 小写编码器(默认) | LowercaseColorLevelEncoder 小写编码器带颜色
	// CapitalLevelEncoder 大写编码器 | CapitalColorLevelEncoder 大写编码器带颜色
	LogEncodeLevel = "CapitalLevelEncoder"
	// LogLevel 日志级别
	// debug | info | warn | error | dpanic | panic | fatal
	LogLevel = "info"
)

// 缓存

const (
	// CacheTempPath 临时文件夹
	CacheTempPath = "./temp"
)

// 接口状态

const (
	// InterfaceStatusFail 接口运行失败
	InterfaceStatusFail = -2
	// InterfaceStatusStop 接口暂停运行
	InterfaceStatusStop = -1
	// InterfaceStatusRun 接口运行中
	InterfaceStatusRun = 0
	// InterfaceStatusSuccess 接口运行成功
	InterfaceStatusSuccess = 1
)

// 更新状态

const (
	// UpdateStatusSuccess 正常
	UpdateStatusSuccess = 0

	// UpdateStatusVersionNotFound 软件版本号>最新版本号，非官方发行版本号
	UpdateStatusVersionNotFound    = 1
	UpdateStatusVersionNotFoundMsg = "软件未通过安全校验，未知渠道发行版本，请注意个人数字资产安全！"

	// UpdateStatusVersionNotTrue 客户端版本号!=服务端版本号
	UpdateStatusVersionNotTrue    = 2
	UpdateStatusVersionNotTrueMsg = "软件未通过安全校验，客户端与服务端版本号不一致，会有兼容问题！"

	// UpdateStatusHashClientNotTrue 客户端HASH不对，客户端疑似被篡改
	UpdateStatusHashClientNotTrue    = 3
	UpdateStatusHashClientNotTrueMsg = "软件未通过安全校验，客户端疑似被篡改，请注意个人数字资产安全！"

	// UpdateStatusHashServerNotTrue 服务端HASH不对，服务端疑似被篡改
	UpdateStatusHashServerNotTrue    = 4
	UpdateStatusHashServerNotTrueMsg = "软件未通过安全校验，服务端疑似被篡改，请注意个人数字资产安全！"

	// UpdateStatusFoundNewVersion 软件版本号<最新版本号，发现新版本
	UpdateStatusFoundNewVersion    = 5
	UpdateStatusFoundNewVersionMsg = "发现新版本，请及时前往更新!"
)
