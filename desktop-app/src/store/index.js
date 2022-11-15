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
      usersOnline : new Set(),
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
    },
    updateUsersOnline(state, usersOnline){
      state.usersOnline = usersOnline
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
         ws.addEventListener('open', (e)=> {console.log("Open the ws:", e)})
          ws.addEventListener('close', (e)=> {console.log("Close the ws:", e)})
         ws.addEventListener("message", (e) => {
            const data = JSON.parse(e.data);
            let onlineUsers = state.usersOnline
            switch (data.action) {
              case "chat":
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
                  dispatch('addNewChatMessage', data.chatMessage)
                }else{
                  dispatch('addUnreadMessage',chatId )
                }
                break;
               case "join":
                onlineUsers.add(data.message)
                commit('updateUsersOnline', onlineUsers)
                break
                case "left":
                  onlineUsers.delete(data.message)
                  commit('updateUsersOnline', onlineUsers)
                break
                case "onlineUsers":
                  data.onlineUsers.forEach(user => {
                    onlineUsers.add(user)
                  });
                  commit('updateUsersOnline', onlineUsers)
                break
                  default:
                console.log("Unhaddaled message in ws. Message:", data.message)
                break;
            }           
         })
         commit("updateWebSocketConn", ws) 
    }
  }
})