import type { RouteRecordRaw } from 'vue-router';

const AuthRoutes: RouteRecordRaw = {
  path: '/',
  redirect: '/sign-in',
  component: () => import('@/layout/BlankLayout.vue'),
  children: [
    {
      name: 'sign-in',
      path: '/sign-in',
      component: () => import('@/modules/auth/views/SignInView.vue')
    },
    {
      name: 'sign-up',
      path: '/sign-up',
      component: () => import('@/modules/auth/views/SignUpView.vue')
    },
  ]
};

export default AuthRoutes;
