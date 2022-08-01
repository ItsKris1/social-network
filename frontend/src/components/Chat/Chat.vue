<template>

    <div class="messaging-wrapper" ref="messagingWrapper">
        <ChatBox v-for="chat in chats" :name="chat.name" :userid="chat.userid" @closeChat="removeChat"
                 :key="chat.userid"></ChatBox>

        <div class=" messaging" @click="toggleShowContent">

            <div class="messaging-header">
                <p>Messaging</p>
                <i class="uil uil-angle-up" :class="{ rotate: showContent }"></i>
            </div>

            <div class="messaging-content" v-show="showContent" v-if="usersIFollow.length > 0">
                <ul class="item-list">

                    <li v-for="user in usersIFollow">
                        <div class="user-picture small"></div>
                        <div class="item-text" @click.stop="openChat($event, user.id)">{{ user.nickname }}</div>
                    </li>

                </ul>

                <ul class="item-list">
                    <li>
                        <img src="../../assets/icons/users-alt.svg" alt="" class="small">
                        <div class="item-text">Group 1</div>
                    </li>
                    <li>
                        <img src="../../assets/icons/users-alt.svg" alt="" class="small">
                        <div class="item-text">Group 2</div>
                    </li>
                    <li>
                        <img src="../../assets/icons/users-alt.svg" alt="" class="small">
                        <div class="item-text">Group 3</div>
                    </li>
                </ul>
            </div>

        </div>
    </div>

</template>


<script>
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
    },


    methods: {
        async getUsersIFollow() {
            await this.$store.dispatch("getMyUserID");

            const response = await fetch('http://localhost:8081/following?userId=' + this.$store.state.id, {
                credentials: 'include'
            });

            const data = await response.json();

            this.usersIFollow = data.users;

        },


        toggleShowContent() {
            // console.log("Content toggled!")
            this.showContent = !this.showContent
        },

        openChat(e, userid) {
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
                "userid": userid,
                "name": e.target.textContent,
            });



            // this.$nextTick(() => {
            //     console.log("Viewport width", window.innerWidth)
            //     console.log("Messaging area width", this.$refs.messagingWrapper.clientWidth)

            // })




        },


        removeChat(name) {
            // console.log(name)
            // console.log("Removing chat")
            this.chats = this.chats.filter((chat) => {
                return chat.name !== name

            })




        }
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
</style>