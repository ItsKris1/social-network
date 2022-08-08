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


        getUnreadMessagesCount: ({ unreadMessages }, getters, { id }) => (userId) => {
            // console.log("uid", userId)
            const userUnreadMsgs = unreadMessages.filter((msg) => {
                // console.log(msg.receiverId)
                return msg.senderId === userId && msg.receiverId === id
            })


            // console.log(userUnreadMsgs.length)
            return userUnreadMsgs.length
        },

        getUnreadGroupMessagesCount: ({ unreadMessages }, getters) => (groupId) => {
            const userUnreadMsgs = unreadMessages.filter((msg) => {
                return msg.receiverId === groupId
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
            let unreadMsgs;
            // console.log(payload)
            if (payload.type === "GROUP") {
                unreadMsgs = state.unreadMessages.filter((msg) => {
                    if (msg.receiverId === payload.receiverId) {
                        return false
                    } else {
                        return true
                    }

                })
            } else {
                unreadMsgs = state.unreadMessages.filter((msg) => {
                    if (msg.type === "PERSON" && msg.senderId === payload.receiverId) {
                        return false
                    } else {
                        return true
                    }
                })
            }

            commit('updateUnreadMessages', unreadMsgs);
        }
    },
}


