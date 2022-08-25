<template>
  <t-menu v-model="expanded" theme="light" value="dashboard" expand-mutex height="550px" :onChange="change">
    <template #logo>
      <img width="136" class="logo" src="https://www.tencent.com/img/index/menu_logo_hover.png" alt="logo" />
    </template>
    <t-menu-item value="main">
      <template #icon>
        <t-icon name="dashboard" />
      </template>
      仪表盘
    </t-menu-item>
    <t-submenu value="article1" title="文章管理">
      <template #icon>
        <t-icon name="file-word" />
      </template>
      <t-menu-item value="writearticle">
        <span>写文章</span>
      </t-menu-item>
      <t-menu-item value="article">
        <span>文章列表</span>
      </t-menu-item>
    </t-submenu>
    <t-menu-item value="category">
      <template #icon>
        <t-icon name="server" />
      </template>
      分类列表
    </t-menu-item>
    <t-menu-item value="user">
      <template #icon>
        <t-icon name="user" />
      </template>
      用户列表
    </t-menu-item>
  </t-menu>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const expanded = ref<string | symbol>('')

const router = useRouter()
const route = useRoute()
// 获取当前路由的路径
onMounted(() => {
  expanded.value = route.name || ''
})

const change = (e: any) => {
  if ((e == 'main')) {
    router.push('/')
    return
  }
  router.push('/' + e)
}
</script>

<style scoped></style>
