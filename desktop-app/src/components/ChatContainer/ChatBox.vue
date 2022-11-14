<template>
    <div class="chat-box-container">
        <h2 class="chat-header">{{openChatName}}</h2>
        <ChatContent/>
        <MessageInput/>
    </div>
</template>

<script>
import MessageInput from './ChatBox/MessageInput.vue';
import ChatContent from './ChatBox/ChatContent.vue'
export default{
    components:{MessageInput, ChatContent},
    computed:{
        openChatName(){ 
            let openChat = this.$store.state.chat.openChat
            if (openChat){
                return openChat.name
            }else{
                return "Click on user to open chat"
            }
        }
    }

}
</script>

<style scoped>
.chat-box-container{
    width:100%;
    height:100vh;
    background-color: var(--color-white);
    padding: 15px 15px 15px 15px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
}
.chat-header{
    color: var(--color-blue);
}
</style>