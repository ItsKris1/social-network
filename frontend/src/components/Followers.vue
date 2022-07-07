<template>
    <div>
        <span>Followers</span>
        <div>

        </div>
    </div>
</template>


<script>
export default {
    name: 'Followers',
    created() {
        this.getFollowers()
    },
    methods: {
        async getFollowers() {
            // console.log("getFollowers");
            await fetch('http://localhost:8081/followers?userId=' + this.$route.params.id, {
                credentials: 'include'
            })
                .then((response => response.json()))
                // .then((json=>console.log(json)))
        }
    }
}
</script>


<style>
</style>