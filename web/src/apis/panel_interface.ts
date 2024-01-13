import request from "./index";
import {InterfaceData, MenuTree, PromiseRes} from "./types.ts";

// 获取 - 文件树
export const GetPanelInterfaceDir = (): PromiseRes<MenuTree[]> => request({
    method: "get",
    url: "/panel/interface/dir",
})

// 创建/更新 - 文件树
export const SetPanelInterfaceDir = (data: MenuTree[]): PromiseRes<null> => request({
    method: "post",
    url: "/panel/interface/dir",
    data: {
        tree: data,
    },
})

// 创建 - 接口
export const SetPanelInterfaceInfo = (data: string, name: string): PromiseRes<null> => request({
    method: "post",
    url: "/panel/interface/info",
    data: {
        uid: data,
        name: name,
    },
})

// 获取 - 接口
export const GetPanelInterfaceInfo = (uid: string): PromiseRes<InterfaceData> => request({
    method: "get",
    url: "/panel/interface/info",
    params: {
        uid: uid,
    }
})

// 更新 - 接口
export const UpdatePanelInterfaceInfo = (data: InterfaceData): PromiseRes<null> => request({
    method: "put",
    url: "/panel/interface/info",
    data: {
        ...data,
    }
})

// 删除 - 接口
export const DelPanelInterfaceInfo = (data: string): PromiseRes<null> => request({
    method: "delete",
    url: "/panel/interface/info",
    data: {
        uid: data
    }
})