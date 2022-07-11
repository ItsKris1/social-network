<template>
    <div>
        <span>Following</span>
        <div v-for="user in this.following" :key="user.id">           
                {{user.nickname}}
        </div>
    </div>
</template>


<script>
export default {
    name: 'Following',
    data(){
        return {
            following:[]
        }
    },
    created(){
        this.getFollowing()
    },
    watch: { //watching changes in route
        $route() {
            this.getFollowing()
        }
    },
    methods: {
        async getFollowing() {
            // console.log("getFollowing");
            await fetch('http://localhost:8081/following?userId=' + this.$route.params.id, {
                credentials: 'include'
            })
                .then((response => response.json()))
                .then((json=>{
                    console.log("following:",json) 
                    return json
                    }))
                
                .then((json=>{
                    this.following = json.users
                }))


        }
    }
}
</script>


<style>
</style>