package ht

import (
	"HttpTest-Server/global"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// Update 更新信息
type Update struct {
	Version string `json:"version"`
	Hash    struct {
		Server string `json:"server"`
		Client string `json:"client"`
	} `json:"hash"`
	Log string `json:"log"`
}

// Get 获取更新信息
func (u Update) Get() (data []Update, err error) {
	result, err := resty.New().R().Get(global.SysLink)
	if err != nil {
		global.APP_LOG.Error("[接口] 获取HT更新接口失败", zap.Error(err))
		return
	}

	if result.StatusCode() != 200 {
		global.APP_LOG.Error("[接口] 获取HT更新接口失败", zap.Error(err))
		return
	}

	err = json.Unmarshal(result.Body(), &data)
	if err != nil {
		global.APP_LOG.Error("[接口] 获取HT更新接口, 解析json失败", zap.Error(err))
		return
	}

	return
}
