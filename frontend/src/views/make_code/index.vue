<template>
  <BsHeader
    :title="mdTitle"
    description="数据模型提供基础结构，通过绑定对应的代码模板生成代码"
  >
    <template #actions>
      <el-button @click="onBackClick">返回</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <el-tabs v-model="activeName" class="h-full">
        <el-tab-pane
          v-for="(item, index) in tableData"
          :label="item.title"
          :key="index"
          :name="item.id"
          style="height: calc(100% - 5px)"
        >
          <div
            class="h-full border border-solid border-[#d9d9d9]"
            style="height: 100%"
          >
            <Codemirror :value-model:value="logData" :options="cmOptions" />
          </div>
        </el-tab-pane>
      </el-tabs>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import Codemirror, { createLog } from "codemirror-editor-vue3";

import { List } from "../../../wailsjs/go/main/Template";
import { Ref } from "vue";
import { ElMessage } from "element-plus";
import { repository, model } from "../../../wailsjs/go/models";
// import { VxeColumnPropTypes, VxeTableEvents } from "vxe-table";

// 数据
const router = useRouter();
const route = useRoute();

const activeName = ref();

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);
const tpIds = ref(route.query["tpIds"]);
console.log("id", id.value);
console.log("tpIds", tpIds.value);

const logData = ref("=========测试日志==========\n");

const cmOptions = {
  mode: "fclog",
  theme: "default",
  lineWrapping: true,
  foldGutter: true,
};

const tableData: Ref<model.CodeTemplate[]> = ref([]);

const getData = async () => {
  let send = repository.RequestTemplateQuery.createFrom({
    ids: tpIds.value,
  });
  const data = await List(send, 0, 0);
  console.log("data", data);
  tableData.value = data;

  if (tableData.value.length > 0) {
    activeName.value = tableData.value[0].id;
  }
};

onMounted(() => {
  getData();
});

const onBackClick = () => {
  router.push("/");
};
</script>


<style scoped>
::v-deep(.el-tabs__content) {
  height: calc(100% - 60px);
}
</style>