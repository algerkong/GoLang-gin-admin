<template>
  <div>
    <t-layout class="layout">
      <t-aside>
        <Menu/>
      </t-aside>
      <t-layout>
        <t-header>
          <Header />
        </t-header>
        <t-content>
          <router-view v-slot="{ Component }">
            <component :is="Component" />
          </router-view>
        </t-content>
        <t-footer>Copyright @ 2019-{{ new Date().getFullYear() }} Tencent. All Rights Reserved</t-footer>
      </t-layout>
    </t-layout>
  </div>
</template>

<script setup lang="ts">
import Header from '@/layout/header.vue'
import Menu from '@/layout/menu.vue'
</script>

<style scoped>
.layout {
  height: 100vh;
}
</style>
