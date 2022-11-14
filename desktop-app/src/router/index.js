import { createRouter, createWebHashHistory} from "vue-router";
import LoginView from "../views/LoginView.vue"
import ChatView from "../views/ChatView.vue"
import store from "@/store";


const routes = [
     {
    path: "/",
    name: "login",
    component: LoginView,
  },
  {
    path: "/chat",
    name: "chat",
    component: ChatView,
  },
]

const router = createRouter({
    history: createWebHashHistory(),
  routes,
});

router.beforeEach(async (to) => {
  const isAuthenticated = await store.dispatch("isLoggedIn");

  if (!isAuthenticated && to.name !== "login") {
      return { name: "login" }
  }else if(isAuthenticated && to.name === "login"){
    store.dispatch("createWebSocketConn")
    return { name: "chat" }
  }
})

export default router;