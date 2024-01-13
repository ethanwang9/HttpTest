<script setup lang="ts">
import {useRouter} from "vue-router";

// 路由
const router = useRouter()

// 返回上一页
const goBackUpPage = () => {
  router.back(-2)
}

</script>

<template>
  <div class="main">
    <h1>404 页面未找到</h1>
    <p>努力更新中，敬请期待...</p>
    <el-button type="primary" @click="goBackUpPage">
      <span>返回上一页</span>
      <template #icon>
        <i-ep-Back/>
      </template>
    </el-button>
  </div>
</template>

<style scoped lang="less">
.main {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;

  h1 {
    font-size: 3em;
    margin-bottom: 10px;
  }

  p {
    margin-bottom: 1.5em;
  }
}
</style>