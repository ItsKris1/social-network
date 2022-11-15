<template>
    <li @click="openChat" :title="user.nickname">
        <div class="user">
            <div class="user-picture"></div>
            <div class="item-text">{{ user.nickname }}</div>
        </div>
        <div class="indicator-container">
            <!-- <div class="unreadMessagesCount"></div> -->
            <div class="unreadMessagesCount" v-if="unreadMessageCount >0"></div>
            <div class="user-online"></div>
        </div>
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
    background-color: brown;
    width:10px;
    height:10px;
    border-radius: 50%;
    justify-self: flex-end;
}
.user-online {
    padding: 2.5px 5px;
    background-color: rgb(60, 148, 34);
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
