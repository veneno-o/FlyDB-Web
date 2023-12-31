import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login/index.vue';
import Case from '../views/Case/index.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/',
      name: 'case',
      component: Case,
    },
    {
      path:'/:catchAll(.*)',
      name:"404",
      component:()=>(import("../views/404.vue"))
    }
  ],
});

export default router;
