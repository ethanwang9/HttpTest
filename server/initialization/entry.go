package initialization

import (
	"HttpTest-Server/global"
	"HttpTest-Server/service"
	"HttpTest-Server/utils"
	"os"
)

// HTInit 初始化
func HTInit() {
	// 初始化服务
	global.APP_LOG = service.Zap()
	global.APP_DB = service.Gorm()

	// 初始化数据库
	gormInit()

	// 初始化缓存目录
	cacheTemp()
}

func cacheTemp() {
	// 判断是否有 temp 文件夹
	if ok, _ := utils.PathExists(global.CacheTempPath); !ok {
		global.APP_LOG.Info("[初始化] 创建临时文件目录 temp")
		_ = os.Mkdir(global.CacheTempPath, os.ModePerm)
	}
}
