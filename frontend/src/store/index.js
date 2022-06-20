import { createStore } from "vuex";

export default createStore({
  // state is like a variables, which hold a values.
  state: {
    user:{
      
    },
    notifications: {
      isNotificationsOpen: false,
    },
    posts:{
      allposts:[]
    }

  },
  // getters is a way for check state values.
  getters: {
    NOTIFICATIONS_STATE(state){
      return state.notifications.isNotificationsOpen
    },
    allPosts(state){
      return state.posts.allposts
    }


  },
  // mutations is a way for change state.
  mutations: {
    CHANGE_NOTIFICATIONS_STATE(state) {
      state.isNotificationsOpen = !state.isNotificationsOpen;
    },
    updatePosts(state, posts){
      state.posts.allposts = posts
    }
  },
  actions: {
    TOGGLE_NOTIFICATIONS({commit}) {
      commit('CHANGE_NOTIFICATIONS_STATE')
    },
    async fetchPosts(){
      await fetch("http://localhost:8081/allPosts", {
            credentials: 'include',
        })
            .then((res => res.json()))
            .then((json => {
                const posts = json.posts
                this.commit('updatePosts', posts)
            }))
            
    }
  },
  modules: {},
});
