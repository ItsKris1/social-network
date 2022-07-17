<template>
<div class="temporary">
    <div>{{this.groupData.name}}</div>
    <div>{{this.groupData.description}}</div>
</div>
    
</template>


<script>
export default {
    name: 'Group',
    created(){
        this.getGroupInfo()
    },
    watch: {
        $route() {
            this.getGroupInfo()
        }
    },
    data(){
        return{
            groupData:null
        }
    },
    methods: {
        async getGroupInfo() {
            await fetch("http://localhost:8081/groupInfo?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
            .then((r=>r.json()))
            .then((json=>{
                console.log(json)
                this.groupData = json.groups[0]
                }))

        },
    },
}
</script>


<style>
.temporary{
    margin-top: 100px;
}
</style>