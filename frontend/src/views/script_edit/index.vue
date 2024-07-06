<template>
    <BsHeader :title="mdTitle" description="生成对应代码的模板脚本">
        <template #actions>
            <el-button @click="onBackClick">返回</el-button>
        </template>
    </BsHeader>
    <BsMain>
        <template #body>
            <div ref="editorContainer" class="h-full"></div>
        </template>
    </BsMain>
</template>

<script setup lang="ts">
import { ref, onMounted, toRaw } from 'vue'
import * as monaco from 'monaco-editor'
import { useRouter, useRoute } from 'vue-router';
import { GetCode, SaveCode } from '../../../wailsjs/go/main/Template';
import { ElMessage } from 'element-plus';


const router = useRouter();
const route = useRoute();

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);

const editorContainer = ref<any>(null)
const editor = ref<any>(null)
onMounted(async () => {
    const code = await GetCode(id.value as string);
    console.log('code', code);

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
        theme: 'vs-dark',
        minimap: {
            enabled: true, // 是否启用预览图
        }
    })

    // console.log('editor', editor.value);

    // 设置快捷键 ctrl + s
    editor.value.addAction({
        id: 'save',
        label: '保存',
        keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS],
        run: () => {
            let code = toRaw(editor.value).getValue()
            // console.log('code', code)
            SaveCode(id.value as string, code).then(() => {
                ElMessage.success('保存成功');
            })
        }
    })
})

// 当组件销毁时，关闭 editer
// onUnmounted(() => {
//    editor.value.dispose()
// })


const onBackClick = () => {
    router.push("/templates");
}
</script>