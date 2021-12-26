import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

import Layout from '@/layout'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/product',
    component: Layout,
    children: [
      {
        path: 'list',
        name: 'product-list',
        component: () => import('@/views/product/list')
      },
      {
        path: 'add',
        name: 'product-add',
        component: () => import('@/views/product/add')
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
