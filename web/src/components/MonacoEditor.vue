<script setup lang="ts">
import {onBeforeUnmount, onMounted, ref} from "vue";
import * as monaco from "monaco-editor";
import editorWorker from "monaco-editor/esm/vs/editor/editor.worker?worker";
import jsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker";

// editor 元素
const editorElm = ref<HTMLElement | null>(null)

// editor实例
let editor: monaco.editor.IStandaloneCodeEditor

// v-model
const model = defineModel()

// 参数
const props = defineProps({
  readonly: {
    type: Boolean,
    default: true,
  }
})

// editor 初始化环境信息 -- worker
self.MonacoEnvironment = {
  getWorker(_: string, label: string) {
    if (label === "json") {
      return new jsonWorker();
    }
    return new editorWorker();
  },
};

// 挂载 editor
onMounted(() => {
  editor = monaco.editor.create(editorElm.value as HTMLElement, {
    value: model.value as string,
    language: "json",
    theme: 'vs-dark',
    scrollBeyondLastLine: false,
    automaticLayout: true,
    autoIndent: true,
    glyphMargin: true,
    minimap: {
      enabled: false,
    },
    readOnly: props.readonly
  })

  editor.onDidChangeModelContent(() => {
    model.value = editor.getValue()
  })

})

// 取消挂载
onBeforeUnmount(() => {
  editor.dispose()
})

// 监听 model 动态修改
watchEffect(() => {
  let d: string = model.value as string | ""

  if (d === "null") {
    editor.setValue("")
  } else if (d.startsWith("!@#@")) {
    editor.setValue(d.substring(4))
    editor.getAction('editor.action.formatDocument')?.run()
    editor.setValue(editor.getValue())
  }
})
</script>

<template>
  <div ref="editorElm" class="editor"></div>
</template>

<style scoped lang="less">
.editor {
  width: 100%;
  height: 100%;
}
</style>