import { createStore } from "vuex";

export default createStore({
  //------------------------------------- state is like a variables, which hold a values.
  state: {
    id: "",
    profileInfo: {},
    myFollowers: null,
    posts: {
      allposts: [],
      myposts: [],
    },
    users: {
      allusers: [],
    },
    groups: {
      allGroups: [],
    },
  },
  //------------------------------------ getters is a way for check state values.
  getters: {
    allPosts(state) {
      return state.posts.allposts;
    },
    myPosts(state) {
      return state.posts.myposts;
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
    }
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

      // .then((json) => console.log("get posts -", json));
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

    }
  },
  modules: {},
});
