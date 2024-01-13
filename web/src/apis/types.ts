// 全局接口返回类型
export type PromiseRes<T> = Promise<RequestRes<T>>

export interface RequestRes<T = any> {
    code: number,
    message: string,
    data: T,
}

// 菜单树
export interface MenuTree {
    label?: null | string;
    data?: null | string;
    type?: null | string;
    children?: Array<null | MenuTree> | null;
}

// 接口数据
export interface InterfaceData {
    uid: string;
    method: string;
    name: string;
    test: {
        num: number,
        time: number,
        sleep: number,
    };
    url?: string | null;
    header?: Record<string, any> | null;
    data?: string | null;
}

// 场景数据
export interface TestData {
    uid: string;
    name: string;
    list: Array<null | string> | null;
    test: {
        num: number,
        time: number,
        sleep: number,
    };
}

// 场景数据 - 接口列表展示数据
export interface TestDataList {
    uid: string;
    method: string;
    name: string;
    url: string;
}

// 分页信息
export interface PageInfo {
    total: number,
    current: number,
    size: number,
}

// 软件相关信息
export interface SoftwareInfo {
    doc: string,
    git: string,
}

// 接口编辑器请求参数
export interface MonacoParams {
    header?: Record<string, any> | null;
    param?: Record<string, any> | null;
    body?: Record<string, any> | null;
}

// 测试报告数据
export interface TestReport {
    uid: string;
    status: number;
    total: number;
    success: number;
    fail: number;
    speed_time: number;
    created_at: string;
}

// 测试报告详细信息
export interface TestReportDetails {
    count: number;
    details: TestReportDetailsArray[] | null;
    report: TestReportDetailsReport;
}

export interface TestReportDetailsArray {
    uid: string;
    name: string;
    status: number;
    status_code: number;
    method: string;
    url: string;
    speed_time: number;
}

export interface TestReportDetailsReport {
    fail: number;
    speed_time: number;
    status: number;
    success: number;
    total: number;
}

// 代理池列表
export interface ProxyList {
    uid: string,
    url: string,
}

// 分页信息
export interface PageInfo {
    total: number,
    current: number,
    size: number,
}

// 测试接口请求信息
export interface ReportRequest {
    body: Record<string, any> | null;
    cookie: Record<string, any> | null;
    header: Record<string, any> | null;
}

// 系统更新信息
export interface SystemUpdate {
    log: string;
    msg: string | undefined;
    status: number;
    version: SystemUpdateVersion | null;
}

export interface SystemUpdateVersion {
    client: string;
    release: string;
    server: string;
}
