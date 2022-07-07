<template>
    <div>
        <span>Following</span>
        <div>

        </div>
    </div>
</template>


<script>
export default {
    name: 'Following',
    methods: {
        async getFollowing() {
            // console.log("getFollowing");
            await fetch('http://localhost:8081/following?userId=' + this.$route.params.id, {
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