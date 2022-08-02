<template>
    <div class="chatbox-wrapper">
        <div class="header">
            <p>{{ this.name }}</p>
            <i class="uil uil-times close" @click.stop="$emit('closeChat', this.name)"></i>
        </div>
        <div class="content" ref="contentDiv">
            <!-- <div class="message" v-for="message in allMessages" :style="getStyle(message)">
                <p v-if="!message.sequentMessage && allMessages.length > 0 && message.type === 'recieved'">User 1</p>
                <p :class="getClass(message)">{{ message.msg }}</p>
            </div> -->

            <div class="message" v-for="message in allMessages" :style="msgPosition(message)">
                <!-- <p v-if="!message.sequentMessage && allMessages.length > 0 && message.type === 'recieved'">User 1</p> -->
                <p>{{ message.content }}</p>
            </div>

            <!-- <div class="message" v-for="message in recentMessages" :style="msgPosition(message)">
                <p>{{ message.content }}</p>
            </div> -->

        </div>
        <div class=" send-message">
            <form @submit.prevent="sendMessage" autocomplete="off">
                <input type="text" name="sent-message" id="sent-message__input" placeholder="Send a message"
                       ref="sendMessageInput">
                <button type="submit"><i class="uil uil-message"></i></button>
            </form>

            <!-- <button type="button" @click="recieveMsg">Recieve msg</button> -->
        </div>
    </div>

</template>

<script>

export default {
    props: ["name", "userid"],
    emits: ['closeChat'],

    data() {
        return {
            previousMessages: [],
        }
    },

    created() {
        this.getPreviousMessages();
    },

    computed: {
        allMessages() {
            return [...this.previousMessages, ...this.$store.state.newChatMessages]

        }
    },

    watch: {
        allMessages() {
            this.$nextTick(() => {
                this.$refs.contentDiv.scrollTop = this.$refs.contentDiv.scrollHeight;

            })

        }
    },

    methods: {

        async getPreviousMessages() {
            const response = await fetch("http://localhost:8081/messages", {
                credentials: 'include',
                method: "POST",
                body: JSON.stringify({
                    type: "PERSON",
                    receiverId: this.userid
                })
            })

            const data = await response.json();
            console.log("/messages data", data)
            this.previousMessages = data.chatMessage;


        },

        async sendMessage() {
            const sendMessageInput = this.$refs.sendMessageInput;

            if (sendMessageInput.value === "") {
                return
            }

            const msgObj = {
                type: "PERSON",
                receiverId: this.userid,
                content: sendMessageInput.value
            }

            const response = await fetch("http://localhost:8081/newMessage", {
                body: JSON.stringify(msgObj),
                method: "POST",
                credentials: "include"
            });

            const data = await response.json();

            console.log("/newMessage data", data)


            this.$store.commit("addNewChatMessage", { content: sendMessageInput.value, receiverId: this.userid })
            sendMessageInput.value = "";

        },

        recieveMsg() {
            let isSequentMessage = false; // tracks if user sent more than two messages in a row

            if (this.allMessages.length !== 0 && this.allMessages[this.allMessages.length - 1].type === "recieved") {
                isSequentMessage = true;
            }

            this.allMessages.push({ msg: "Hey! How are you mate?", type: "recieved", sequentMessage: isSequentMessage })
        },

        // determines whether to display sender name in chatbox
        displayAuthor(message) {
            return this.allMessages.length > 0 && message.type === 'recieved' && message.sequentMessage === false;
        },

        getClass(message) {
            if (message.type === "recieved") {
                return { "recieved-message": true }
            } else {
                return { "sent-message": true }
            }
        },

        msgPosition(message) {
            const isSentMsg = message.receiverId === this.userid
            return {
                alignSelf: isSentMsg ? "flex-end" : "flex-start"
            }
        }

    },



}
</script>
<style scoped>
.chatbox-wrapper {
    height: 400px;
    width: 300px;
    display: flex;
    flex-direction: column;
    box-shadow: var(--container-shadow);
    border-radius: 5px 5px 0 0;
    overflow: hidden;
    --padding: 15px;
    --msg-border-rad: 13px;
    --msg-padding: 7.5px 10px;
    --msg-margin-b: 5px;
}

.header {
    display: flex;
    justify-content: space-between;
    background-color: var(--color-blue);
    color: var(--button-content);
    padding: 12px 20px;
}

.content {
    flex: 1;
    background-color: var(--color-white);
    padding: var(--padding);
    color: var(--color-lg-black);
    font-size: 14px;
    overflow-y: auto;
    overscroll-behavior: contain;
    display: flex;
    flex-direction: column;
    gap: 10px;
}

/* .recieved-messages,
.sent-messages {
    display: flex;
    flex-direction: column;
} */

/* .sent-messages>*:not(:last-child),
.recieved-messages>*:not(:last-child) {
    margin-bottom: var(--msg-margin-b);
} */

.sent-message,
.recieved-message {
    padding: var(--msg-padding);
    border-radius: var(--msg-border-rad);
    word-break: break-word;
}

.message {
    max-width: 80%;
}

.recieved-messages {
    align-items: flex-start;
}


.recieved-message {
    background-color: rgb(185, 185, 185);
    /* align-self: flex-start; */
}


.recieved-message .message-author {
    padding-bottom: 3.5px;
}


.sent-message {
    background-color: rgb(212, 212, 212);
    /* align-self: flex-end; */

}


/* SEND MESSAGE FORM */
.send-message {
    background-color: var(--color-white);
    padding: 10px;
}

.send-message form {
    display: flex;
    border-radius: 30px;
    padding: 10px 20px;
    overflow: hidden;
    box-shadow: var(--container-shadow);
}

.send-message form input {
    padding: 0;
    box-shadow: none;
    background-color: inherit;
}

.send-message form button {
    border: none;
    background-color: inherit;
    font-size: 1.25em;
}
</style>