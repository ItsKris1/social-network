export default {
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

    async getUserGroups(context) {
        const response = await fetch(`http://localhost:8081/userGroups`, {
            credentials: 'include'
        });

        const data = await response.json();
        console.log("/getUserGroups data", data)
        // context.state.groups.userGroups.loaded = true;

        context.commit("updateUserGroups", data.groups)
        context.commit("updateDataLoaded", "userGroups")

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

    createWebSocketConn({ commit, dispatch, state }) {
        const ws = new WebSocket("ws://localhost:8081/ws");
        ws.addEventListener("open", () => {
            console.log("WS: Connection has established")
        })

        ws.addEventListener("message", (e) => {
            // console.log("New message")
            const data = JSON.parse(e.data);
            if (data.action == "chat") {
                // only broadcast messages when participants(sender and reciever) chat is open

                const isParticipantsChatOpen = state.chat.openChats.some((chat) => {
                    // chat is open with receiver
                    if (chat.receiverId === data.chatMessage.receiverId) {
                        return true
                    }

                    // chat is open with sender
                    // check type cuz we may have a scenario where we have the chat open with sender
                    // but not the group and that will cause to broadcast the message to GROUP chat;
                    if (data.chatMessage.type === "PERSON" && chat.receiverId === data.chatMessage.senderId) {
                        return true
                    }
                })
                if (isParticipantsChatOpen) {
                    // console.log("Dispatching a message..")
                    dispatch("addNewChatMessage", data.chatMessage)
                    dispatch("markMessageRead", data.chatMessage)
                } else {
                    // console.log("Unread msg..")
                    dispatch("addUnreadChatMessage", data.chatMessage)
                }
            } else if (data.action == "notification") {
                // console.log(data.notification)
                dispatch("addNewNotification", data.notification)
            }

        })

        ws.addEventListener("close", (e) => {
            console.log("Connection closed");
        })



        commit("updateWebSocketConn", ws)
    }

}