import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/403',
    component: () => import('@/views/403'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '主页', icon: 'dashboard' }
    }]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/shoplist',
    component: Layout,
    redirect: '/shoplist',
    children: [{
      path: 'shoplist',
      name: 'Shoplist',
      component: () => import('@/views/table/shopList'),
      meta: { title: '商铺列表', icon: 'table', roles: ['admin'] }
    }]
  },
  {
    path: '/userlist',
    component: Layout,
    redirect: '/userlist',
    children: [{
      path: 'userlist',
      name: 'Userlist',
      component: () => import('@/views/table/userList'),
      meta: { title: '用户列表', icon: 'table', roles: ['admin'] }
    }]
  },
  {
    path: '/itemlist',
    component: Layout,
    redirect: '/itemlist',
    children: [{
      path: 'itemlist',
      name: 'Itemlist',
      component: () => import('@/views/table/productList'),
      meta: { title: '商品列表', icon: 'table', roles: ['admin'] }
    }]
  },
  {
    path: '/recordlist',
    component: Layout,
    redirect: '/recordlist',
    children: [{
      path: 'recordlist',
      name: 'Recordlist',
      component: () => import('@/views/table/recordList'),
      meta: { title: '记录列表', icon: 'table', roles: ['admin'] }
    }]
  },
  {
    path: '/algo',
    component: Layout,
    redirect: '/algo',
    children: [{
      path: 'algorithm',
      name: 'Algorithm',
      component: () => import('@/views/algo/algo'),
      meta: { title: '模型算法', icon: 'skill', roles: ['admin'] }
    }]
  },
  {
    path: '/shopitemlist',
    component: Layout,
    redirect: '/shopitemlist',
    children: [{
      path: 'shopitemlist',
      name: 'Shopitemlist',
      component: () => import('@/views/table/index'),
      meta: { title: '商品列表', icon: 'table', roles: ['shop'] }
    }]
  },
  {
    path: '/useritemlist',
    component: Layout,
    redirect: '/useritemlist',
    children: [{
      path: 'useritemlist',
      name: 'Useritemlist',
      component: () => import('@/views/table/userItemList'),
      meta: { title: '商品列表', icon: 'table', roles: ['user'] }
    }]
  },
  {
    path: '/userinfo',
    component: Layout,
    redirect: '/userinfo',
    children: [{
      path: 'userinfo',
      name: 'Userinfo',
      component: () => import('@/views/form/userinfo'),
      meta: { title: '个人信息', icon: 'form', roles: ['user'] }
    }]
  },
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'https://github.com/RyanMrW/somewhere',
        meta: { title: '源码链接', icon: 'github', roles: ['admin'] }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
