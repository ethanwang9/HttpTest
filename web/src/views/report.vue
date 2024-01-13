<script setup lang="ts">
import {useRouter} from "vue-router";
import {reactive} from "vue";
import {GetPanelTestReportDetails} from "@/apis/panel_test.ts";
import {PageInfo, TestReportDetails, TestReportDetailsArray, TestReportDetailsReport} from "@/apis/types.ts";
import {convertTime} from "@/utils/tools.ts";
import {clearInterval} from "node:timers";

// ==========
// 路由
// ==========

// 路由
const router = useRouter()

// query 测试报告详细记录ID
const {id} = router.currentRoute.value.query

// 返回上一页
const goBackUpPage = () => {
  router.back()
}

// 跳转查看详情页
const goTestDetailPage = (id: string) => {
  router.push({
    name: "TestDetail",
    query: {
      id,
    }
  })
}

// ==========
// 分页
// ==========

// 分页信息
const pageInfo = reactive<PageInfo>({
  // 记录数
  total: 0,
  // 当前页数
  current: 1,
  // 一页显示记录数
  size: 15,
})

// 页码改变
const changePage = async () => {
  await getDetails()
}

// ==========
// 测试报告
// ==========

// 获取 - 测试报告
const dataReport = ref<TestReportDetailsReport>({})

// 计算 - 成功
const success = computed(() => {
  let t = (dataReport.value.success / dataReport.value.total) * 100
  return t.toFixed(2) + "%"
})

// 计算 - 失败
const fail = computed(() => {
  let t = (dataReport.value.fail / dataReport.value.total) * 100
  return t.toFixed(2) + "%"
})

// ==========
// 详细信息
// ==========

// 是否加载成功
const isLoading = ref<boolean>(false)

// 定时器
const isLoadingID = ref<number>()

// 数据 - 测试报告详细数据
const dataDetails = reactive<TestReportDetailsArray[]>([])
watchEffect(() => {
  // 监听运行状态并实时更新
  if (isLoading.value && dataReport.value.status === 1) {
    isLoadingID.value = window.setInterval(() => {
      getDetails()
    }, 1500)
  } else {
    window.clearInterval(isLoadingID.value)
  }
})

// 获取 - 测试报告详细信息
const getDetails = async () => {
  dataDetails.length = 0
  isLoading.value = false
  await GetPanelTestReportDetails(id, pageInfo.size, pageInfo.current).then((data: TestReportDetails) => {
    pageInfo.total = data.count
    data.details?.forEach((v: TestReportDetailsArray) => {
      dataDetails.push(v)
    })
    dataReport.value = data.report
    isLoading.value = true
  }).catch((err: string) => {
    ElMessage({
      message: "获取测试报告详细信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
    isLoading.value = false
    return
  })
}
getDetails()
</script>

<template>
  <el-page-header content="测试报告详细信息" @back="goBackUpPage"/>
  <div class="result">
    <div class="result-bar">
      <div>
        <span>状态&nbsp;&nbsp;</span>
        <el-tag effect="plain" type="danger" v-if="dataReport.status === -2">失败</el-tag>
        <el-tag effect="plain" type="warning" v-else-if="dataReport.status === -1">已终止</el-tag>
        <el-tag effect="plain" v-else-if="dataReport.status === 1">运行中</el-tag>
        <el-tag effect="plain" type="success" v-else-if="dataReport.status === 2">已完成</el-tag>
      </div>
      <div>请求数&nbsp;&nbsp;<b>{{ dataReport.total }}</b></div>
      <div>通过&nbsp;&nbsp;{{ success }}&nbsp;&nbsp;<b>{{ dataReport.success }}</b></div>
      <div>失败&nbsp;&nbsp;{{ fail }}&nbsp;&nbsp;<b>{{ dataReport.fail }}</b></div>
      <div>总耗时&nbsp;&nbsp;<b>{{ convertTime(dataReport.speed_time) }}</b></div>
    </div>
    <div class="result-data">
      <el-table
          :data="dataDetails"
          :show-header="false"
          empty-text="没有数据"
          scrollbar-always-on
          style="height: calc(100vh - 15px - 24px - 20px - 60px - 10px - 15px - 24px - 15px);"
          v-loading="!isLoading"
      >
        <el-table-column width="60">
          <template #default="scope">
            <el-tag effect="dark" type="success" v-if="scope.row.status === 1">通过</el-tag>
            <el-tag effect="dark" type="danger" v-else>失败</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="method" width="75"/>
        <el-table-column prop="url" :show-overflow-tooltip="true"/>
        <el-table-column prop="name" :show-overflow-tooltip="true"/>
        <el-table-column width="230">
          <template #default="scope">
            <span>状态码:&nbsp;&nbsp;{{ scope.row.status_code }}&nbsp;&nbsp;&nbsp;&nbsp;</span>
            <span>耗时:&nbsp;&nbsp;{{ convertTime(scope.row.speed_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column width="120">
          <template #default="scope">
            <el-button text type="primary" @click="goTestDetailPage(scope.row.uid)">更多详情</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="result-page">
      <el-pagination
          small
          background
          hide-on-single-page
          layout="total, prev, pager, next, jumper"
          :total="pageInfo.total"
          v-model:current-page="pageInfo.current"
          v-model:page-size="pageInfo.size"
          @current-change="changePage"
      />
    </div>
  </div>
</template>

<style scoped lang="less">
.result {
  width: 100%;
  margin-top: 20px;

  &-bar {
    height: 60px;
    display: flex;
    justify-content: space-around;
    align-items: center;
    flex-wrap: nowrap;
    background-color: #ffffff;
    box-shadow: 0 0 6px rgba(0, 0, 0, 0.2);
    border-radius: 6px;
    overflow: hidden;
    margin-bottom: 10px;
  }

  &-page {
    width: 100%;
    display: flex;
    justify-content: flex-end;
    padding-top: 15px;
  }
}
</style>