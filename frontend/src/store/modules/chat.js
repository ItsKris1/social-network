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


        getUnreadMessagesCount: ({ state }) => (userId) => {
            const userUnreadMsgs = state.unreadMessages.filter((msg) => {
                msg.receiverId === userId
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

        updateUnreadChatMessages(state, unreadMsgs) {
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
            commit("updateUnreadChatMessages", unreadChatMsgs)
        }
    },
}


