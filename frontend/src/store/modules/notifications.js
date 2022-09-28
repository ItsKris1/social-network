export default {
    state: () => ({
        allNotifications: [],
    }),

    mutations: {
        updateAllNotifications(state, allNotifs) {
            state.allNotifications = allNotifs
        }
    },

    actions: {
        addNewNotification({ state, commit }, payload) {
            const allNotifs = state.allNotifications;
            allNotifs.push(payload);
            commit("updateAllNotifications", allNotifs)
        },

        removeNotification({ state, commit }, notifID) {
            const allNotifs = state.allNotifications.filter((notif) => notif.id !== notifID);
            commit("updateAllNotifications", allNotifs)
        }
    }
}