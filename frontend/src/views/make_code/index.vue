<template>
  <BsHeader :title="mdTitle" description="数据模型提供基础结构，通过绑定对应的代码模板生成代码">
    <template #actions>
      <el-button @click="onBackClick">返回</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <el-tabs v-model="activeName">
        <el-tab-pane>

        </el-tab-pane>
      </el-tabs>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import {
  GetModelInfo,
  SaveModelInfo,
} from "../../../wailsjs/go/main/DataModel";
import { model } from "../../../wailsjs/go/models";
import { Ref } from "vue";
import { ElMessage } from "element-plus";
// import { VxeColumnPropTypes, VxeTableEvents } from "vxe-table";

// 数据
const router = useRouter();
const route = useRoute();

const activeName = ref();


const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);

onMounted(() => {});

const tableData: Ref<model.ModelInfo[]> = ref([]);

const getData = async () => {
  const data = await GetModelInfo(id.value as string);
  console.log("data", data);
  tableData.value = data;
};

onMounted(() => {
  getData();
});

const onBackClick = () => {
  router.push("/");
};

</script>