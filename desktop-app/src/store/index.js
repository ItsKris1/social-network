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

         ws.addEventListener("message", (e) => {
            const data = JSON.parse(e.data);
            console.log("Message received:", data)
            if (data.action != "chat") return
            const isParticipantsChatOpen = state.chat.openChats.some((chat) =>{
                if (data.chatMessage.type === "PERSON" && data.chatMessage.senderId  === chat.receiverId) {
                        return true
                    }
                if (data.chatMessage.type === "GROUP" && data.chatMessage.receiverId === chat.receiverId) {
                    return true
                }
            })
            if (isParticipantsChatOpen) {
                // dispatch("addNewChatMessage", data.chatMessage)
                // dispatch("markMessageRead", data.chatMessage)
            } else {
                if (data.message === "NEW") {
                    dispatch("fetchChatUserList");
                }

                // dispatch("addUnreadChatMessage", data.chatMessage)
            }
            
         })
         commit("updateWebSocketConn", ws) 
    }
  }
})