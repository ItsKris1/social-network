import { createRouter, createWebHistory } from "vue-router";
import Auth from "../components/Auth.vue";
// import SignIn from '../views/SignInView.vue'
// import RegisterView from '../views/RegisterView.vue'

const routes = [
  {
    path: "/",
    name: "auth",
    component: Auth,

  },
  {
    path: "/sign-in",
    name: "sign-in",
    component: () => import("../views/SignInView.vue"),
  },
  {
    path: "/reg",
    name: "register",
    component: () => import("../views/RegisterView.vue"),
  },
  {
    path: "/main",
    name: "mainpage",
    component: () => import("../views/MainView.vue"),
  },
  {
    path: "/profile/:id",
    name: "Profile",
    component: () => import("../views/ProfileView.vue"),
  },
  {
    path: "/group/:id",
    name: "Group",
    component: () => import("../views/GroupView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
