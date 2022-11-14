<template>
    <form autocomplete="off" class="send-message">
        <input type="text" name="sent-message" id="sent-message__input" placeholder="Send a message" v-model="message">
        <button type="submit"><i class="uil uil-message"></i></button>
            <EmojiSelector @addEmoji="addEmoji"/>
        </form>
</template>

<script>
import EmojiSelector from './EmojiSelector.vue';
export default {
    components: { EmojiSelector },
   data() {
        return {
            message:""
        }
    },
    methods:{
        addEmoji(emoji){
            this.message += emoji
        }
    }
}
</script>


<style scoped>
.send-message {
    background-color: var(--color-white);
    padding: 10px;
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
    align-items: center;
}
.send-message input {
    padding: 12px 20px;
    border-radius: 30px;
    flex: 1;
}
.send-message button {
    border: none;
    background-color: inherit;
    font-size: 1.25em;
}
.send-message button:hover {
    color: var(--hover-color);
}
</style>