
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
        },
        async sendMessage({state,dispatch}, payload){
            const msgObj = {
                receiverId: state.openChat.id,
                content: payload,
                type: state.openChat.type
            };
             let response = await fetch("http://localhost:8081/newMessage", {
                body: JSON.stringify(msgObj),
                method: "POST",
                credentials: "include"
            });
            const data = await response.json();
            if (data.type == "Success"){
                dispatch("addNewChatMessage", data.chatMessage[0]);
                return true
            }else{
                return false
            }
        },
        addNewChatMessage({ commit, state }, payload) {
            let openMessages = state.openChatMessages;
            openMessages.push(payload)
            commit("updateOpenChatMessages", openMessages)
        },
    },
}