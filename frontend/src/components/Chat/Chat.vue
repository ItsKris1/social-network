<template>

    <div class="messaging-wrapper">
        <!-- <component v-for="chat in chats"
                   :is="{ ...chat.ChatBox }"
                   :key="chat.id"
                   :name="chat.name"
                   @closeChat="removeChat">
        </component> -->

        <ChatBox v-for="chat in chats" :name="chat.name" @closeChat="removeChat" :key="chat.id"></ChatBox>

        <div class="messaging" @click="toggleShowContent">

            <div class="messaging-header">
                <p>Messaging</p>
                <i class="uil uil-angle-up" :class="{ rotate: showContent }"></i>
            </div>

            <div class="messaging-content" v-show="showContent">
                <ul class="item-list">
                    <li>
                        <div class="user-picture small"></div>
                        <div class="item-text" @click.stop="openChat">User 1</div>
                    </li>
                    <li>
                        <div class="user-picture small"></div>
                        <div class="item-text" @click.stop="openChat">User 2</div>
                    </li>
                    <li>
                        <div class="user-picture small"></div>
                        <div class="item-text" @click.stop="openChat">User 3</div>
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
            chatID: 0,
        }
    },

    methods: {
        toggleShowContent() {
            // console.log("Content toggled!")
            this.showContent = !this.showContent
        },

        openChat(e) {
            // console.log("Trying to add a chatbox")
            // console.log(e.target.textContent)

            const found = this.chats.some(chat => chat.name === e.target.textContent);
            if (!found) {
                this.chats.push({
                    id: this.chatID,
                    "name": e.target.textContent,
                    ChatBox
                });

                this.chatID++;

            }

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