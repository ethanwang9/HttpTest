import request from "./index";
import {PromiseRes, SoftwareInfo} from "./types.ts";

// 获取面板基本信息
export const GetSystemBasic = (): PromiseRes<SoftwareInfo> => request({
    method: "get",
    url: "/sys/basic",
})