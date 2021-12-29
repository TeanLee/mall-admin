import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import Layout from '@/layout'

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/product',
    hidden: true,
  },
  {
    path: '/product',
    name: 'product',
    title: '商品管理',
    component: Layout,
    redirect: '/product/list',
    children: [
      {
        path: 'list',
        name: 'product-list',
        title: '商品列表',
        meta: { title: 'Dashboard', icon: 'dashboard', affix: true },
        component: () => import('@/views/product/list')
      },
      {
        path: 'add',
        name: 'product-add',
        title: '添加商品',
        meta: { title: 'Documentation', icon: 'documentation', affix: true },
        component: () => import('@/views/product/add')
      }
    ]
  },
  {
    path: '/order',
    name: 'order',
    title: '订单管理',
    component: Layout,
    redirect: '/order/list',
    children: [
      {
        path: 'list',
        name: 'order-list',
        title: '订单管理',
        meta: { title: 'Documentation', icon: 'el-icon-s-order', affix: true },
        component: () => import('@/views/order/list')
      }
    ]
  },
  {
    path: '/category',
    name: 'category',
    title: '分类管理',
    component: Layout,
    redirect: '/category/list',
    children: [
      {
        path: 'list',
        name: 'category-list',
        title: '分类列表',
        meta: { title: 'Documentation', icon: 'el-icon-s-order', affix: true },
        component: () => import('@/views/category/list')
      }
    ]
  },
  {
    path: '/user',
    name: 'user',
    title: '用户管理',
    component: Layout,
    redirect: '/user/list',
    children: [
      {
        path: 'list',
        name: 'user-list',
        title: '用户列表',
        meta: { title: 'Documentation', icon: '', affix: true },
        component: () => import('@/views/user/list')
      },
      {
        path: 'summary',
        name: 'user-summary',
        title: '用户统计',
        meta: { title: 'Documentation', icon: '', affix: true },
        component: () => import('@/views/user/summary')
      }
    ]
  },
  {
    path: '/permission',
    name: 'permission',
    title: '权限管理',
    component: Layout,
    redirect: '/permission/role',
    children: [
      {
        path: 'role',
        name: 'permission-role',
        title: '权限管理',
        meta: { title: 'Documentation', icon: 'el-icon-s-order', affix: true },
        component: () => import('@/views/permission/role')
      }
    ]
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  { path: "/:catchAll(.*)", redirect: '/404', hidden: true }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
