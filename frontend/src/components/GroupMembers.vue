<template>
    <div class="item-list__wrapper" id="groups">
        <h3>Members:</h3>
        <ul class="item-list">
            <li v-for="member in this.groupMembers">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">{{ member.nickname }}</div>
            </li>
        </ul>
    </div>
</template>


<script>
export default {
    name: 'GroupMembers',
    data(){
        return{
            groupMembers:null,
        }
    },
    created(){
        this.getGroupMembers()
    },
    watch: {
        $route() {
            this.getGroupMembers();
        }
    },
    methods: {
        async getGroupMembers() {
            await fetch('http://localhost:8081/groupMembers?groupId=' + this.$route.params.id, {
                credentials: 'include'
            })
                .then((response => response.json()))
                .then((json => {
                    console.log("GroupMembers:", json)
                    this.groupMembers = json.users
                }))
        }
    }
}
</script>


<style>
</style>