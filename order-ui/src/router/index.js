import { createRouter, createWebHistory } from 'vue-router';
import Index from '@/pages/Home.vue';
import Login from '@/pages/Login.vue';
import Dashboard from '@/pages/dashboard/Index.vue';
import DemandList from '@/pages/dashboard/Demand.vue';
import AccountList from '@/pages/dashboard/Account.vue';
import CustomerList from '@/pages/dashboard/Customer.vue';
import NoFound from '@/pages/404.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Index,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/404',
    name: 'NoFound',
    component: NoFound,
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404',
  },
  {
    path: '/dashboard/index',
    name: 'Dashboard',
    component: Dashboard,
    meta: {
        requiresAuth: true
    },
  },
  {
    path: '/dashboard/account',
    name: 'Account',
    component: AccountList,
    meta: {
        requiresAuth: true
    },
  },
  {
    path: '/dashboard/demend',
    name: 'Demend',
    component: DemandList,
    meta: {
        requiresAuth: true
    },
  },
  {
    path: '/dashboard/customer',
    name: 'Customer',
    component: CustomerList,
    meta: {
        requiresAuth: true
    },
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token'); // 检查是否登录
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth); // 检查目标路由是否需要登录

  if (requiresAuth && !isAuthenticated) {
    next('/login'); // 如果需要登录但未登录，跳转到登录页面
  } else {
    next(); // 否则正常跳转
  }
});

export default router;