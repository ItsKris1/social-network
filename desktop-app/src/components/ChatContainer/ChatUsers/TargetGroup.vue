<template>
    <li class="group" @click="openChat">
        <img src="../../../assets/icons/users-alt.svg" alt=""/>
        <div class="item-text">{{ group.name }}</div>
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
    methods:{
        openChat(){
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
    justify-content: flex-start;
}
</style>