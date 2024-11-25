import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:'/signed-in-redirect',
      redirect:'/dashboard/home'
    },
    {
      path: '/:pathMatch(.*)*',
      component: () => import('@/views/Error404Page.vue')
    },

  ],
})


export default router
