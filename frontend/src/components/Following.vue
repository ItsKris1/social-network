<template>
    <div class="item-list__wrapper" id="following">
        <h3>Following</h3>
        <ul class="item-list users">
            <li v-for="user in this.following" :key="user.id">
                <div class="user-picture small"></div>
                <div class="item-text">{{ user.nickname }}</div>
            </li>
            <!--  -->
        </ul>
    </div>
    <!-- 
    <div>
        <span>Following</span>
        <div v-for="user in this.following" :key="user.id">
            {{ user.nickname }}
        </div>
    </div>


    <li>
        <div class="user-picture small"></div>
        <div class="item-text">User 2</div>
    </li>
    <li>
        <div class="user-picture small"></div>
        <div class="item-text">User 3</div>
    </li>
    <li>
        <div class="user-picture small"></div>
        <div class="item-text">User 4</div>
    </li> -->
</template>


<script>
export default {
    name: 'Following',
    data() {
        return {
            following: []
        }
    },
    created() {
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
                .then((json => {
                    console.log("following:", json)
                    return json
                }))

                .then((json => {
                    this.following = json.users
                    console.log(json.users)
                }))


        }
    }
}
</script>


<style>
</style>