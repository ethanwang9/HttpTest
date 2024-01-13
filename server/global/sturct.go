package global

// MsgBack 消息返回接口
type MsgBack struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// MenuTree 菜单树
type MenuTree struct {
	Type     string     `json:"type,omitempty"`
	Label    string     `json:"label,omitempty"`
	Data     string     `json:"data,omitempty"`
	Children []MenuTree `json:"children,omitempty"`
}
