import { createRouter, createWebHistory } from 'vue-router'

import SignInView from '../views/SignInView.vue'
import RegisterView from '../views/RegisterView.vue'

const routes = [
  {
    path: '/',
    name: 'sign-in',
    component: SignInView
  },
  {
    path: '/reg',
    name: 'register',
    component: () => import ('../views/RegisterView.vue')
  },
  // {
  //   path: '/reg',
  //   name: 'reg',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/RegisterView.vue')
  // },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
