<template>
  <div class="m-4 p-4 bg-white rounded-lg">
    <Form v-model="formData" v-bind="formConfig">
      <template #btn>
        <t-button theme="primary" type="submit" style="margin-right: 10px" @click="onSubmit">提交</t-button>
        <t-button theme="default" variant="base" type="reset" style="margin-right: 10px">重置</t-button>
      </template>
    </Form>
    <div>
      <div class="table-header">
        <t-button theme="primary" @click="addUserHandle">添加</t-button>
      </div>
      <t-table
        row-key="index"
        :data="userList"
        :columns="columns"
        stripe
        bordered
        hover
        table-layout="auto"
        size="medium"
        :pagination="pagination"
        cell-empty-content="-"
        resizable
        :async-loading="loadingNode"
        :selected-row-keys="selectedRowKeys"
        @select-change="rehandleSelectChange"
      >
        <template #op-column>
          <p>操作</p>
        </template>
        <template #op="slotProps">
          <t-button theme="primary" @click="updateUserHandle(slotProps)">编辑</t-button>
          <t-button theme="danger" @click="deleteUserHandle(slotProps)" :style="{ marginLeft: '10px' }">删除</t-button>
        </template>
      </t-table>
    </div>
    <t-dialog v-model:visible="visibleModal" header="添加用户" mode="modal" draggable @confirm="onSubmit">
      <Form v-model="formData1" v-bind="formConfig" layout="vertical"></Form>
    </t-dialog>
  </div>
</template>

<script setup lang="ts">
import api from '@/api'
import { UserData } from '@/types/user'
import { PageData } from '@/types/page'
import { onMounted, reactive, ref } from 'vue'
import { MessagePlugin, PageInfo, SubmitContext } from 'tdesign-vue-next'
import Form from '@/base-ui/form/form.vue'
import { IForm, IFormItem } from '@/base-ui/types'

const formData = ref({})
const formData1 = ref<UserData>({
  username: '',
  password: '',
  age: 0,
  avatar: '',
  role: 0,
})
const visibleModal = ref(false)
const formType = ref('add')
const formConfig: IForm = {
  formItems: [
    {
      field: 'username',
      label: '用户名',
      type: 'input',
      placeholder: '请输入用户名',
    },
    {
      field: 'age',
      label: '年龄',
      type: 'input',
      placeholder: '请输入年龄',
    },
    {
      field: 'avatar',
      label: '头像',
      type: 'input',
      placeholder: '请输入头像',
    },
    {
      field: 'role',
      label: '角色',
      type: 'select',
      placeholder: '请选择角色',
      options: [
        {
          value: 0,
          label: '管理员',
        },
        {
          value: 1,
          label: '普通用户',
        },
      ],
    },
  ],
}

const selectedRowKeys = ref([])
const rehandleSelectChange = (value: any, { currentRowData }: any) => {
  console.log('value', value)
  selectedRowKeys.value = value
  console.log('first', selectedRowKeys.value)
}

const onSubmit = async () => {
  switch (formType.value) {
    case 'add':
      const { data } = await api.user.AddUser(formData1.value)
      if (data.code == 200) {
        MessagePlugin.success('添加成功')
        getDataList()
      }
      break
    case 'update':
      const id = formData1.value.ID || 0
      delete formData1.value.ID
      const { data: updataData } = await api.user.UpdateUser(id, formData1.value)
      if (updataData.code == 200) {
        MessagePlugin.success('修改成功')
      }
      break
  }
}

onMounted(() => {
  getDataList()
})

const page = ref<PageData>({
  pageNum: 1,
  pageSize: 5,
})
const loadingNode = ref('')

const pagination = reactive({
  defaultCurrent: page.value.pageNum,
  defaultPageSize: page.value.pageSize,
  total: 0,
  onChange: (pageInfo: PageInfo) => {
    page.value = {
      pageNum: pageInfo.current,
      pageSize: pageInfo.pageSize,
    }
    getDataList()
  },
})

const userList = ref<UserData[]>([])
const getDataList = async () => {
  loadingNode.value = 'loading'
  const { data } = await api.user.GetUsers(page.value)
  userList.value = data.data || []
  pagination.total = data.total || 0
  loadingNode.value = ''
}

const columns = [
  {
    colKey: 'row-select',
    type: 'multiple',
    width: 50,
  },
  {
    width: '100',
    colKey: 'index',
    title: '序号',
    // 对齐方式
    align: 'center',
    render(h: any, context: any) {
      const { rowIndex } = context
      return rowIndex + 1
    },
  },
  {
    colKey: 'username',
    title: '用户名',
  },
  {
    colKey: 'age',
    title: '年龄',
  },
  {
    colKey: 'avatar',
    title: '头像',
  },
  {
    colKey: 'role',
    title: '权限',
  },
  {
    colKey: 'op',
    width: 200,
    title: 'op-column',
    cell: 'op',
  },
]

const addUserHandle = () => {
  formType.value = 'add'
  visibleModal.value = true
  formData1.value = {
    username: '',
    password: '',
    age: 0,
    avatar: '',
    role: 0,
  }
}

const updateUserHandle = ({ row }: any) => {
  formType.value = 'update'
  visibleModal.value = true
  delete row.DeletedAt
  delete row.CreatedAt
  delete row.UpdatedAt
  delete row.password
  formData1.value = row
}

const deleteUserHandle = ({ row }: any) => {
  console.log('row', row)
}
</script>

<style scoped>
.table-header {
  @apply flex justify-end items-center mb-4;
}
</style>
