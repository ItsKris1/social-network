<template>
    <div class="chat-wrapper">
        <ChatUsers />
        <ChatBox/>
        <SearchSection v-if="chatOpen"/>
    </div>
</template>

<script>
import ChatUsers from "./ChatContainer/ChatUsers.vue"
import ChatBox from "./ChatContainer/ChatBox.vue"
import SearchSection from './ChatContainer/SearchSection.vue'

export default {
    name:"ChatContainer",
    components:{ChatUsers, ChatBox, SearchSection},
    computed:{
        chatOpen(){
            return this.$store.state.chat.openChat? true:false;
        }
    },
    created(){
        // fetch unread messages from back end
        this.$store.dispatch("fetchUnreadMessages");
    }
}
</script>

<style>
.chat-wrapper{
    display:flex;
    justify-content: flex-start;
    flex-wrap: nowrap;
    width:100%;
}
</style>