import { createRouter, createWebHistory } from 'vue-router'
import Auth from '../components/Auth.vue' 
// import SignIn from '../views/SignInView.vue'
// import RegisterView from '../views/RegisterView.vue'

const routes = [
  {
    path: '/',
    name: 'auth',
    component: Auth
  },
  {
    path: '/sign-in',
    name: 'sign-in',
    component: () => import ('../views/SignInView.vue')
  },
  {
    path: '/reg',
    name: 'register',
    component: () => import ('../views/RegisterView.vue')
  },
  {
    path: '/main',
    name: 'mainpage',
    component: () => import ('../views/MainView.vue')
  },
  {
    path: '/profile',
    name: 'profile',
    component: () => import ('../views/ProfileView.vue')
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
