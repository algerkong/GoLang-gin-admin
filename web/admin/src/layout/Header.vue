<template>
  <div class="header">
    <t-avatar :image="userInfo?.avatar" size="large" />
    <div class="user">{{ userInfo?.username }}</div>
    <t-button theme="danger" class="logout" @click="logout">
      <template #icon>
        <icon name="poweroff" />
      </template>
      退出
    </t-button>
    <t-dialog v-model:visible="isLogout" theme="warning" header="提示" body="确认退出登录吗?" @confirm="handleLogout" />
  </div>
</template>

<script setup lang="ts">
import { Icon, MessagePlugin } from 'tdesign-vue-next'
import { useStorage } from '@/hooks'
import { UserData } from '@/types/user'
import { ref } from 'vue'
import { useRouter } from 'vue-router';
const userInfo = ref<UserData>(useStorage.getStorage('user') || {})

const isLogout = ref(false)
const logout = () => {
  isLogout.value = true
}
const router =useRouter()
const handleLogout = ()=>{
    useStorage.clearStorage()
    router.push('/login')
    MessagePlugin.success('退出成功')
}
</script>

<style scoped>

.header {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  height: 100%;
  padding: 0 40px;
}

.user {
  font-size: 16px;
  margin-left: 10px;
  margin-right: 10px;
}
</style>
