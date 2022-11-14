<template>
    <li  @click="openChat" :title="group.name">
        <div class="group">
            <img src="../../../assets/icons/users-alt.svg" alt=""/>
            <div class="item-text">{{ group.name }}</div>
        </div>
        <p class="unreadMessagesCount" v-if="unreadMessageCount >0">{{unreadMessageCount}}</p>
    </li>
</template>

<script>

export default {
    name:"TargetGroup",
    props: ['group'],
     data() {
        return {
            type:"GROUP"
        }
    },
    computed:{
        unreadMessageCount(){
            let unreadMessages = this.$store.state.chat.unreadMessages
            let count = unreadMessages[this.group.id]? unreadMessages[this.group.id]:0;
            return count
        }
    },
    methods:{
        openChat(){
            this.$store.dispatch('removeUnreadMessages', this.group.id)
            this.$store.commit("openNewChat",{id:this.group.id, type: this.type, name:this.group.name})
            this.$store.dispatch('fecthChatMessages')

        }
    }
}
</script>
<style scoped>
.group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
}
.unreadMessagesCount {
    padding: 2.5px 5px;
    color: var(--color-white);
    background-color: brown;
    font-size: 12px;
    border-radius: 5px;
}
</style>