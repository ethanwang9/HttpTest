<script setup lang="ts">
import {RouteRecordRaw, useRouter} from "vue-router";
import {computed, ref} from "vue";
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import en from 'element-plus/dist/locale/en.mjs'
import {SoftwareInfo} from "@/apis/types.ts";
import {GetSystemBasic} from "@/apis/system.ts";

const router = useRouter()
const currentRoute = router.currentRoute

// 软件相关信息
const info = reactive<SoftwareInfo>({})
const getInfo = async function () {
  await GetSystemBasic().then((result: SoftwareInfo) => {
    for (let key in result) {
      info[key] = result[key]
    }
  }).catch((err: string) => {
    ElMessage({
      message: "获取面板基本信息失败 Error: " + err,
      type: 'warning',
      showClose: true,
    })
  })
}
getInfo()

// 跳转页面
const goPage = (url: string) => {
  window.open(url)
}

// 跳转 Router 地址
const goRouter = (name: string) => {
  router.push({name})
}

// 获取 Index 路由
const indexRouter = ref<RouteRecordRaw[] | undefined>()
const indexRouterFun = () => {
  router.getRoutes().map(item => {
    if (item.name === "Index") {
      indexRouter.value = item.children
    }
  })
}
indexRouterFun()

// i18n
const language = ref('zh-cn')
const locale = computed(() => (language.value === 'zh-cn' ? zhCn : en))
</script>

<template>
  <el-config-provider :locale="locale">
    <div class="main">
      <div class="left">
        <div class="left-top">
          <div class="left-top-item left-top-logo">
            <img src="@/assets/icon.svg" alt="panel logo">
          </div>
          <template v-for="(item, index) in indexRouter" :key="index">
            <div
                class="left-top-item"
                :class="[currentRoute.name == item.name?'left-top-item-active':'']"
                @click="goRouter(<string>item.name)"
            >
              <div class="left-top-item-icon">
                <i-ep-Histogram v-if="item.name=='Interface'"/>
                <i-ep-List v-else-if="item.name=='Test'"/>
                <i-ep-Tools v-else-if="item.name=='System'"/>
              </div>
              <div class="left-top-item-text">{{ item.meta!.title }}</div>
            </div>
          </template>
        </div>
        <div class="left-bottom">
          <div class="left-top-item" @click="goPage(info.doc)">
            <div class="left-top-item-icon">
              <i-ep-Management/>
            </div>
            <div class="left-top-item-text">接口文档</div>
          </div>
          <div class="left-top-item" @click="goPage(info.git)">
            <div class="left-top-item-icon">
              <i-ep-ChromeFilled/>
            </div>
            <div class="left-top-item-text">项目地址</div>
          </div>
        </div>
      </div>
      <div class="right">
        <router-view></router-view>
      </div>
    </div>
  </el-config-provider>
</template>

<style scoped lang="less">
.main {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  @left-width: 80px;

  .left {
    width: @left-width;
    height: 100vh;
    padding: 20px 0 5px 0;
    box-sizing: border-box;
    background-color: #F9FAFB;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    border-right: 1px solid #eaecf3;

    &-top {
      display: flex;
      flex-direction: column;
      align-items: center;
      @color: rgba(16, 24, 40, 0.56);

      &-logo {
        width: 40px;
        height: 40px;
        box-sizing: border-box;
        border-radius: 6px;
        background-color: #fff;
        box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);

        img {
          width: 24px;
          height: 24px;
        }
      }

      &-item {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 0 5px;

        &:not(:last-child) {
          margin-bottom: 26px;
        }

        &-icon {
          color: @color;
          font-size: 20px;
          font-weight: bold;
          padding-bottom: 5px;
        }

        &-text {
          font-size: 12px;
          color: @color;
        }

        &-active {
          @color: #8c78e6;

          .left-top-item-icon {
            color: @color;
          }

          .left-top-item-text {
            color: @color;
          }
        }
      }
    }

    &-bottom {
      font-size: 10px;
    }
  }

  .right {
    width: calc(100vw - @left-width);
    height: 100vh;
    overflow: hidden;
  }
}
</style>
