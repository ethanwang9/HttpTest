<script setup lang="ts">
import MonacoEditor from "@/components/MonacoEditor.vue";
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {ReportRequest} from "@/apis/types.ts";
import {GetPanelTestReportRequest} from "@/apis/panel_test.ts";
import JsonFormat from "@/utils/json-format.ts";

// ==========
// 路由
// ==========

// 路由
const router = useRouter()

// tabs 下标
const tabsIndex = ref("body")

// query 测试报告详细记录ID
const {id} = router.currentRoute.value.query

// 返回上一页
const goBackUpPage = () => {
  router.go(-1)
}

// ==========
// 编辑器
// ==========

// 代码编辑器值
const editValue = reactive<ReportRequest>({
  body: "",
  cookie: "",
  header: "",
})

// 获取请求详细信息
const getRequestInfo = async () => {
  await GetPanelTestReportRequest(id).then((data: ReportRequest) => {
    for (let key in data) {
      editValue[key] = "!@#@" + JsonFormat(JSON.stringify(data[key]))
    }
  }).catch((err: string) => {
    ElMessage({
      message: "获取测试请求信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}
getRequestInfo()
</script>

<template>
  <el-page-header content="测试接口详细信息" @back="goBackUpPage"/>
  <el-tabs v-model="tabsIndex" class="tabs">
    <el-tab-pane label="Body" name="body">
      <MonacoEditor v-model="editValue.body" class="tabs-item"></MonacoEditor>
    </el-tab-pane>
    <el-tab-pane label="Cookie" name="cookie">
      <MonacoEditor v-model="editValue.cookie" class="tabs-item"></MonacoEditor>
    </el-tab-pane>
    <el-tab-pane label="Header" name="header">
      <MonacoEditor v-model="editValue.header" class="tabs-item"></MonacoEditor>
    </el-tab-pane>
  </el-tabs>
</template>

<style scoped lang="less">
.tabs {
  width: 100%;
  height: 100%;

  &-item {
    height: calc(100vh - 15px - 24px - 40px - 15px - 15px);
  }
}
</style>