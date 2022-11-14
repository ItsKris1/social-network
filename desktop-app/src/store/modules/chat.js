export default {
    state: () => ({
        openChatMessages:[],
        
        openChat: null,
        chatUserList: [],
    }),
    getters: {

    },
    mutations: {
        updateOpenChatMessages(state, msgs) {
            console.log("Update open chat messages: ",  msgs)
            state.openChatMessages = msgs
        },
        openNewChat(state, openChat) {
            console.log("Open new chat: ", openChat)
            state.openChat = openChat
        },
        updateChatUserList(state, userList) {
            state.chatUserList = userList
        }
    },
    actions: {
        async fetchChatUserList({rootState, commit, dispatch}) {
            await dispatch("getMyUserID");
           
            const response = await fetch('http://localhost:8081/chatList?userId=' + rootState.id, {
                credentials: 'include'
            });
            const data = await response.json();
            commit("updateChatUserList", data.users);
        },
        async fecthChatMessages({state,commit}){
            const response = await fetch("http://localhost:8081/messages", {
                credentials: "include",
                method: "POST",
                body: JSON.stringify({
                    type: state.openChat.type,
                    receiverId: state.openChat.id
                })
                
            });
            const data = await response.json();
            let previousMessages = data.chatMessage ? data.chatMessage : [];
            console.log("MSg received from db: ", previousMessages)
            commit('updateOpenChatMessages', previousMessages)
        }
    },
}