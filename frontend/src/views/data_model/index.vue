<template>
  <BsHeader :title="mdTitle" description="数据模型信息">
    <template #actions>
      <el-button @click="onBackClick">返回</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <div class="mb-3">
        <el-button type="primary" @click="onAddOneCellClick(-1)"
          >新增一行</el-button
        >
        <el-button type="primary" @click="onSaveClick">保存</el-button>
      </div>

      <div class="h-[calc(100% - 60px)]">
        <vxe-table
          border
          show-overflow
          keep-source
          size="mini"
          ref="tableRef"
          :data="tableData"
          :menu-config="menuConfig"
          :edit-rules="validRules"
          max-height="550"
          :edit-config="{ trigger: 'click', mode: 'cell', showStatus: true }"
          @menu-click="onCtxMenuClickEvent"
        >
          <vxe-column type="seq" title="序号" width="70"></vxe-column>
          <vxe-column field="title" title="标题" :edit-render="{}">
            <template #edit="{ row }">
              <el-input v-model="row.title" placeholder="请输入名称" />
            </template>
          </vxe-column>
          <vxe-column field="name" title="名称" :edit-render="{}">
            <template #edit="{ row }">
              <el-input v-model="row.name" placeholder="请输入名称" />
            </template>
          </vxe-column>
          <vxe-column
            field="columnType"
            title="数据类型"
            :edit-render="{}"
            width="180"
          >
            <template #edit="{ row }">
              <el-select v-model="row.columnType" placeholder="请选择数据类型">
                <el-option value="int" label="int" />
                <el-option value="nvarchar" label="nvarchar" />
                <el-option value="bool" label="bool" />
                <el-option value="text" label="text" />
                <el-option value="datetime" label="datetime" />
                <el-option value="decimal" label="decimal" />
              </el-select>
            </template>
          </vxe-column>
          <vxe-column field="size" title="长度" :edit-render="{}" width="180">
            <template #edit="{ row }">
              <el-input-number v-model="row.size" controls-position="right" />
            </template>
          </vxe-column>
        </vxe-table>
      </div>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import {
  GetDataModelInfo,
  SaveModelInfo,
} from "../../../wailsjs/go/main/DataModel";
import { model } from "../../../wailsjs/go/models";
import { Ref } from "vue";

// 数据
const tableRef = ref();

const router = useRouter();
const route = useRoute();

const menuConfig = reactive({
  body: {
    options: [
      [
        { code: "insertAt", name: "插入", disabled: false },
        { code: "remove", name: "删除", disabled: false },
      ],
    ],
  },
});

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);

const validRules = ref({
  title: [{ required: true, message: "标题不能为空" }],
  name: [{ required: true, message: "名称不能为空" }],
  columnType: [{ required: true, message: "数据类型不能为空" }],
});

onMounted(() => {});

const tableData: Ref<model.ModelInfo[]> = ref([]);

const getData = async () => {
  const data = await GetDataModelInfo(id.value as string);
  console.log("data", data);

  tableData.value = data;
};

onMounted(() => {
  getData();
});

const onBackClick = () => {
  router.push("/");
};

const onSaveClick = () => {
  validEvent();
};

const validEvent = async () => {
  const $table = tableRef.value;
  if ($table) {
    const errMap = await $table.validate();
    if (!errMap) {
      console.log("tableData", tableData.value);

      await SaveModelInfo(id.value as string, tableData.value);
      getData();
    }
  }
};

const onAddOneCellClick = async (row: number) => {
  const $table = tableRef.value;
  if ($table) {
    const record = {
      title: "",
      name: "",
      columnType: "varchar",
    };
    const { row: newRow } = await $table.insertAt(record, row);
    await $table.setEditCell(newRow, "name");
  }
};

const onCtxMenuClickEvent = ({ menu, row, column }) => {
  const $table = tableRef.value;
  if ($table) {
    switch (menu.code) {
      case "insertAt": {
        $table.insertAt({}, row || -1).then(({ row }) => {
          $table.setEditCell(row, column || "name");
        });
        break;
      }
      case "remove":
        $table.remove(row);
        break;
    }
  }
};
</script>