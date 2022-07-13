<template>
    <div>
        <span><b>Followers</b></span>
        <div v-for="user in this.followers" :key="user.id">
            {{ user.nickname }}
        </div>
    </div>
</template>


<script>
export default {
    name: 'Followers',
    data() {
        return {
            followers: []
        }
    },
    created() {
        this.getFollowers()
    },
    watch: { //watching changes in route
        $route() {
            this.getFollowers()
        }
    },
    methods: {
        async getFollowers() {
            // console.log("getFollowers");
            await fetch('http://localhost:8081/followers?userId=' + this.$route.params.id, {
                credentials: 'include'
            })
                .then((response => response.json()))
                .then((json => {
                    // console.log("Followers:", json)
                    this.followers = json.users
                    }))
        }
    }
}
</script>


<style>
</style>