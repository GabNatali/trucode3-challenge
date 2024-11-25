import authApi from '@/modules/auth/actions/authApi';
import AuthRoutes from '@/modules/auth/routes'
import { useAuthStore } from '@/stores/auth.store';
import DashboardRoutes from '@/modules/dasboard/routes';
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:'/signed-in-redirect',
      redirect:'/dashboard/home',
    },
    {
      path: '/:pathMatch(.*)*',
      component: () => import('@/views/Error404Page.vue')
    },

    AuthRoutes,
    DashboardRoutes
  ],
})

router.beforeEach(async(to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('accessToken');
  if (to.meta.requiresAuth && !isAuthenticated) {
    return next({ name: 'sign-in' });
  }

  if (to.name === 'sign-in' && isAuthenticated) {
    return next({ name: 'home' });
  }


  const authStore = useAuthStore();

  if (isAuthenticated && !authStore.userDataResponse) {
    try {
      const { data } = await authApi.getMe();
      authStore.userDataResponse = data;
    } catch (error) {
      console.error('Error loading user data', error);
      return;
    }
  }

  next();
});

export default router
