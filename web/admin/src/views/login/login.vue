<template>
  <div class="login">
    <div class="login-top">
      <h2>欢迎登录</h2>
      <h3>GIN-BLOG后台管理系统</h3>
    </div>
    <div class="login-panel">
      <div class="title">GIN-BLOG后台管理系统</div>
      <t-form ref="form" :rules="loginRules" :data="formData" :colon="true" :label-width="0" @reset="onReset" @submit="onSubmit">
        <div class="value">用户名</div>
        <t-form-item name="username">
          <t-input v-model="formData.username" clearable placeholder="请输入账户名" autofocus>
            <template #prefix-icon>
              <desktop-icon />
            </template>
          </t-input>
        </t-form-item>
        <div class="value">密码</div>
        <t-form-item name="password">
          <t-input v-model="formData.password" type="password" clearable placeholder="请输入密码">
            <template #prefix-icon>
              <lock-on-icon />
            </template>
          </t-input>
        </t-form-item>

        <t-form-item style="padding-top: 8px">
          <t-button theme="primary" type="submit" block shape="round" class="btn">登录</t-button>
        </t-form-item>
        <div class="register">没有账号?去注册</div>
      </t-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import api from '@/api'
import { LoginData } from '@/types/user'
import { loginRules } from './login.rules'
import { MessagePlugin } from 'tdesign-vue-next'
import { useStorage } from '@/hooks'
import { useRouter } from 'vue-router'

const formData = ref<LoginData>({
  username: '',
  password: '',
})

const onReset = () => {
  formData.value.username = ''
  formData.value.password = ''
}

const router = useRouter()
const onSubmit = async ({ validateResult, firstError, e }: any) => {
  e.preventDefault()
  if (validateResult === true) {
    try {
      const { data } = await api.user.Login(formData.value)
      useStorage.setStorage('token', data.token)
      useStorage.setStorage('user', data.data)
      MessagePlugin.success('登录成功')
      setTimeout(() => {
        router.push('/')
      }, 1000)
    } catch (error) {}
  } else {
    MessagePlugin.warning(firstError)
    return
  }
}
</script>

<style scoped>
.login {
  @apply p-6;
}

.login-top {
  background-image: url(../../assets/img/login_bg.png);
  height: 520px;

  @apply bg-no-repeat bg-auto w-auto rounded-xl overflow-hidden flex items-center flex-col text-gray-50 text-6xl;
}

.login-top h2 {
  @apply mt-20;
}

.login-top h3 {
  @apply mt-10 text-4xl;
}

.login-panel {
  width: 400px;
  @apply bg-white rounded-xl shadow-lg p-6 fixed top-80 left-1/2 right-1/2 -translate-x-1/2 pb-10;
}

.title {
  @apply text-center font-bold text-xl my-4;
}

.btn {
  height: 40px;
  @apply mt-4;
}

.register {
  @apply text-center text-blue-500 text-xs my-4 cursor-pointer hover:text-blue-800 mt-4;
}
</style>
