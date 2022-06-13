import { createStore } from "vuex";

export default createStore({
  state: {
    notifications: {
      isNotificationsOpen: false,
    },
  },
  getters: {
    NOTIFICATIONS_STATE(state){
      return state.notifications.isNotificationsOpen
    }
  },
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
