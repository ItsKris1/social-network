import { createStore } from "vuex";
import router from "@/router";

export default createStore({
  //------------------------------------- state is like a variables, which hold a values.
  state: {
    id: "", // id is currently logged in user ID
    wsConn: null,
    newChatMessages: [],
    profileInfo: {},
    myFollowers: null,
    posts: {
      allposts: [],
      myposts: [],
      groupPosts: [],
    },
    users: {
      allusers: [],
    },
    groups: {
      allGroups: [],
    },

    recievedMessage: "",
  },
  //------------------------------------ getters is a way for check state values.
  getters: {
    getId(state) {
      return state.id;
    },
    allPosts(state) {
      return state.posts.allposts;
    },
    myPosts(state) {
      return state.posts.myposts;
    },
    groupPosts(state) {
      return state.posts.groupPosts;
    },
    userInfo(state) {
      return state.profileInfo;
    },
    allUsers(state) {
      return state.users.allusers;
    },
    allGroups(state) {
      return state.groups.allGroups;
    },
    filterUsers: (state) => (searchquery) => {
      if (searchquery === "") {
        return [];
      }
      let arr = [];
      state.users.allusers.filter((user) => {
        if (user.nickname.includes(searchquery)) {
          arr.push(user);
        }
      });
      return arr;
    },

    filterGroups: (state) => (searchquery) => {
      if (searchquery === "") {
        return [];
      }
      let arr = [];
      state.groups.allGroups.filter((group) => {
        if (group.name.includes(searchquery)) {
          arr.push(group);
        }
      });
      return arr;
    },

    getMyFollowersNames({ myFollowers }) {
      if (myFollowers === null) {
        return null
      }

      return myFollowers.map((follower) => {
        if (follower.nickname) {
          return follower.nickname
        } else {
          return follower.firstName + follower.lastName
        }
      })
    },


    getMessages: ({ newChatMessages, id }) => (receiverId) => {
      return newChatMessages.filter((e) => {
        return (e.receiverId === receiverId && e.senderId === id) || (e.receiverId === id && e.senderId === receiverId)
      })

    },
  },
  //-------------------------------------- mutations is a way for change state.
  mutations: {
    updatePosts(state, posts) {
      state.posts.allposts = posts;
    },
    updateMyPosts(state, myposts) {
      state.posts.myposts = myposts;
    },
    updateProfileInfo(state, userinfo) {
      state.profileInfo = userinfo;
    },
    updateAllUsers(state, users) {
      state.users.allusers = users;
    },
    updateAllGroups(state, groups) {
      state.groups.allGroups = groups;
    },

    updateMyFollowers(state, followers) {
      state.myFollowers = followers;
    },

    updateMyUserID(state, id) {
      state.id = id;
    },
    updateGroupPosts(state, posts) {
      state.posts.groupPosts = posts;
    },
    updateWebSocketConn(state, wsConn) {
      state.wsConn = wsConn
    },

    updateNewChatMessages(state, msgs) {
      state.newChatMessages = msgs
    },

    addNewChatMessage(state, msg) {
      state.newChatMessages.push(msg);
    }
  },
  //------------------------------------------Actions
  actions: {
    //fetch all posts of all users.
    async fetchPosts() {
      await fetch("http://localhost:8081/allPosts", {
        credentials: "include",
      })
        // .then((r=>console.log(r)))
        .then((res) => res.json())
        .then((json) => {
          // console.log(json);
          const posts = json.posts;
          this.commit("updatePosts", posts);
        });
    },
    //fetch current logged in user posts.
    async fetchMyPosts() {
      let id = "";
      await fetch("http://localhost:8081/currentUser", {
        //first get my ID
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => {
          // console.log("get id - ", json);
          id = json.users[0].id;
        });
      await fetch("http://localhost:8081/userPosts?id=" + id, {
        //then fetch all posts with this ID
        credentials: "include",
      })
        .then((r) => r.json())
        .then((r) => {
          const myposts = r.posts;
          this.commit("updateMyPosts", myposts);
          // console.log(myposts);
        });
      console.log("here")

      // .then((json) => console.log("get posts -", json));
    },

    async getMyUserID({ commit }) {

      await fetch("http://localhost:8081/currentUser", {
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => {
          // console.log("JSON response", json)
          commit("updateMyUserID", json.users[0].id)
        });
    },

    async getMyProfileInfo(context) {
      await context.dispatch("getMyUserID");
      const userID = context.state.id;
      await fetch("http://localhost:8081/userData?userId=" + userID, {
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => {
          let userInfo = json.users[0];
          // console.log(userInfo);
          this.commit("updateProfileInfo", userInfo);
          // console.log("userinfo -", json);
        });
    },
    async getAllUsers() {
      await fetch("http://localhost:8081/allUsers", {
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => {
          let users = json.users;
          this.commit("updateAllUsers", users);
          // console.log("allUsers:", json.users);
        });
    },
    async getAllGroups() {
      await fetch("http://localhost:8081/allGroups", {
        credentials: "include",
      })
        .then((r) => r.json())
        .then((json) => {
          let groups = json.groups;
          this.commit("updateAllGroups", groups);
          // console.log("Allgroups:", json.groups);
        });
    },

    async getMyFollowers(context) {
      await context.dispatch("getMyProfileInfo");
      const myID = context.state.profileInfo.id;

      const response = await fetch(`http://localhost:8081/followers?userId=${myID}`, {
        credentials: 'include'
      });

      const data = await response.json();

      context.commit("updateMyFollowers", data.users)

    },


    async getGroupPosts() {
      await fetch(
        "http://localhost:8081/groupPosts?groupId=" +
        router.currentRoute.value.params.id,
        {
          credentials: "include",
        }
      )
        .then((r) => r.json())
        .then((json) => {
          // console.log(json)
          let posts = json.posts;
          this.commit("updateGroupPosts", posts);
        });
    },

    createWebSocketConn({ commit, state }) {
      const ws = new WebSocket("ws://localhost:8081/ws");
      ws.addEventListener("open", () => {
        console.log("Connection has established")
      })

      ws.addEventListener("message", (e) => {
        const data = JSON.parse(e.data);
        console.log("WS Message", data)
        // state.recievedMessage = data.chatMessage.content;
        if (data.action == "chat") {
          console.log("Message received -> ", data.chatMessage.content)
          // state.chatMessages.push(data.chatMessage.content);


          commit("updateNewChatMessages", [...state.newChatMessages, data.chatMessage])
        }

      })

      ws.addEventListener("close", (e) => {
        console.log("Connection closed");
      })



      commit("updateWebSocketConn", ws)
    },

  },
  modules: {},
});
