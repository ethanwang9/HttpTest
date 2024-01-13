import request from "./index";
import {PromiseRes, ProxyList, SystemUpdate} from "./types.ts";

// 获取 - 代理池列表
export const GetPanelSystemProxyList = (): PromiseRes<ProxyList[]> => request({
    method: "get",
    url: "/panel/system/proxy",
})

// 创建/更新 - 代理池列表
export const SetPanelSystemProxyList = (data: ProxyList[]): PromiseRes<null> => request({
    method: "post",
    url: "/panel/system/proxy",
    data: {
        proxy: JSON.stringify(data),
    },
})

// 获取 - 系统版本信息
export const SetPanelSystemUpdate = (version: string, hash: string): PromiseRes<SystemUpdate> => request({
    method: "get",
    url: "/panel/system/update",
    params: {
        v: version,
        h: hash,
    },
})