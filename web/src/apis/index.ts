import axios from "axios";
import {RequestRes} from "./types.ts";

const request = axios.create({
    baseURL: import.meta.env.VITE_SERVER_PATH,
})


// 请求拦截 - 发送前
request.interceptors.request.use(function (config) {
    // 修改请求头
    config.headers!["Content-Type"] = "application/x-www-form-urlencoded"

    // 序列化请求数据
    let method = config.method as string
    if (method.toLowerCase() === "put" || method.toLowerCase() === "delete") {
        config.data = JSON.stringify(config.data)
    }

    return config;
}, function (error) {
    return Promise.reject(error);
})

// 请求拦截- 接收后
request.interceptors.response.use(function (response) {
    // 判断返回接口状态码
    const result = response.data as RequestRes<null>
    if (result.code !== 0) {
        return Promise.reject(result?.message)
    }

    return result?.data;
}, function (error) {
    return Promise.reject(error);
})

export default request
