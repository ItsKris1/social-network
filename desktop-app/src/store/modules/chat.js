
export default {
    state: () => ({
        openChatMessages:[],
        unreadMessages:{},
        openChat: null,
        chatUserList: [],
    }),
    getters: {

    },
    mutations: {
        updateOpenChatMessages(state, msgs) {
            state.openChatMessages = msgs
        },
        openNewChat(state, openChat) {
            state.openChat = openChat
        },
        updateChatUserList(state, userList) {
            state.chatUserList = userList
        },
        updateUnreadMessages(state, msgs) {
            state.unreadMessages = msgs
        },
    },
    actions: {
        async fetchUnreadMessages({state, commit}){
            const response = await fetch('http://localhost:8081/unreadMessages', {
                credentials: 'include'
            });
            const data = await response.json();
            if (data.type === "Error") {
                // state.unreadMsgsStatsFromDB = null;
            } else if (data.chatStats){
                let unreadMessages = state.unreadMessages
                data.chatStats.forEach(msg => {
                    unreadMessages[msg.id] = msg.unreadMsgCount
                    
                });
                commit("updateUnreadMessages", unreadMessages)
            }
        },
        async fetchChatUserList({rootState, commit, dispatch}) {
            await dispatch("getMyUserID");
           
            const response = await fetch('http://localhost:8081/chatList?userId=' + rootState.id, {
                credentials: 'include'
            });
            const data = await response.json();
            commit("updateChatUserList", data.users);
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
                console.log("Success on sending: ", data.chatMessage[0], data.chatMessage[0].receiverId)
                dispatch("addNewChatMessage", {payload: data.chatMessage[0], id:state.openChat.id});
                return true
            }else{
                return false
            }
        },
        addUnreadMessage({ commit, state }, payload){
            let unreadMessages = state.unreadMessages
            if(unreadMessages[payload]){
               unreadMessages[payload] = unreadMessages[payload]+1
            }else{
                unreadMessages[payload] = 1
            }
            commit("updateUnreadMessages", unreadMessages)
        },
         removeUnreadMessages({ commit, state }, payload){
            let unreadMessages = state.unreadMessages
            if(unreadMessages[payload]){
               delete unreadMessages[payload]
               commit("updateUnreadMessages", unreadMessages)
            }
        }
    },
}