<template>
    <div class="chat-container">
        <SideBar :auth="true"/>
        <ChatContainer />
    </div>
</template>

<script>
import ChatContainer from "@/components/ChatContainer.vue"
import SideBar from '@/components/SideBar.vue'

export default {
    name: 'ChatView',
    components:{
        ChatContainer,
        SideBar
    }
}
</script>

<style>
.chat-container{
  display:flex;
  width:100%;
  height:100%;
}
</style>