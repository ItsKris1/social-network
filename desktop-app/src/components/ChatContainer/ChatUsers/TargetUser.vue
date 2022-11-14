<template>
    <li @click="openChat" :title="user.nickname">
        <div class="user">
            <div class="user-picture"></div>
            <div class="item-text">{{ user.nickname }}</div>
        </div>
        <p class="unreadMessagesCount" v-if="unreadMessageCount >0">{{unreadMessageCount}}</p>
    </li>
</template>

<script>

export default {
    name:"TargetUser",
    props: ['user'],
    data() {
        return {
            type:"PERSON"
        }
    },
    computed:{
        unreadMessageCount(){
            let unreadMessages = this.$store.state.chat.unreadMessages
            let count = unreadMessages[this.user.id]? unreadMessages[this.user.id]:0;
            return count
        }
    },
    methods:{
        openChat(){
            this.$store.dispatch('removeUnreadMessages', this.user.id)
            this.$store.commit("openNewChat",{id:this.user.id, type: this.type, name:this.user.nickname})
            this.$store.dispatch('fecthChatMessages')
        },
        
    }
}
</script>
<style scoped>

.user {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    cursor: pointer;
}
.unreadMessagesCount {
    padding: 2.5px 5px;
    color: var(--color-white);
    background-color: brown;
    font-size: 12px;
    border-radius: 5px;
    justify-self: flex-end;
}
</style>
