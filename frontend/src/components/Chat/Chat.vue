<template>

    <div class="messaging-wrapper" ref="messagingWrapper">
        <ChatBox v-for="chat in chats" v-bind="chat" @closeChat="removeChat" :key="chat.receiverId"></ChatBox>

        <div class=" messaging" @click="toggleShowContent">

            <div class="messaging-header">
                <p>Messaging</p>
                <i class="uil uil-angle-up" :class="{ rotate: showContent }"></i>
            </div>

            <div class="messaging-content" v-show="showContent">
                <ul class="item-list" v-if="usersIFollow.type && usersIFollow.users !== null">

                    <li v-for="user in usersIFollow.users">
                        <div class="user">
                            <div class="user-picture small"></div>
                            <div class="item-text"
                                 @click.stop="openChat($event, { receiverId: user.id, type: 'PERSON' })">
                                {{ user.nickname }}</div>
                        </div>

                        <p class="unreadMessagesCount" v-if="getUnreadMessagesCount(user.id)">{{
                                getUnreadMessagesCount(user.id)
                        }}</p>

                    </li>

                </ul>

                <ul class="item-list" v-if="userGroups.type && userGroups.groups !== null">
                    <li v-for="group in userGroups.groups">

                        <div class="group">
                            <img src="../../assets/icons/users-alt.svg" alt="" class="small">

                            <div class="item-text"
                                 @click.stop="openChat($event, { receiverId: group.id, type: 'GROUP' })">
                                {{ group.name }}</div>
                        </div>

                        <p class="unreadMessagesCount" v-if="getUnreadGroupMessagesCount(group.id)">
                            {{ getUnreadGroupMessagesCount(group.id) }}</p>

                    </li>

                </ul>

                <p class="additional-info" v-if="usersIFollow.users === null && userGroups.groups === null">
                    No one to message with :(</p>
            </div>

        </div>
    </div>

</template>


<script>
import { handleError } from 'vue';
import { mapGetters, mapState } from 'vuex';
import ChatBox from './ChatBox.vue'
export default {
    name: '',
    components: { ChatBox },
    data() {
        return {
            showContent: false,
            chats: [],
            usersIFollow: [],

        }
    },

    mounted() { },

    created() {
        this.getUsersIFollow();
        this.$store.dispatch("getUserGroups");
        // this.getUnreadMessagesCount()
    },

    computed: {
        ...mapState({
            userGroups: state => state.groups.userGroups,
            unreadMessages: state => state.chat.unreadMessages,
        }),

        ...mapGetters(['getUnreadMessagesCount', 'getUnreadGroupMessagesCount']),
    },

    watch: {
        // unreadMessages: {
        //     handler: (newVal) => console.log("Val", newVal),
        //     deep: true
        // },

        // getUnreadMessagesCount(newVal) {
        //     console.log("Messages count", newVal)
        // }
    },

    methods: {
        async getUsersIFollow() {
            if (this.$store.state.id === "") {
                await this.$store.dispatch("getMyUserID");
            }

            const response = await fetch('http://localhost:8081/following?userId=' + this.$store.state.id, {
                credentials: 'include'
            });

            const data = await response.json();


            this.usersIFollow = data;

        },

        toggleShowContent() {
            // console.log("Content toggled!")
            this.showContent = !this.showContent
        },

        openChat(e, obj) {
            // console.log("Trying to add a chatbox")
            // console.log(e.target.textContent)
            const found = this.chats.some(chat => chat.name === e.target.textContent);
            if (found) {
                return
            };

            if (this.$refs.messagingWrapper.clientWidth + 300 > window.innerWidth) {
                this.chats.shift();
            }
            this.chats.push({
                "name": e.target.textContent,
                ...obj
            });

            this.$store.dispatch("removeUnreadMessages", { receiverId: obj.receiverId, type: obj.type })
            this.$store.commit("updateOpenChats", this.chats)
        },


        removeChat(name) {
            // console.log(name)
            // console.log("Removing chat")
            this.chats = this.chats.filter((chat) => {
                return chat.name !== name

            })

            this.$store.commit("updateOpenChats", this.chats)


        },



    }
}
</script>


<style scoped>
.messaging-wrapper {
    position: fixed;
    bottom: 0;
    right: 0;
    display: flex;
    align-items: flex-end;
}

.messaging {

    width: 300px;
    background-color: var(--color-white);
}

.messaging-header {
    display: flex;
    justify-content: space-between;
    padding: 12px 20px;
    background-color: var(--color-blue);
    border-top-left-radius: 5px;
    color: var(--color-white);
}

.messaging-header i {
    font-size: 18px;
}

.messaging-header p {
    font-weight: 300;
    letter-spacing: 0.35px;
}

.messaging-content {
    padding: 20px;
}

.messaging .item-list:first-of-type {
    padding-bottom: 20px;
}

.rotate {
    transform: rotate(180deg);
}

.user {
    display: flex;
    justify-content: space-between;
}

.item-list li {
    justify-content: space-between;
}

.user,
.group {
    display: flex;
    align-items: center;
    gap: 5px;
}
</style>