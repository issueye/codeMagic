<template>
  <BsHeader title="代码模板" description="生成对应代码的模板脚本">
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
            prop="fileName"
            label="文件名称"
            width="150"
            show-overflow-tooltip
          />
          <el-table-column
            prop="fileType"
            label="目标文件类型"
            width="90"
            show-overflow-tooltip
          >
            <template #default="{ row }">
              <el-tag>{{ getTargetFileType(row.fileType) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="mark" label="备注" show-overflow-tooltip />
          <el-table-column
            label="操作"
            width="190"
            align="center"
            fixed="right"
          >
            <template v-slot="{ row }">
              <el-button
                type="primary"
                link
                size="small"
                @click="onEditScriptClick(row)"
                >编辑脚本</el-button
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
    @open="onOpen"
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
        <el-form-item label="文件名称" prop="fileName">
          <el-input
            v-model="dataForm.fileName"
            placeholder="请输入文件名称"
            :disabled="true"
            clearable
          />
        </el-form-item>
        <el-form-item label="文件类型" prop="fileType">
          <el-select v-model="dataForm.fileType" placeholder="请选择文件类型">
            <el-option :value="0" label="Go" />
            <el-option :value="1" label="Java" />
            <el-option :value="2" label="vue" />
            <el-option :value="3" label="javascript" />
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
import { nanoid } from "nanoid";
import {
  List,
  Create,
  Delete,
  Modify,
} from "../../../wailsjs/go/main/Template";
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
const dataForm = reactive<model.CodeTemplate>(
  model.CodeTemplate.createFrom({
    id: "",
    title: "",
    fileName: "",
    fileType: "",
    mark: "",
  })
);

onMounted(() => {});

const tableData: Ref<model.CodeTemplate[]> = ref([]);

const getData = async () => {
  const data = await List(form.condition, pageNum.value, pageSize.value);
  tableData.value = data;
};

// 重置表单数据
const resetForm = () => {
  let id = nanoid();
  dataForm.id = "";
  dataForm.title = "";
  dataForm.fileName = id;
  dataForm.fileType = 0;
  dataForm.mark = "";
};

const getTargetFileType = (value: number) => {
  switch (value) {
    case 0:
      return "golang";
    case 1:
      return "java";
    case 2:
      return "vue";
    case 3:
      return "javascript";
    default:
      return "";
  }
};

// 赋值表单数据
const setForm = (value: any) => {
  dataForm.id = value.id;
  dataForm.title = value.title;
  dataForm.fileName = value.fileName;
  dataForm.fileType = value.fileType;
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

const onEditClick = (value: any) => {
  operationType.value = 1;
  title.value = `[编辑]${nameTitle}`;
  setForm(value);
  visible.value = true;
};

const onEditScriptClick = (value: any) => {
  console.log("value", value);
  router.push({
    path: "/script_edit",
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
          visible.value = false;
          getData();
          break;
        }

        case 1: {
          await Modify(dataForm);
          // ElMessage.success(res.message);
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

// 弹窗打开时
const onOpen = async () => {
  // const res = await apiUserGroupList();
  // if (res.code === 200) {
  //   groupTableData.value = res.data;
  // }
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