<script setup lang="ts">
import MonacoEditor from "@/components/MonacoEditor.vue";
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {
  DelPanelInterfaceInfo,
  GetPanelInterfaceDir,
  GetPanelInterfaceInfo,
  SetPanelInterfaceDir, SetPanelInterfaceInfo, UpdatePanelInterfaceInfo
} from "@/apis/panel_interface.ts";
import {InterfaceData, MenuTree, MonacoParams, PageInfo} from "@/apis/types.ts";
import Uuid from "@/utils/uuid.ts";
import JsonFormat from "@/utils/json-format.ts";
import {isValidURL} from "@/utils/validator.ts";
import {convertTime} from "@/utils/tools.ts";

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
  await GetPanelInterfaceDir().then((data: MenuTree[] | null) => {
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
  await SetPanelInterfaceDir(JSON.stringify(treeData.value)).catch((err: string) => {
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
  if (e.type === "interface") {
    treeDataActive.value = e.data
    GetInterfaceInfo()
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

// 新建接口
const showCreateInterface = ref<boolean>(false)
const interfaceName = ref<string>("")
const createInterface = async () => {
  let u = Uuid()

  // 新建接口数据
  await SetPanelInterfaceInfo(u, interfaceName.value.trim()).catch((err: string) => {
    ElMessage({
      message: "创建接口失败, Error: " + err,
      type: 'warning',
      showClose: true,
    })
    return
  })

  // 新建接口
  treeData.value.unshift({
    "label": interfaceName.value.trim(),
    "type": "interface",
    "data": u,
  })

  // 上传菜单树数据
  await setMenuTree()
  // 关闭对话框
  interfaceName.value = ""
  showCreateInterface.value = !showCreateInterface.value
}

// 删除接口
const delInterface = async (uid: string) => {
  await DelPanelInterfaceInfo(uid).then(() => {
    ElMessage({
      message: "删除接口成功",
      type: 'success',
      showClose: true,
    })
  }).catch((err: string) => {
    ElMessage({
      message: "删除接口失败, Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}

// 删除接口或目录
const delMenuTree = (data: MenuTree, temp: MenuTree[] = [], mode: string = "") => {
  if (mode.length === 0) {
    mode = data.type
    temp = treeData.value
  }

  switch (mode) {
    case "interface":
      for (let i = 0; i < temp.length; i++) {
        if (temp[i].type === "interface" && temp[i].$treeNodeId === data.$treeNodeId) {
          // 删除接口
          delInterface(temp[i].data)
          temp.splice(i, 1)
          setMenuTree()
          break
        } else if (temp[i].type === "dir" && temp[i].children?.length > 0) {
          for (let j = 0; j < temp[i].children?.length; j++) {
            delMenuTree(data, temp[i].children, "interface")
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

// ==========
// 接口信息
// ==========

// 获取接口请求状态
const apiStatus = ref<boolean>(false)

// http 请求数据
const httpData = reactive<InterfaceData>({})
const httpDataUrlShowTips = ref<boolean>(false)
watchEffect(() => {
  // 监听 URL 错误
  httpDataUrlShowTips.value = !isValidURL(httpData.url);
})


// 递归修改菜单接口名称
const changeMenuTreeName = (uid: string, name: string, arr: MenuTree[] = []) => {
  if (arr.length === 0) {
    arr = treeData.value
  }
  for (let i = 0; i < arr.length; i++) {
    if (arr[i].type === "interface" && arr[i].data === uid) {
      arr[i].label = name
      break
    } else if (arr[i].children?.length > 0 && arr[i].type === "dir") {
      changeMenuTreeName(uid, name, arr[i].children)
    }
  }
}

// test 选项数据
const httpDataTest = reactive<InterfaceData["test"]>({})
watchEffect(() => {
  for (let key in httpDataTest) {
    if (apiStatus.value && httpDataTest[key] !== httpData.test[key]) {
      httpData.test[key] = httpDataTest[key]
    }
  }
})

// http 请求类型
const methodList = [
  {
    value: 'get',
    label: 'GET',
  },
  {
    value: 'post',
    label: 'POST',
  },
  {
    value: 'put',
    label: 'PUT',
  },
  {
    value: 'delete',
    label: 'DELETE',
  },
  {
    value: 'options',
    label: 'OPTIONS',
  },
  {
    value: 'head',
    label: 'HEAD',
  },
  {
    value: 'patch',
    label: 'PATCH',
  }
]

// 编辑器数据
const monacoShow = ref<boolean>(false)
const monacoDataShowTips = ref<boolean>(false)
const monacoData = ref<string>("")
watchEffect(() => {
  if (monacoData.value === "") {
    monacoDataShowTips.value = false
    return false
  }
  if (apiStatus.value && monacoData.value.length !== 0) {
    try {
      // 判断数据结构是否正确
      let v: MonacoParams = JSON.parse(monacoData.value)
      if (typeof v !== "object") {
        monacoDataShowTips.value = true
        return false
      }
      for (let key in v) {
        if (key !== "header" && key !== "param" && key !== "body") {
          monacoDataShowTips.value = true
          return false
        }
      }
      // 写入数据
      httpData.header = v
      monacoDataShowTips.value = false
    } catch (e) {
      monacoDataShowTips.value = true
      return false
    }
  }
})

// 获取接口信息
const GetInterfaceInfo = async () => {
  apiStatus.value = false
  await GetPanelInterfaceInfo(treeDataActive.value).then((data: InterfaceData) => {
    for (let key in data) {
      httpData[key] = data[key]
    }
    for (let key in data.test) {
      httpDataTest[key] = data.test[key]
    }
    if (monacoShow.value === false) {
      monacoData.value = "!@#@" + JsonFormat(JSON.stringify(data.header))
      monacoShow.value = true
    } else {
      monacoData.value = "!@#@" + JSON.stringify(data.header)
    }
    apiStatus.value = true
  }).catch((err: string) => {
    ElMessage({
      message: "获取接口信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
    apiStatus.value = true
  })
}

// 保存接口信息
const saveInterfaceInfo = async () => {
  // 判断状态
  if (monacoDataShowTips.value) {
    ElMessage({
      message: "【请求参数】配置错误，请填写正确后再次重试！",
      type: 'warning',
      showClose: true,
    })
    return
  }
  if (httpDataUrlShowTips.value) {
    ElMessage({
      message: "【请求地址】配置错误，请填写正确后再次重试！",
      type: 'warning',
      showClose: true,
    })
    return
  }
  if (httpData.url.trim().length === 0) {
    ElMessage({
      message: "【请求地址】不能为空，请填写正确后再次重试！",
      type: 'warning',
      showClose: true,
    })
    return
  }
  if (httpData.method.trim().length === 0) {
    ElMessage({
      message: "【请求类型】不能为空，请填写正确后再次重试！",
      type: 'warning',
      showClose: true,
    })
    return
  }

  // 处理数据
  let t: InterfaceData = {}
  for (let key in httpData) {
    if (key !== "data") {
      t[key] = httpData[key]
    }
  }

  // 保存接口名
  changeMenuTreeName(treeDataActive.value, t.name)
  await setMenuTree()

  // 发送请求
  await UpdatePanelInterfaceInfo(t).then(() => {
    ElMessage({
      message: "保存成功",
      type: 'success',
      showClose: true,
    })
  }).catch((err: string) => {
    ElMessage({
      message: "保存接口信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}

// ==========
// 其他
// ==========
// TODO 发送测试
const sendTest = () => {
  ElMessage({
    message: "TODO",
    type: 'warning',
    showClose: true,
  })
}

// TODO 停止测试
const stopTest = () => {
  ElMessage({
    message: "TODO",
    type: 'warning',
    showClose: true,
  })
}
</script>

<template>
  <div class="main">
    <div class="left">
      <div class="left-toolbar">
        <el-button type="danger" @click="showCreateInterface=!showCreateInterface">
          <span>新建接口</span>
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
                      <!--<el-dropdown-item v-if="data.type==='dir'">新建接口</el-dropdown-item>-->
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
        <el-empty description="请选择接口进行测试"/>
      </div>
    </template>
    <div class="right" v-else>
      <div class="right-base">
        <div class="right-base-method">
          <el-select
              clearable
              v-model="httpData.method"
              placeholder="请求类型"
          >
            <el-option
                v-for="item in methodList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </div>
        <div class="right-base-url">
          <el-input v-model="httpData.url" placeholder="链接地址" clearable/>
        </div>
        <div class="right-base-btns">
          <el-button type="primary" @click="sendTest">发送</el-button>
          <el-button type="danger" @click="stopTest">停止</el-button>
          <el-button type="success" @click="saveInterfaceInfo">保存</el-button>
        </div>
        <div class="right-base-tips" v-show="httpDataUrlShowTips">
          <span>请求地址格式错误</span>
        </div>
      </div>
      <div class="right-name">
        <div class="right-title">接口名称</div>
        <el-input v-model="httpData.name" placeholder="接口名称"/>
      </div>
      <div class="right-config">
        <div class="right-config-request">
          <div class="right-title">
            <span>请求参数</span>
            <span
                style="background-color:red;color: #FFFFFF;padding: 2px 8px;border-radius: 4px;"
                v-show="monacoDataShowTips"
            >请求参数格式或内容错误</span>
          </div>
          <MonacoEditor v-model="monacoData" :readonly="false" style="height: calc(350px - 35px);"></MonacoEditor>
        </div>
        <div class="right-config-test">
          <div class="right-title">测试参数</div>
          <p>并发数</p>
          <el-slider :min="1" :max="10000" v-model="httpDataTest.num" show-input/>
          <p>运行时间(分钟)</p>
          <el-slider :min="1" :max="360" v-model="httpDataTest.time" show-input/>
          <p>间隔停顿（毫秒）</p>
          <el-slider :min="0" :max="60000" v-model="httpDataTest.sleep" show-input/>
        </div>
      </div>
      <div class="right-result">
        <div class="right-title">响应数据</div>
        <div class="right-result-bar">
          <div>
            <span>状态&nbsp;&nbsp;</span>
            <el-tag effect="plain" type="danger" v-if="2 === -2">失败</el-tag>
            <el-tag effect="plain" type="warning" v-else-if="2 === -1">已终止</el-tag>
            <el-tag effect="plain" v-else-if="2 === 1">运行中</el-tag>
            <el-tag effect="plain" type="success" v-else-if="2 === 2">已完成</el-tag>
          </div>
          <div>请求数&nbsp;&nbsp;<b>1000</b></div>
          <div>通过&nbsp;&nbsp;1&nbsp;&nbsp;<b>12</b></div>
          <div>失败&nbsp;&nbsp;2&nbsp;&nbsp;<b>12</b></div>
          <div>总耗时&nbsp;&nbsp;<b>100ms</b></div>
        </div>
        <el-button type="primary" class="right-result-btn">查看详情</el-button>
      </div>
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
    <el-dialog v-model="showCreateInterface" title="新建接口" width="380" draggable>
      <div style="padding-bottom: 20px;">
        <el-input placeholder="接口名称" v-model="interfaceName"/>
      </div>
      <div style="display: flex;justify-content: flex-end;">
        <el-button @click="showCreateInterface = !showCreateInterface">取消</el-button>
        <el-button type="primary" @click="createInterface">确定</el-button>
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
    overflow-y: auto;

    &-title {
      height: 16px;
      font-size: 14px;
      padding: 5px 0 10px 0;
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    &-base {
      width: 100%;
      height: 40px;
      display: flex;
      justify-content: center;
      align-items: center;
      flex-wrap: wrap;
      margin-bottom: 10px;

      &-method {
        width: 120px;
      }

      &-url {
        width: calc(100% - 120px - 220px);
      }

      &-btns {
        width: 220px;
        display: flex;
        justify-content: center;
        align-items: center;
      }

      &-tips {
        width: 100%;
        margin: 8px 0 0 120px;

        span {
          background-color: red;
          color: #FFFFFF;
          padding: 2px 8px;
          border-radius: 4px;
          font-size: 14px;
        }
      }
    }

    &-name {
      width: 100%;
      margin-bottom: 10px;
    }

    &-config {
      display: flex;
      justify-content: space-between;
      width: 100%;
      height: 350px;
      margin-bottom: 10px;

      &-request {
        width: 65%;
        padding-right: 30px;
        box-sizing: border-box;
      }

      &-test {
        width: 35%;
      }
    }

    &-result {
      width: 100%;

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

      &-btn {
        width: 100%;
        margin: 10px 0 15px 0;
      }

      &-page {
        width: 100%;
        display: flex;
        justify-content: flex-end;
        padding-top: 15px;
      }
    }
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