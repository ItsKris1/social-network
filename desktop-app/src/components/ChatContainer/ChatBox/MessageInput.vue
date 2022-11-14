<template>
    <form @submit.prevent="sendMessage" autocomplete="off" class="send-message">
        <input type="text" name="sent-message" id="sent-message__input" placeholder="Send a message" v-model="message">
        <button type="submit"><i class="uil uil-message"></i></button>
            <EmojiSelector @addEmoji="addEmoji"/>
        </form>
</template>

<script>
import { mapState } from 'vuex';
import EmojiSelector from './EmojiSelector.vue';

export default {
    components: { EmojiSelector },
   data() {
        return {
            message:""
        }
    },
    computed:{
        ...mapState({
            myID: state => state.id
        })
    },
    methods:{
        addEmoji(emoji){
            this.message += emoji
        },
        async sendMessage(){
            if (this.message === "") {return}
            let result = this.$store.dispatch('sendMessage', this.message)
            if (result){
                this.message = "";
            }else{
                console.log("Err: message not sent")
            }
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