package api

import "HttpTest-Server/api/ht"

type Api struct {
	Ht ht.Ht
}

// New Api实例
var New = new(Api)
