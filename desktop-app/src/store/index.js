import { createStore } from 'vuex'
import chat from "@/store/modules/chat.js"

// Create a new store instance.
export default createStore({
    modules: {
    chat
  },
  state:{
      id: "", // id is currently logged in user ID
      wsConn: {},
    //   users: {
    //     allusers: [],
    //   },
    groups: {
    //   allGroups: [],
      userGroups: [],
    },
      dataLoaded: {
      userGroups: false
    }
  },
   getters:{
    // getId(state) {
    //   return state.id;
    // },
   },
  mutations: {
    updateMyUserID(state, id) {
      state.id = id;
    },
    updateWebSocketConn(state, wsConn) {
      state.wsConn = wsConn
    },
    updateUserGroups(state, userGroups) {
      state.groups.userGroups = userGroups
    },
    updateDataLoaded(state, data) {
      state.dataLoaded[data] = true;
    }
  },
  actions:{
    async isLoggedIn() {
        const response = await fetch('http://localhost:8081/sessionActive', {
            credentials: 'include'
        });
        const data = await response.json();
        if (data.message === "Session active") {
            return true
        } else {
            return false
        }
    },
     async getMyUserID({ commit }) {
        await fetch("http://localhost:8081/currentUser", {
            credentials: "include",
        })
            .then((r) => r.json())
            .then((json) => {
                commit("updateMyUserID", json.users[0].id)
            });
    },
    async getUserGroups(context) {
        const response = await fetch(`http://localhost:8081/userGroups`, {
            credentials: 'include'
        });

        const data = await response.json();
        context.commit("updateUserGroups", data.groups)
        context.commit("updateDataLoaded", "userGroups")
    },
    createWebSocketConn({ commit, dispatch, state }) {
        const ws = new WebSocket("ws://localhost:8081/ws");
        console.log("Create web socket:", ws)
         ws.addEventListener("message", (e) => {
            const data = JSON.parse(e.data);
            switch (data.action) {
              case "chat":
                console.log("Received chat message")
                var isChatOpen = false
                var chatId = ""
                if (data.chatMessage.type === "PERSON"){
                  chatId = data.chatMessage.senderId
                  if(state.chat.openChat && data.chatMessage.senderId === state.chat.openChat.receiverId){
                    isChatOpen = true
                  }
                }
                if (data.chatMessage.type === "GROUP" ){
                  if (state.chat.openChat && data.chatMessage.receiverId === state.chat.openChat.receiverId){
                    isChatOpen = true
                  }
                  chatId = data.chatMessage.receiverId
                }
                if (isChatOpen){
                  console.log("Chat open")
                  dispatch('addNewChatMessage', data.chatMessage)
                }else{
                  console.log("Chat closed")
                  dispatch('addUnreadMessage',chatId )
                }
                break;
              default:
                console.log("Message:", data)
                break;
            }           
         })
         commit("updateWebSocketConn", ws) 
    }
  }
})