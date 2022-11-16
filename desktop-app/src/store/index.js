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
      groups: {
        userGroups: [],
      },
      dataLoaded: {
      userGroups: false
    },
    chatStack:{}
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
    },
    updateChatStack(state,stack){
            state.chatStack = stack
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
    async fecthChatStack({state, commit}){
      let chatStack = {}
      state.chat.chatUserList.forEach(user => {
          fetch("http://localhost:8081/messages", {
                credentials: "include",
                method: "POST",
                body: JSON.stringify({type: "PERSON",receiverId: user.id})
            })
            .then((res)=> res.json())
            .then((res)=> chatStack[user.id] = res.chatMessage)
      });
      state.groups.userGroups.forEach(group=> {
          fetch("http://localhost:8081/messages", {
                credentials: "include",
                method: "POST",
                body: JSON.stringify({type: "GROUP",receiverId: group.id})
            })
            .then((res)=> res.json())
            .then((res)=> chatStack[group.id] = res.chatMessage)
      });
      commit('updateChatStack', chatStack)
      console.log("Combined list:", chatStack)
      },
      addNewChatMessage({ commit, state }, {payload, id}) {
        console.log("Add message: ", payload, id)
        let chatStack = state.chatStack;
        chatStack[id].push(payload)
       commit('updateChatStack', chatStack)
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
                  if(data.chatMessage.senderId === state.chat.openChat?.id){
                    isChatOpen = true
                  }
                }
                if (data.chatMessage.type === "GROUP" ){
                  if ( data.chatMessage.receiverId === state.chat.openChat?.id){
                    isChatOpen = true
                  }
                  chatId = data.chatMessage.receiverId
                }
                dispatch('addNewChatMessage', {payload:data.chatMessage, id:chatId})
                if (!isChatOpen){
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