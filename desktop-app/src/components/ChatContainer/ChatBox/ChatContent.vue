<template>
    <div class="chat-content" ref="contentDiv">
       <div class="message" v-for="(message, index) in allMessages" :style="msgPosition(message)" :key="index">
            <p class="message-author" v-if="displayName(message, index)">{{ message.sender.nickname }}</p>
            <p :class="getClass(message)">{{ message.content }}</p>
       </div> 
    </div>
</template>

<script>
import { mapState } from 'vuex';
export default{
    data() {
        return {

        };
    },
    computed:{
        allMessages(){
            let openChat = this.$store.state.chat.openChat
            if (!openChat) {return []}
            let messages = this.$store.state.chatStack[openChat.id]
             return messages? messages :[];
        },
        ...mapState({
            myID: state => state.id
        })
    },
    methods:{
        // determines whether to display sender name in chatbox
        displayName(message, index) {
            let isSentMsg = message.senderId === this.myID;
            if (isSentMsg || message.type =="PERSON") {
                return false;
            }
            if (index < 1) {
                return true;
            }
            let isSequentMsg = message.senderId === this.allMessages[index - 1].senderId;
            if (isSequentMsg) {
                return false;
            }
            return true;
        },
        getClass(message) {
            let isSentMsg = message.senderId === this.myID;
            return isSentMsg ? { "sent-message": true } : { "recieved-message": true };
        },
        msgPosition(message) {
            let isSentMsg = message.senderId === this.myID;
            return {
                alignSelf: isSentMsg ? "flex-end" : "flex-start"
            };
        }
    }
}
</script>
<style scoped>
.chat-content {
    flex: 1;
    color: var(--color-lg-black);
    font-size: 14px;
    overflow-y: auto;
    overscroll-behavior: contain;
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.sent-message,
.recieved-message {
    /* padding: var(--msg-padding);
    border-radius: var(--msg-border-rad); */
     padding: 8px;
    border-radius: 10px;
    word-break: break-word;
    display: inline-block;
}

.message {
    max-width: 80%;
}
.recieved-messages {
    align-items: flex-start;
}
.recieved-message {
    background-color: rgb(212, 212, 212);
}
.message-author {
    padding-bottom: 3.5px;
}

.sent-message {
    background-color: rgb(201, 201, 201);
}
</style>