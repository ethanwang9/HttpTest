import request from "./index";
import {MenuTree, PromiseRes, ReportRequest, TestData, TestReport, TestReportDetails} from "./types.ts";

// 获取 - 文件树
export const GetPanelTestDir = (): PromiseRes<MenuTree[]> => request({
    method: "get",
    url: "/panel/test/dir",
})

// 创建/更新 - 文件树
export const SetPanelTestDir = (data: MenuTree[]): PromiseRes<null> => request({
    method: "post",
    url: "/panel/test/dir",
    data: {
        tree: data,
    },
})

// 创建 - 场景
export const SetPanelTestInfo = (uid: string, name: string): PromiseRes<null> => request({
    method: "post",
    url: "/panel/test/info",
    data: {
        uid: uid,
        name: name,
    },
})

// 删除 - 场景
export const DelPanelTestInfo = (data: string): PromiseRes<null> => request({
    method: "delete",
    url: "/panel/test/info",
    data: {
        uid: data
    }
})

// 获取 - 场景
export const GetPanelTestInfo = (data: string): PromiseRes<TestData> => request({
    method: "get",
    url: "/panel/test/info",
    params: {
        uid: data,
    }
})

// 更新 - 场景
export const UpdatePanelTestInfo = (data: TestData): PromiseRes<null> => request({
    method: "put",
    url: "/panel/test/info",
    data: {
        ...data,
    }
})

// 获取 - 测试报告
export const GetPanelTestReport = (test_uid: string): PromiseRes<TestReport[]> => request({
    method: "get",
    url: "/panel/test/report",
    params: {
        "uid": test_uid,
    }
})

// 删除 - 测试报告
export const DelPanelTestReport = (data: string): PromiseRes<null> => request({
    method: "delete",
    url: "/panel/test/report",
    data: {
        uid: data
    }
})

// 获取 - 测试报告详细信息
export const GetPanelTestReportDetails = (uid: string, size: string, page: string): PromiseRes<TestReportDetails> => request({
    method: "get",
    url: "/panel/test/report/details",
    params: {
        "uid": uid,
        "size": size,
        "page": page,
    }
})

// 获取 - 获取测试接口请求信息
export const GetPanelTestReportRequest = (uid: string): PromiseRes<ReportRequest> => request({
    method: "get",
    url: "/panel/test/report/request",
    params: {
        "uid": uid,
    }
})