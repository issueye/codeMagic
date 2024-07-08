<template>
  <BsHeader :title="mdTitle" description="生成对应代码的模板脚本">
    <template #actions>
      <div class="flex w-[300px] justify-end">
        <div class="flex grow mr-3">
          <el-select v-model="dmId">
            <el-option
              v-for="(item, index) in tableData"
              :key="index"
              :label="item.title"
              :value="item.id"
            />
          </el-select>
          <el-button type="primary" @click="onTestRunClick">测试</el-button>
        </div>
        <el-button @click="onBackClick">返回</el-button>
      </div>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <div class="flex h-full">
        <div ref="editorContainer" class="mr-1" style="width: 70%"></div>
        <div class="w-[30%]">
          <Codemirror v-model:value="logData" :options="cmOptions" border />
        </div>
      </div>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { ref, onMounted, toRaw, Ref } from "vue";
import * as monaco from "monaco-editor";
import { useRouter, useRoute } from "vue-router";
import { GetCode, SaveCode } from "../../../wailsjs/go/main/Template";
import { ElMessage } from "element-plus";

import { List, TestRunCode } from "../../../wailsjs/go/main/DataModel";

import Codemirror, { createLog } from "codemirror-editor-vue3";
import { model } from "../../../wailsjs/go/models";
import { EventsOn } from "../../../wailsjs/runtime/runtime";

const router = useRouter();
const route = useRoute();

const dmId = ref("");
const tableData: Ref<model.DataModel[]> = ref([]);

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);

const getData = async () => {
  const data = await List("", 0, 0);
  tableData.value = data;
};

onMounted(() => {
  getData();
});

const editorContainer = ref<any>(null);
const editor = ref<any>(null);
onMounted(async () => {
  const code = await GetCode(id.value as string);
  console.log("code", code);

  editor.value = monaco.editor.create(editorContainer.value, {
    value: code || "",
    language: "javascript",
    folding: true, // 是否折叠
    foldingHighlight: true, // 折叠等高线
    foldingStrategy: "indentation", // 折叠方式  auto | indentation
    showFoldingControls: "always", // 是否一直显示折叠 always | mouseover
    disableLayerHinting: true, // 等宽优化
    emptySelectionClipboard: false, // 空选择剪切板
    selectionClipboard: false, // 选择剪切板
    automaticLayout: true, // 自动布局
    codeLens: false, // 代码镜头
    scrollBeyondLastLine: false, // 滚动完最后一行后再滚动一屏幕
    colorDecorators: true, // 颜色装饰器
    accessibilitySupport: "off", // 辅助功能支持  "auto" | "off" | "on"
    lineNumbers: "on", // 行号 取值： "on" | "off" | "relative" | "interval" | function
    lineNumbersMinChars: 5, // 行号最小字符   number
    readOnly: false, //是否只读  取值 true | false
    theme: "vs-dark",
    minimap: {
      enabled: true, // 是否启用预览图
    },
  });

  // console.log('editor', editor.value);

  // 设置快捷键 ctrl + s
  editor.value.addAction({
    id: "save",
    label: "保存",
    keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS],
    run: () => {
      let code = toRaw(editor.value).getValue();
      console.log("code", code);
      SaveCode(id.value as string, code).then(() => {
        ElMessage.success("保存成功");
      });
    },
  });
});

const logData = ref();

const cmOptions = {
  mode: "log",
  theme: "default",
  lineWrapping: true,
  foldGutter: true,
};

// 当组件销毁时，关闭 editer
// onUnmounted(() => {
//   editor.value.dispose();
// });

const onBackClick = () => {
  router.push("/templates");
};

const onTestRunClick = () => {
  logData.value = "";
  EventsOn("console", (data: any) => {
    logData.value += createLog(data, "info");
  });

  TestRunCode(dmId.value, id.value as string);
};
</script>