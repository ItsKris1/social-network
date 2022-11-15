<template>
    <li  @click="openChat" :title="group.name">
        <div class="group">
            <img src="../../../assets/icons/users-alt.svg" alt=""/>
            <div class="item-text">{{ group.name }}</div>
        </div>
        <div class="indicator-container">
        <p class="unreadMessagesCount" v-if="unreadMessageCount >0"></p>
    </div>
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
    cursor: pointer;
}
.unreadMessagesCount {
    padding: 2.5px 5px;
    background-color: brown;
    width:10px;
    height:10px;
    border-radius: 50%;
    justify-self: flex-end;
}
.indicator-container{
    display:flex;
    gap:5px;
}

@media screen and (max-width: 700px){
    .indicator-container{
        position:absolute;
        flex-direction: column;
        right:-15px;
        bottom:0;
    }
}
</style>