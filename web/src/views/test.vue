<script setup lang="ts">
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {InterfaceData, MenuTree, TestData, TestDataList, TestReport} from "@/apis/types.ts";
import Uuid from "@/utils/uuid.ts";
import {
  DelPanelTestInfo, DelPanelTestReport,
  GetPanelTestDir,
  GetPanelTestInfo, GetPanelTestReport,
  SetPanelTestDir,
  SetPanelTestInfo, UpdatePanelTestInfo
} from "@/apis/panel_test.ts";
import {TabPaneName} from "element-plus";
import {GetPanelInterfaceDir, GetPanelInterfaceInfo} from "@/apis/panel_interface.ts";
import {convertTime} from "@/utils/tools.ts"

// ==========
// 路由
// ==========
const router = useRouter()

// 监听当前路由是否为子路由
const isChildNode = computed(() => {
  let t = router.currentRoute.value.path.substring(1).split("/")
  return t.length > 1;
})

// ==========
// 文件树
// ==========

// 文件树加载
const treeLoading = ref<boolean>(true)

// 文件树数据
const treeData = ref<MenuTree[]>([])

// 文件树选中ID
const treeDataActive = ref("")

// req: 获取菜单树
const getMenuTree = async () => {
  treeLoading.value = true
  await GetPanelTestDir().then((data: MenuTree[] | null) => {
    treeLoading.value = false
    if (data !== null && data.length !== 0) {
      treeData.value = data
    }
  }).catch((err: string) => {
    treeLoading.value = false
    ElMessage({
      message: err,
      type: 'warning',
      showClose: true,
    })
  })
}
getMenuTree()

// 上传菜单树
const setMenuTree = async () => {
  await SetPanelTestDir(JSON.stringify(treeData.value)).catch((err: string) => {
    ElMessage({
      message: err,
      type: 'warning',
      showClose: true,
    })
  })
}

// 移动文件树
const moveTree = () => {
  setMenuTree()
}

// 点击文件树
const clickTree = (e: MenuTree) => {
  if (e.type === "scene") {
    treeDataActive.value = e.data
    getPanelTestInfo()
  }
}

// 文件树禁止规则
const stopTree = (n1: Node, n2: Node) => {
  // 1. 禁止dir移入interface中
  return n2.data.type !== "interface";
}

// 新建文件夹
const showCreateDir = ref<boolean>(false)
const dirName = ref<string>("")
const createDir = () => {
  // 新建目录
  treeData.value.unshift({
    "label": dirName.value.trim(),
    "type": "dir",
  })
  // 上传数据
  setMenuTree()
  // 关闭对话框
  dirName.value = ""
  showCreateDir.value = !showCreateDir.value
}

// 新建场景
const showCreateScene = ref<boolean>(false)
const sceneName = ref<string>("")
const createScene = async () => {
  let u = Uuid()

  // 新建场景数据
  await SetPanelTestInfo(u, sceneName.value.trim()).catch((err: string) => {
    ElMessage({
      message: "创建场景失败, Error: " + err,
      type: 'warning',
      showClose: true,
    })
    return
  })

  // 新建场景
  treeData.value.unshift({
    "label": sceneName.value.trim(),
    "type": "scene",
    "data": u,
  })

  // 上传菜单树数据
  await setMenuTree()
  // 关闭对话框
  sceneName.value = ""
  showCreateScene.value = !showCreateScene.value
}

