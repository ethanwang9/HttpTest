import type {RouteRecordRaw} from 'vue-router'
import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "Index",
        meta: {
            title: "控制面板",
        },
        redirect: {name: "Interface"},
        children: [
            {
                path: "interface",
                name: "Interface",
                meta: {
                    title: "接口管理",
                },
                component: () => import("@/views/interface.vue"),
                children: [
                    {
                        path: "detail",
                        name: "InterfaceDetail",
                        meta: {
                            title: "接口数据详情信息",
                        },
                        component: () => import("@/views/detail.vue"),
                    },
                ]
            },
            {
                path: "test",
                name: "Test",
                meta: {
                    title: "自动化测试",
                },
                component: () => import("@/views/test.vue"),
                children: [
                    {
                        path: "report",
                        name: "Report",
                        meta: {
                            title: "测试报告",
                        },
                        component: () => import("@/views/report.vue"),
                    },
                    {
                        path: "detail",
                        name: "TestDetail",
                        meta: {
                            title: "接口数据详情信息",
                        },
                        component: () => import("@/views/detail.vue"),
                    },
                ]
            },
            {
                path: "system",
                name: "System",
                meta: {
                    title: "系统设置",
                },
                component: () => import("@/views/system.vue"),
            },
        ]
    },
    {
        path: "/404",
        name: "404",
        meta: {
            title: "页面不存在",
        },
        component: () => import("@/views/404.vue"),
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.beforeEach((to, _, next) => {
    // flag
    let flag = true

    // 动态设置网页标题
    window.document.title = to.meta.title as string || 'HttpTest | 自动化批量测试服务'

    // 404
    if (!router.getRoutes().find((v => v.path === to.path))) {
        flag = false
        next({name: '404'})
    }

    // 请求正常，返回正确页面
    if (flag) {
        next()
    }
})

export default router