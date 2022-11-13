<template>
    <div class="chat-wrapper">
        <ChatUsers />
        <ChatBox/>
    </div>
</template>

<script>
import ChatUsers from "./ChatContainer/ChatUsers.vue"
import ChatBox from "./ChatContainer/ChatBox.vue"
export default {
    name:"ChatContainer",
    components:{ChatUsers, ChatBox}
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