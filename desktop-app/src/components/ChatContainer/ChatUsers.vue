<template>
    <div class="user-container">
        <!-- header -->
        <div class="chat-header">
            <h2>Chats</h2>
        </div>
        <!-- list -->
        <ul class="item-list">
            <TargetUser v-for="user in chatUserList" :key="user.id" :user="user"/>
        </ul>

        <ul class="item-list">
            <TargetGroup v-for="group in userGroups" :key="group.id" :group="group"/>
        </ul>

        <p v-if="chatUserList === null && userGroups === null">
        No one to message with :(</p>
    </div>
    
</template>

<script>
import TargetUser from './ChatUsers/TargetUser.vue'
import TargetGroup from './ChatUsers/TargetGroup.vue'

import { mapState } from 'vuex';

export default{
    name:"ChatUsers",
    components:{TargetUser, TargetGroup},

    async created(){
        this.$store.dispatch("fetchChatUserList")
        .then(()=>  this.$store.dispatch("getUserGroups"))
        .then(()=> this.$store.dispatch("fecthChatStack"))
        
    },

     computed:{
         ...mapState({
            userGroups: state => state.groups.userGroups,
            chatUserList: state => state.chat.chatUserList,
         })
     }
}
</script>

<style>
.chat-header{
    margin-bottom:15px;
}
.user-container{
    height:100vh;
    width:500px;
    background-color:var(--page-bg);
    padding: 15px 15px 15px 15px;
}
  .item-list img, .item-list .user-picture{
   height: 1.5em;
    width: 1.5em;
}
.item-list:nth-child(3){
    margin-top:10px;
}
.item-list li {
    justify-content: space-between;
}
@media screen and (max-width: 700px) {
  .user-container {
    width: 75px;
  }
  .chat-header{
    visibility:hidden;
  }
  .item-list .item-text{
    display:none;
    /* visibility:hidden; */
  }
  .item-list img, .item-list .user-picture{
   height: 3em;
    width: 3em;
    
    }
    .item-list li{
        display:block;
        position:relative;
    }

}
</style>
