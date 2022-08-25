import { createRouter, createWebHistory } from 'vue-router'
import layoutVue from '@/layout/layout.vue'
import { useStorage } from '@/hooks'
import { MessagePlugin } from 'tdesign-vue-next'

const routes = [
  {
    path: '/',
    component: layoutVue,
    children: [
      {
        path: '/',
        name: 'main',
        mate: {
          title: '首页',
          KeepAlive: true,
        },
        component: () => import('@/views/main/main.vue'),
      },
      {
        path: 'article',
        name: 'article',
        mate: {
          title: '文章管理',
          KeepAlive: true,
        },
        component: () => import('@/views/article/article.vue'),
      },
      {
        path: 'writearticle',
        name: 'writearticle',
        mate: {
          title: '写文章',
          KeepAlive: true,
        },
        component: () => import('@/views/writeArticle/writeArticle.vue'),
      },
      {
        path: 'category',
        name: 'category',
        mate: {
          title: '分类管理',
          KeepAlive: true,
        },
        component: () => import('@/views/category/category.vue'),
      },
      {
        path: 'user',
        name: 'user',
        mate: {
          title: '用户管理',
          KeepAlive: true,
        },
        component: () => import('@/views/user/user.vue'),
      },
    ],
  },
  {
    path: '/login',
    name: 'login',
    mate: {
      title: '登录',
      KeepAlive: true,
    },
    component: () => import('@/views/login/login.vue'),
  },
]

// 路由前置守卫
const beforeEach = (to: any, from: any, next: any) => {
  // 判断是否登录
  if (to.name !== 'login' && !useStorage.getStorage('token')) {
    console.log('to', to)
    MessagePlugin.error('用户未登录, 或者登录已过期')
    next('/login')
  } else {
    next()
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(beforeEach)

export default router
