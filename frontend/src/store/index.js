import { createStore } from "vuex";

export default createStore({
  // state is like a variables, which hold a values.
  state: {
    user: {},
    notifications: {
      isNotificationsOpen: false,
    },
    posts: {
      allposts: [],
    },
  },
  // getters is a way for check state values.
  getters: {
    allPosts(state) {
      return state.posts.allposts;
    },
  },
  // mutations is a way for change state.
  mutations: {
    updatePosts(state, posts) {
      state.posts.allposts = posts;
    },
  },
  actions: {
    //fetch all posts of all users.
    async fetchPosts() {
      await fetch("http://localhost:8081/allPosts", {
        credentials: "include",
      })
        // .then((r=>console.log(r)))
        .then((res) => res.json())
        .then((json) => {
          console.log(json);
          const posts = json.posts;
          this.commit("updatePosts", posts);
        });
    },
    //fetch current logged in user posts.
    async fetchMyPosts() {
      let id = "";
      await fetch("http://localhost:8081/currentUser", {
        //first get my ID
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => {
          console.log("get id - ", json);
          id = json.users[0].id;
        });
      await fetch("http://localhost:8081/userPosts?id=" + id, {
        //then fetch all posts with this ID
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => console.log("get posts -", json)); 
        // BUT response is only 'success' type, without any data
    },
  },
  modules: {},
});
