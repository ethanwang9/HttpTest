<script setup lang="ts">
import {reactive} from "vue";
import Uuid from "@/utils/uuid.ts";
import {isValidURL} from "@/utils/validator.ts";
import {ProxyList, SystemUpdate, SystemUpdateVersion} from "@/apis/types.ts";
import {GetPanelSystemProxyList, SetPanelSystemProxyList, SetPanelSystemUpdate} from "@/apis/panel_system.ts";

// ==========
// 代理池
// ==========

// 标签下表
const tabActive = ref<string>("proxy")

// 标签切换
const tabChange = (name: string) => {
  switch (name) {
    case "proxy":
      getProxyList()
      break
    case "version":
      break
  }
}

// 代理池列表
const proxy_list = reactive<ProxyList[]>([])

// 添加代理
const addProxyValue = ref<string>("")
const showAddProxy = ref<boolean>(false)
const addProxy = () => {
  // 显示弹窗
  showAddProxy.value = true
  addProxyValue.value = ""
}
const addProxyFun = async () => {
  if (!isValidURL(addProxyValue.value)) {
    ElMessage({
      message: "不是一个有效的 HTTP/HTTPS 代理池链接",
      type: 'warning',
      showClose: true,
    })
    return
  }
  proxy_list.push({
    "uid": Uuid(),
    "url": addProxyValue.value,
  })
  showAddProxy.value = false
  await setProxyList()
}

// 删除代理
const delProxyFun = async (i: number) => {
  proxy_list.splice(i, 1)
  await setProxyList()
}

// 提交代理列表
const setProxyList = async () => {
  await SetPanelSystemProxyList(proxy_list).catch((err: string) => {
    ElMessage({
      message: "设置系统代理池失败 Err: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}

// 获取代理列表
const getProxyList = async () => {
  proxy_list.length = 0
  await GetPanelSystemProxyList().then((data: ProxyList[] | null) => {
    if (data !== null && data.length != 0) {
      for (let i = 0; i < data.length; i++) {
        proxy_list[i] = data[i]
      }
    }
  }).catch((err: string) => {
    ElMessage({
      message: "获取系统代理池失败 Err: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}
getProxyList()

// ==========
// 更新
// ==========

// 数据
const dataUpdate = reactive<SystemUpdate>({})

// 版本信息
const dataUpdateVersion = reactive<SystemUpdateVersion>({})

// 更新加载
const updateLoading = ref<boolean>(false)

// 是否显示错误
const isShowUpdateError = computed(() => {
  return dataUpdate.msg != undefined && dataUpdate.msg.length !== 0 && dataUpdate.status >= 1 && dataUpdate.status <= 5
})

// 获取更新信息
const getUpdate = async () => {
  updateLoading.value = false
  await SetPanelSystemUpdate("1.0.0", "1.0.0").then((data: SystemUpdate) => {
    for (let key in data) {
      if (key === "version") {
        for (let key2 in data[key]) {
          dataUpdateVersion[key2] = data[key][key2]
        }
      } else {
        dataUpdate[key] = data[key]
      }
    }
    updateLoading.value = true
  }).catch((err: string) => {
    ElMessage({
      message: "获取系统更新信息失败 Err: " + err,
      type: 'warning',
      showClose: true,
    })
    updateLoading.value = true
  })
}
getUpdate()
</script>

<template>
  <el-row class="config">
    <el-col :span="16">
      <el-card shadow="hover" style="height: 500px;">
        <el-tabs v-model="tabActive" tab-position="left" @tab-change="tabChange">
          <el-tab-pane label="代理池配置" name="proxy" style="overflow-y: auto;height: 460px;">
            <el-alert
                effect="dark"
                title="代理池配置要求"
                type="info"
                description="请使用 HTTP/HTTPS 类型的代理池链接，返回内容请选择 text 类型并且一行一个代理IP"
                show-icon
                style="margin-bottom: 10px;"
            />
            <el-table
                empty-text="暂无代理池"
                :show-header="false"
                :data="proxy_list"
                :show-overflow-tooltip="true"
            >
              <el-table-column width="70">
                <template #default="scope">
                  <span>#{{ scope.$index + 1 }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="url"></el-table-column>
              <el-table-column width="90">
                <template #default="scope">
                  <el-button text type="danger" @click.stop="delProxyFun(scope.$index)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
            <div class="proxy-btn">
              <el-button text type="primary" class="proxy-btn-add" @click="addProxy">
                <span>添加代理池</span>
                <template #icon>
                  <el-icon>
                    <i-ep-Plus/>
                  </el-icon>
                </template>
              </el-button>
            </div>
          </el-tab-pane>
          <el-tab-pane label="版本更新" name="version" style="overflow-y: auto;height: 460px;" v-loading="!updateLoading">
            <el-alert
                effect="dark"
                title="系统警告"
                type="error"
                :description="dataUpdate.msg"
                show-icon
                style="margin-bottom: 10px;"
                v-if="isShowUpdateError"
            />
            <div class="version">
              <div class="version-left">
                <el-icon :size="60">
                  <i-ep-Promotion/>
                </el-icon>
              </div>
              <div class="version-right">
                <h4>最新版本号 {{ dataUpdateVersion.release }}</h4>
                <p v-if="dataUpdate.status >= 6 && dataUpdate.status <= 8">{{ dataUpdate.msg }}</p>
                <p v-else>当前为最新版本，无需更新</p>
              </div>
              <div class="version-bottom">
                <span>客户端版本号: {{ dataUpdateVersion.client }}</span>
                <span>服务器版本号: {{ dataUpdateVersion.server }}</span>
              </div>
            </div>
            <h4>更新日志</h4>
            <div v-html="dataUpdate.log" class="log"></div>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-col>
  </el-row>
  <!--  添加代理弹窗-->
  <el-dialog
      v-model="showAddProxy"
      title="添加代理"
      width="500"
  >
    <el-input v-model="addProxyValue" placeholder="请输入代理池链接"></el-input>
    <template #footer>
      <el-button @click="showAddProxy = !showAddProxy">取消</el-button>
      <el-button type="primary" @click="addProxyFun">添加</el-button>
    </template>
  </el-dialog>
</template>

<style scoped lang="less">
.config {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  overflow: hidden;
  background-color: #F1F3FE;

  .proxy-btn {
    display: flex;
    flex-direction: column;
    align-items: center;

    &-add {
      margin-top: 10px;
    }
  }

  .version {
    display: flex;
    flex-wrap: wrap;
    background-image: linear-gradient(to right, #DA22FF 0%, #9733EE 51%, #DA22FF 100%);
    background-size: 200% auto;
    text-transform: uppercase;
    transition: 0.5s;
    color: #FFFFFF;
    padding: 15px;
    box-sizing: border-box;
    border-radius: 15px;

    &:hover {
      background-position: right center;
      color: #fff;
      text-decoration: none;
    }

    &-left {
      width: 80px;
      height: 80px;
      display: flex;
      justify-content: center;
      align-items: center;
    }

    &-right {
      padding-left: 15px;

      h4 {
        font-size: 18px;
        font-weight: bold;
        padding: 10px 0;
        margin: 0;
      }

      p {
        font-size: 14px;
        padding: 5px 0;
        margin: 0;
      }
    }

    &-bottom {
      width: 100%;
      text-align: right;
      font-size: 12px;

      & > *:not(:last-child) {
        padding-right: 1em;
      }
    }
  }
}

.log {
  line-height: 1.6em;
}
</style>