// 删除场景
const delScene = async (uid: string) => {
  await DelPanelTestInfo(uid).then(() => {
    ElMessage({
      message: "删除场景成功",
      type: 'success',
      showClose: true,
    })
  }).catch((err: string) => {
    ElMessage({
      message: "删除场景失败, Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}

// 删除场景或目录
const delMenuTree = (data: MenuTree, temp: MenuTree[] = [], mode: string = "") => {
  if (mode.length === 0) {
    mode = data.type
    temp = treeData.value
  }

  switch (mode) {
    case "scene":
      for (let i = 0; i < temp.length; i++) {
        if (temp[i].type === "scene" && temp[i].$treeNodeId === data.$treeNodeId) {
          // 删除场景
          delScene(temp[i].data)
          temp.splice(i, 1)
          setMenuTree()
          break
        } else if (temp[i].type === "dir" && temp[i].children?.length > 0) {
          for (let j = 0; j < temp[i].children?.length; j++) {
            delMenuTree(data, temp[i].children, "scene")
          }
        }
      }
      break
    case "dir":
      for (let i = 0; i < temp.length; i++) {
        if (temp[i].$treeNodeId === data.$treeNodeId) {
          temp.splice(i, 1)
          setMenuTree()
          break
        }
      }
      break
  }
}

// 递归修改菜单接口名称
const changeMenuTreeName = (uid: string, name: string, arr: MenuTree[] = []) => {
  if (arr.length === 0) {
    arr = treeData.value
  }
  for (let i = 0; i < arr.length; i++) {
    if (arr[i].type === "scene" && arr[i].data === uid) {
      arr[i].label = name
      break
    } else if (arr[i].children?.length > 0 && arr[i].type === "dir") {
      changeMenuTreeName(uid, name, arr[i].children)
    }
  }
}

// ==========
// 步骤条
// ==========

// 步骤条下标
const stepActive = ref<string>("steps")
// 步骤条点击事件
const stepChange = (name: TabPaneName) => {
  switch (name) {
    case "steps":
      // 更新测试步骤信息
      getPanelTestInfo()
      break
    case "report":
      // 更新测试报告信息
      getReportData()
      break
  }
}

// ==========
// 测试步骤
// ==========

// 获取接口请求状态
const apiStatus = ref<boolean>(false)

// 场景数据
const sceneData = reactive<TestData>({})

// 获取场景信息
const getPanelTestInfo = async () => {
  testList.length = 0
  apiStatus.value = false
  await GetPanelTestInfo(treeDataActive.value).then(async (data: TestData) => {
    for (let key in data) {
      if (key === "list" && data.list === null) {
        sceneData[key] = []
      } else if (key === "test") {
        httpDataTest.sleep = data.test.sleep
        httpDataTest.num = data.test.num
        httpDataTest.time = data.test.time
        sceneData[key] = data[key]
      } else {
        sceneData[key] = data[key]
      }
    }
    // 写入 testList
    await stepsInterfaceUidList(data.list)
    // 写入接口状态
    apiStatus.value = true
  }).catch((err: string) => {
    ElMessage({
      message: "获取场景信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
    apiStatus.value = false
  })
}

// test 选项数据
const httpDataTest = reactive<TestData["test"]>({})
watchEffect(() => {
  for (let key in httpDataTest) {
    if (apiStatus.value && httpDataTest[key] !== sceneData.test[key]) {
      sceneData.test[key] = httpDataTest[key]
    }
  }
})

// 测试接口数据集合
const testList = reactive<TestDataList[]>([])

// 是否显示测试步骤
const isShowSteps = ref<Boolean>(false)

// 接口文件树
const interfaceMenuTree = reactive<MenuTree[]>([])

// 打开添加接口弹窗
const openSteps = async () => {
  // 获取接口菜单树
  await GetPanelInterfaceDir().then((data: MenuTree[]) => {
    interfaceMenuTree.length = 0
    for (let i = 0; i < data.length; i++) {
      interfaceMenuTree.push(data[i])
    }
  }).catch((err: string) => {
    isShowSteps.value = false
    ElMessage({
      message: "获取接口菜单信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
  // 打开弹窗
  isShowSteps.value = true
}

// 接口菜单元素
const interfaceMenuTreeRef = ref()

// 添加测试步骤 - 接口
const addStepsInterface = async () => {
  let node: MenuTree[] = interfaceMenuTreeRef.value.getCheckedNodes()
  let list = []
  let l = sceneData.list
  node.forEach((v: MenuTree) => {
    if (v.type === "interface") {
      list.push(v.data)
      // 写入 sceneData.list
      l?.push(v.data)
    }
  })
  // 查询接口信息
  await stepsInterfaceUidList(list)
  // 关闭弹窗
  isShowSteps.value = false
}

// 通过uid列表查询接口详细信息并渲染到页面中
const stepsInterfaceUidList = async (list: Array<string> | null) => {
  if (list === null || list.length === 0) {
    return
  }
  for (const v: string of list) {
    await GetPanelInterfaceInfo(v).then((data: InterfaceData) => {
      // 写入 testList
      testList.push({
        uid: data.uid,
        name: data.name,
        url: data.url,
        method: data.method,
      })
    }).catch((err: string) => {
      ElMessage({
        message: "获取接口详细信息失败 Error: " + err,
        type: 'warning',
        showClose: true,
      })
    })
  }
}

// 删除接口列表
const delInterfaceList = (uid: string) => {
  // 删除数据
  let l = sceneData.list
  let i = l?.indexOf(uid)
  l?.splice(i, 1)
  // 删除页面数据
  for (let j = 0; j < testList.length; j++) {
    if (testList[j].uid === uid) {
      testList.splice(j, 1)
      break
    }
  }
}

// 更新场景
const updateScene = async () => {
  // 判断状态
  if (!apiStatus.value) {
    ElMessage({
      message: "接口状态异常，请刷新页面后重试",
      type: 'warning',
      showClose: true,
    })
    return
  }
  if (sceneData.name.trim().length === 0) {
    ElMessage({
      message: "【场景名称】配置错误，请填写正确后再次重试！",
      type: 'warning',
      showClose: true,
    })
    return
  }

  // 保存场景名称
  changeMenuTreeName(sceneData.uid, sceneData.name)
  await setMenuTree()

  // 更新场景信息
  // 发送请求
  await UpdatePanelTestInfo(sceneData).then(() => {
    ElMessage({
      message: "保存成功",
      type: 'success',
      showClose: true,
    })
  }).catch((err: string) => {
    ElMessage({
      message: "保存场景信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}

// TODO 运行测试
const runTest = () => {
  ElMessage({
    message: "TODO",
    type: 'warning',
    showClose: true,
  })
}

// ==========
// 测试报告
// ==========

// 测试报告数据
const reportData = reactive<TestReport[]>([])

// 获取测试报告数据
const getReportData = async () => {
  apiStatus.value = false
  reportData.length = 0
  await GetPanelTestReport(sceneData.uid).then((data: TestReport[]) => {
    data.forEach((v: TestReport) => {
      reportData.push(v)
    })
    apiStatus.value = true
  }).catch((err: string) => {
    ElMessage({
      message: "获取测试报告数据失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
    apiStatus.value = false
  })
}

// 删除报告
const delReport = async (uid: string) => {
  // 发送请求
  await DelPanelTestReport(uid).then(() => {
    ElMessage({
      message: "删除成功",
      type: 'success',
      showClose: true,
    })
  }).catch((err: string) => {
    ElMessage({
      message: "删除测试报告失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })

  // 移除缓存
  for (let i = 0; i < reportData.length; i++) {
    if (reportData[i].uid === uid) {
      reportData.splice(i, 1)
      break
    }
  }
}

// 查看测试报告详情信息
const testReportDetails = (row: TestReport) => {
  router.push({
    name: "Report",
    query: {
      id: row.uid,
    },
  })
}
</script>

<template>
  <div class="main">
    <div class="left">
      <div class="left-toolbar">
        <el-button type="danger" @click="showCreateScene=!showCreateScene">
          <span>新建场景</span>
          <template #icon>
            <i-ep-Plus/>
          </template>
        </el-button>
        <el-button type="primary" @click="showCreateDir=!showCreateDir">
          <span>新建目录</span>
          <template #icon>
            <i-ep-Folder/>
          </template>
        </el-button>
      </div>
      <div class="left-dir" v-loading="treeLoading">
        <el-tree
            empty-text="没有数据"
            :data="treeData"
            :allow-drop="stopTree"
            accordion draggable
            @node-drop="moveTree"
            @node-click="clickTree"
        >
          <template #default="{ node, data }">
            <div class="left-dir-box">
              <div class="left-dir-box-content">
                <el-icon size="16px">
                  <i-ep-Folder v-if="data.type==='dir'"/>
                  <i-ep-Postcard v-else/>
                </el-icon>
                <span class="left-dir-box-content-text">{{ data.label }}</span>
              </div>
              <div class="left-dir-box-toolbar">
                <el-dropdown>
                  <el-icon size="12px" class="el-dropdown-link">
                    <i-ep-MoreFilled/>
                  </el-icon>
                  <template #dropdown>
                    <el-dropdown-menu>
                      <!--<el-dropdown-item v-if="data.type==='dir'">新建场景</el-dropdown-item>-->
                      <el-dropdown-item @click="delMenuTree(data)">删除</el-dropdown-item>
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </div>
            </div>
          </template>
        </el-tree>
      </div>
    </div>
    <template v-if="isChildNode">
      <div class="right">
        <router-view></router-view>
      </div>
    </template>
    <template v-else-if="treeDataActive.length == 0">
      <div class="right">
        <el-empty description="请选择场景进行测试"/>
      </div>
    </template>
    <div class="right" v-else>
      <el-tabs v-model="stepActive" @tab-change="stepChange">
        <el-tab-pane label="测试步骤" name="steps">
          <div class="right-steps">
            <div class="right-steps-list">
              <el-input v-model="sceneData.name" placeholder="场景名称"/>
              <el-table
                  empty-text="请添加步骤"
                  :data="testList"
                  :show-overflow-tooltip="true"
                  :show-header="false"
                  style="margin-top: 15px;"
              >
                <el-table-column width="90">
                  <template #default="scope">#{{ scope.$index + 1 }}</template>
                </el-table-column>
                <el-table-column prop="method" width="110">
                  <template #default="scope">
                    <b>{{ scope.row.method.toUpperCase() }}</b>
                  </template>
                </el-table-column>
                <el-table-column prop="name"></el-table-column>
                <el-table-column prop="url"></el-table-column>
                <el-table-column width="85">
                  <template #default="scope">
                    <el-button text type="warning" @click="delInterfaceList(scope.row.uid)">删除</el-button>
                  </template>
                </el-table-column>
              </el-table>
              <el-button text type="primary" style="width: 140px;margin-top: 15px;" @click="openSteps">
                <span>添加步骤</span>
                <template #icon>
                  <i-ep-Plus/>
                </template>
              </el-button>
            </div>
            <div class="right-steps-test">
              <p>并发数</p>
              <el-slider :min="1" :max="10000" v-model="httpDataTest.num" show-input/>
              <p>运行时间(分钟)</p>
              <el-slider :min="1" :max="360" v-model="httpDataTest.time" show-input/>
              <p>间隔停顿（毫秒）</p>
              <el-slider :min="0" :max="60000" v-model="httpDataTest.sleep" show-input/>
              <div class="right-steps-test-btn">
                <el-button
                    @click="runTest"
                    type="primary"
                    style="width: 45%;"
                >
                  <span>运行</span>
                  <template #icon>
                    <el-icon :size="18">
                      <i-ep-CaretRight/>
                    </el-icon>
                  </template>
                </el-button>
                <el-button style="width: 45%;" @click="updateScene">保存</el-button>
              </div>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="测试报告" name="report">
          <el-table
              :data="reportData"
              empty-text="暂无测试报告"
              @row-click="testReportDetails"
              :show-overflow-tooltip="true"
          >
            <el-table-column label="报告信息">
              <template #default="scope">
                <span class="right-report-name">{{ sceneData.name }}</span><br>
                <span class="right-report-time">{{ scope.row.created_at }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="90">
              <template #default="scope">
                <div class="right-report-tag">
                  <el-tag effect="plain" type="danger" v-if="scope.row.status === -2">失败</el-tag>
                  <el-tag effect="plain" type="warning" v-else-if="scope.row.status === -1">已终止</el-tag>
                  <el-tag effect="plain" v-else-if="scope.row.status === 1">运行中</el-tag>
                  <el-tag effect="plain" type="success" v-else-if="scope.row.status === 2">已完成</el-tag>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="" label="结果" width="320">
              <template #default="scope">
                <span class="right-report-details-name">请求数</span>
                <el-tooltip placement="top" :content="scope.row.total">
                  <span class="right-report-details-num">{{ scope.row.total }}</span>
                </el-tooltip>
                <span class="right-report-details-name">成功</span>
                <el-tooltip placement="top" :content="scope.row.success">
                  <span class="right-report-details-num">{{ scope.row.success }}</span>
                </el-tooltip>
                <br>
                <span class="right-report-details-name">失败</span>
                <el-tooltip placement="top" :content="scope.row.fail">
                  <span class="right-report-details-num">{{ scope.row.fail }}</span>
                </el-tooltip>
                <span class="right-report-details-name">耗时</span>
                <el-tooltip placement="top" :content="convertTime(scope.row.speed_time)">
                  <span class="right-report-details-num">{{ convertTime(scope.row.speed_time) }}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column label="操作" width="100">
              <template #default="scope">
                <el-button text type="danger" @click.stop="delReport(scope.row.uid)">删除报告</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </div>
    <!--弹窗-->
    <el-dialog v-model="showCreateDir" title="新建目录" width="380" draggable>
      <div style="padding-bottom: 20px;">
        <el-input placeholder="目录名称" v-model="dirName"/>
      </div>
      <div style="display: flex;justify-content: flex-end;">
        <el-button @click="showCreateDir = !showCreateDir">取消</el-button>
        <el-button type="primary" @click="createDir">确定</el-button>
      </div>
    </el-dialog>
    <el-dialog v-model="showCreateScene" title="新建场景" width="380" draggable>
      <div style="padding-bottom: 20px;">
        <el-input placeholder="场景名称" v-model="sceneName"/>
      </div>
      <div style="display: flex;justify-content: flex-end;">
        <el-button @click="showCreateScene = !showCreateScene">取消</el-button>
        <el-button type="primary" @click="createScene">确定</el-button>
      </div>
    </el-dialog>
    <!--是否显示测试步骤-->
    <el-dialog class="show-steps" width="450" v-model="isShowSteps" title="添加接口" :draggable="true">
      <el-tree
          :data="interfaceMenuTree"
          show-checkbox
          ref="interfaceMenuTreeRef"
          empty-text="无接口数据"
      />
      <div class="show-steps-btn">
        <el-button @click="isShowSteps = !isShowSteps">取消</el-button>
        <el-button type="primary" @click="addStepsInterface">添加</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<style scoped lang="less">
.main {
  width: calc(100vw - 80px);
  height: 100vh;
  @left-width: 270px;

  .left {
    background-color: #FCFCFD;
    height: 100vh;
    width: @left-width;
    border-right: 2px solid #f1f2f6;
    padding: 20px 10px 5px 10px;
    box-sizing: border-box;
    overflow-y: auto;

    &-dir {
      &-box {
        width: 100%;
        display: flex;
        justify-content: space-between;
        align-items: center;

        &:hover {
          .left-dir-box-toolbar {
            display: block;
          }

          .left-dir-box-content > .left-dir-box-content-text {
            max-width: 80px;
          }
        }

        &-content {
          display: flex;
          align-items: center;

          &-text {
            padding-left: 4px;
            max-width: 130px;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            transition: all .15s ease-in;
          }
        }

        &-toolbar {
          display: none;
          padding-right: 0.5em;
          transition: all .15s ease-in;
        }
      }
    }

    &-toolbar {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: row;
      margin-bottom: 15px;
    }
  }

  .right {
    height: 100vh;
    width: calc(100vw - @left-width);
    background-color: #FFFFFF;
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    padding: 15px;
    overflow: hidden;

    &-steps {
      display: flex;
      flex-wrap: nowrap;
      gap: 15px;
      justify-content: space-between;
      box-sizing: border-box;
      @steps-test-width: 330px;

      &-list {
        width: calc(100vw - @left-width - 30px - @steps-test-width);
        height: calc(100vh - 30px - 40px - 15px);
        display: flex;
        align-items: center;
        flex-direction: column;
        overflow-y: auto;
      }

      &-test {
        width: @steps-test-width;
        height: 312px;
        padding: 0 15px 15px 20px;
        border: 2px solid #ececec;
        border-radius: 10px;
        box-sizing: border-box;
        background-color: #FFFFFF;

        &-btn {
          text-align: center;
          padding-top: 15px;
        }
      }
    }

    &-report {
      &-name {
        padding: 0;
        margin: 0;
        font-size: 16px;
      }

      &-time {
        padding: 4px 0 0 0;
        margin: 0;
        font-size: 12px;
        color: gray;
      }

      &-tag {
        text-align: center;
      }

      &-details {
        &-name {
          display: inline-block;
          width: 3em;
          padding-right: 0.5em;
        }

        &-num {
          display: inline-block;
          overflow: hidden;
          width: 6em;
          text-overflow: ellipsis;
          vertical-align: bottom;
        }

        &-num:not(:last-child) {
          padding-right: 1em;
        }
      }
    }
  }
}

.show-steps {
  &-btn {
    text-align: right;
    margin-top: 15px;
  }
}
</style>

<style lang="less">
.el-dropdown-link:focus {
  outline: none;
}

.el-dialog__body {
  padding: var(--el-dialog-padding-primary);
}
</style>