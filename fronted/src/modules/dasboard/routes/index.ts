import type { RouteRecordRaw } from 'vue-router';

const DashboardRoutes: RouteRecordRaw = {
  path: '/dashboard',
  redirect: '/dashboard/home',
  component: () => import('@/layout/DashboardLayout.vue'),
  children: [
    {
      name: 'home',
      path: 'home',
      component: () => import('@/modules/dasboard/views/HomePage.vue'),
      meta: { requiresAuth: true }
    },
    {
      name: 'stats',
      path: 'stats',
      component: () => import('@/modules/dasboard/views/StatsPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      name: 'user',
      path: 'user',
      component: () => import('@/modules/dasboard/views/UserPage.vue'),
      meta: { requiresAuth: true }
    }
  ]
};

export default DashboardRoutes;
