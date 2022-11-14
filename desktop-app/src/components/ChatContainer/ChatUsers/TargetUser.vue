<template>
    <li class="user"  @click="openChat">
        <div class="user-picture"></div>
        <div class="item-text">{{ user.nickname }}</div>
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
    methods:{
        openChat(){
            this.$store.commit("openNewChat",{id:this.user.id, type: this.type, name:this.user.nickname})
            this.$store.dispatch('fecthChatMessages')
        }
    }
}
</script>
<style scoped>

.user {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    justify-content: flex-start;
    cursor: pointer;
}
</style>
