<template>

    <div class="item-list__wrapper" id="followers">
        <h3>Followers</h3>
        <ul class="item-list users" v-if="this.followers">
            <li v-for="user in this.followers" :key="user.id">
                <div class="user-picture small"></div>
                <div class="item-text">{{ user.nickname }}</div>
            </li>
        </ul>

        <p class="additional-info" v-else>No followers</p>
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
                    console.log("Followers:", json)
                    this.followers = json.users
                }))
        }
    }
}
</script>


<style>
</style>