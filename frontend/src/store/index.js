import { createStore } from "vuex";

export default createStore({
  // state is like a variables, which hold a values.
  state: {
    user:{
      
    },
    notifications: {
      isNotificationsOpen: false,
    },

  },
  // getters is a way for check state values.
  getters: {
    NOTIFICATIONS_STATE(state){
      return state.notifications.isNotificationsOpen
    }
  },
  // mutations is a way for change state.
  mutations: {
    CHANGE_NOTIFICATIONS_STATE(state) {
      state.isNotificationsOpen = !state.isNotificationsOpen;
    },
  },
  actions: {
    TOGGLE_NOTIFICATIONS({commit}) {
      commit('CHANGE_NOTIFICATIONS_STATE')
    }
  },
  modules: {},
});
