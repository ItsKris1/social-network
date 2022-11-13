<template>
    <div class="chat-box-container">
        <h2 class="chat-header">Some user name</h2>
        <div class="chat-content">content</div>
        <MessageInput/>
    </div>
</template>

<script>
import MessageInput from './ChatBox/MessageInput.vue';
export default{
    components:{MessageInput}

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
</style>