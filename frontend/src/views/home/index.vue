<template>
  <BsHeader title="数据模型" description="通过数据模型生成代码">
    <template #actions>
      <el-button type="primary" @click="onAddClick">添加</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <el-form inline>
        <el-form-item label="检索">
          <el-input
            v-model="form.condition"
            placeholder="请输入检索内容"
            clearable
            @change="onChange"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onQryClick">查询</el-button>
        </el-form-item>
      </el-form>

      <div class="h-[calc(100% - 60px)]">
        <el-table
          border
          :data="tableData"
          size="small"
          highlight-current-row
          stripe
        >
          <el-table-column
            prop="id"
            label="编码"
            width="150"
            show-overflow-tooltip
          />
          <el-table-column
            prop="title"
            label="标题"
            width="150"
            show-overflow-tooltip
          />
          <el-table-column
            prop="makeType"
            label="生成类型"
            width="90"
            show-overflow-tooltip
          >
            <template v-slot="{ row }">
              <el-tag
                effect="plain"
                :type="row.makeType === 1 ? 'success' : 'danger'"
              >
                {{ row.makeType === 1 ? "数据源" : "手动生成" }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="tpIds" label="脚本模板" show-overflow-tooltip>
            <template v-slot="{ row }">
              <el-tag
                class="mr-1"
                v-for="item in row.tpIds"
                :key="item"
                effect="plain"
                >{{ getTPName(item) }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column prop="mark" label="备注" show-overflow-tooltip />
          <el-table-column
            label="操作"
            width="230"
            align="center"
            fixed="right"
          >
            <template v-slot="{ row }">
              <el-button
                type="primary"
                link
                size="small"
                @click="onMakeCodeClick(row)"
                >生成代码</el-button
              >
              <el-button
                type="primary"
                link
                size="small"
                @click="onEditDataModelClick(row)"
                >编辑模型</el-button
              >
              <el-button
                type="primary"
                link
                size="small"
                @click="onEditClick(row)"
                >编辑</el-button
              >
              <el-button
                type="danger"
                link
                size="small"
                @click="onDeleteClick(row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div style="margin-top: 10px">
        <el-pagination
          small
          background
          :current-page="pageNum"
          :page-size="pageSize"
          :page-sizes="[5, 10, 20]"
          :total="total"
          layout="total, sizes, prev, pager, next"
          @size-change="onSizeChange"
          @current-change="onCurrentChange"
        />
      </div>
    </template>
  </BsMain>

  <BsDialog
    :title="title"
    :width="500"
    :visible="visible"
    @close="onClose"
    @save="onSave"
  >
    <template #body>
      <el-form
        label-width="auto"
        :model="dataForm"
        :rules="rules"
        ref="dataFormRef"
      >
        <el-form-item label="标题" prop="title">
          <el-input
            v-model="dataForm.title"
            placeholder="请输入数据模型标题"
            clearable
          />
        </el-form-item>
        <el-form-item label="生成类型" prop="makeType">
          <el-select v-model="dataForm.makeType" placeholder="请选择生成类型">
            <el-option :value="0" label="手动生成" />
            <el-option :value="1" label="数据源" />
          </el-select>
        </el-form-item>
        <el-form-item label="脚本模板" prop="tpIds">
          <el-select
            v-model="dataForm.tpIds"
            placeholder="请选择脚本模板"
            multiple
            collapse-tags
          >
            <el-option
              v-for="(item, index) in tpTableData"
              :key="index"
              :value="item.id"
              :label="item.title"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="dataForm.mark"
            placeholder="请输入备注"
            type="textarea"
            :row="2"
            clearable
          />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import {
  List,
  Create,
  Delete,
  Modify,
} from "../../../wailsjs/go/main/DataModel";
import { List as TemplateList } from "../../../wailsjs/go/main/Template";
import { RunCode } from "../../../wailsjs/go/main/DataModel";
import { model } from "../../../wailsjs/go/models";
import { Ref } from "vue";
import { useRouter } from "vue-router";

const nameTitle = "数据模型";
// 标题
const title = ref("数据模型");
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据
const dataFormRef = ref();
// 路由
const router = useRouter();

const rules = reactive({
  name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
  account: [{ required: true, message: "请输入账户", trigger: "blur" }],
  groupId: [{ required: true, message: "请选择用户组", trigger: "blur" }],
  tpIds: [{ required: true, message: "请选择脚本模板", trigger: "blur" }],
});

// 分页
const pageNum = ref(1);
const pageSize = ref(10);
const total = ref(0);

// 查询表单
const form = reactive({
  condition: "",
});

// 弹窗表单
const dataForm = reactive<model.DataModel>(
  model.DataModel.createFrom({
    name: "",
    title: "",
    makeType: 0,
    mark: "",
  })
);

onMounted(async () => {
  tpTableData.value = await TemplateList("", 0, 0);
});

const tableData: Ref<model.DataModel[]> = ref([]);
const tpTableData: Ref<model.CodeTemplate[]> = ref([]);

const getData = async () => {
  const data = await List(form.condition, pageNum.value, pageSize.value);
  tableData.value = data;
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.title = "";
  dataForm.makeType = 0;
  dataForm.tpIds = [];
  dataForm.mark = "";
};

// 赋值表单数据
const setForm = (value: any) => {
  dataForm.id = value.id;
  dataForm.title = value.title;
  dataForm.makeType = value.makeType;
  dataForm.tpIds = value.tpIds;
  dataForm.mark = value.mark;
};

onMounted(() => {
  getData();
});

// 事件

/**
 * 添加事件
 */
const onAddClick = () => {
  operationType.value = 0;
  title.value = `[添加]${nameTitle}`;
  resetForm();
  visible.value = true;
};

const onChange = () => {
  getData();
};

const onQryClick = () => {
  getData();
};

// 获取脚本模板名称
const getTPName = (value: string): string => {
  return tpTableData.value.find((item: any) => item.id == value)?.title || "";
};

const onEditClick = (value: any) => {
  operationType.value = 1;
  title.value = `[编辑]${nameTitle}`;
  setForm(value);
  visible.value = true;
};

const onMakeCodeClick = (value: model.DataModel) => {
  // console.log("value", value.id);

  RunCode(value.id);
};

const onEditDataModelClick = (value: any) => {
  console.log("value", value);
  router.push({
    path: "/data_model",
    query: { id: value.id, title: value.title },
  });
};

const onDeleteClick = (value: any) => {
  console.log("value", value.value);

  ElMessageBox.confirm("请确认是否要删除数据？", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      await Delete(value.id);
      getData();
    })
    .catch(() => {
      ElMessage.info("取消删除");
    });
};

const onSave = () => {
  if (!dataFormRef.value) return;
  dataFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      switch (operationType.value) {
        case 0: {
          await Create(dataForm);
          ElMessage.success("创建成功");
          visible.value = false;
          getData();
          break;
        }
        case 1: {
          await Modify(dataForm);
          ElMessage.success("修改成功");
          visible.value = false;
          getData();
          break;
        }
      }
    } else {
      console.log("表单验证失败");
    }
  });
};

const onSizeChange = (value: number) => {
  pageSize.value = value;
  getData();
};

const onCurrentChange = () => {
  getData();
};

const onClose = () => {
  visible.value = false;

  if (!dataFormRef.value) return;
  dataFormRef.value.resetFields();
};
</script>

