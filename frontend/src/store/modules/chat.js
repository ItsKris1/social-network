export default {

    state: () => ({
        newChatMessages: [],
        newGroupChatMessages: [],

        unreadMessages: [],
        openChats: [],
    }),

    getters: {
        getMessages: ({ newChatMessages, newGroupChatMessages }, getters, { id }) => (receiverId, type) => {
            let messages = [];

            if (type === "PERSON") {
                messages = newChatMessages.filter((e) => {
                    return (e.receiverId === receiverId && e.senderId === id) || (e.receiverId === id && e.senderId === receiverId)
                })
            } else {
                messages = newGroupChatMessages.filter((msg) => {
                    // console.log(msg.receiverId === receiverId)
                    return (msg.receiverId === receiverId)
                })

                // console.log("Group messages returned", messages)
            }

            return messages
            // return newChatMessages.filter((e) => {
            //   return (e.receiverId === receiverId && e.senderId === id) || (e.receiverId === id && e.senderId === receiverId)
            // })
        },


        getUnreadMessagesCount({ unreadMessages }, getters, rootState) {
            const userUnreadMsgs = unreadMessages.filter((msg) => {
                return msg.receiverId === rootState.id
            })

            return userUnreadMsgs.length
        },

        getUnreadGroupMessagesCount: ({ unreadMessages }, getters) => (groupId) => {
            const userUnreadMsgs = unreadMessages.filter((msg) => {
                return msg.type === "GROUP" && msg.receiverId === groupId
            })

            return userUnreadMsgs.length
        }
    },

    mutations: {
        updateNewChatMessages(state, msgs) {
            state.newChatMessages = msgs
        },

        updateNewGroupChatMessages(state, msgs) {
            state.newGroupChatMessages = msgs
        },

        updateOpenChats(state, openChats) {
            state.openChats = openChats
        },

        updateUnreadMessages(state, unreadMsgs) {
            state.unreadMessages = unreadMsgs
        }

    },

    actions: {
        addNewChatMessage({ commit, state }, payload) {
            let newMessages;

            if (payload["type"] === "PERSON") {
                newMessages = [...state.newChatMessages, payload]
                commit("updateNewChatMessages", newMessages)
            } else {
                newMessages = [...state.newGroupChatMessages, payload]
                commit("updateNewGroupChatMessages", newMessages)
            }
        },

        addUnreadChatMessage({ commit, state }, payload) {
            const unreadChatMsgs = state.unreadMessages
            unreadChatMsgs.push(payload)
            commit("updateUnreadMessages", unreadChatMsgs)
        },


        removeUnreadMessages({ state, commit }, payload) {
            const unReadMsgs = state.unreadMessages.filter((msg) => {
                return msg.receiverId !== payload.receiverId
            })

            commit('updateUnreadMessages', unReadMsgs);
        }
    },
}


