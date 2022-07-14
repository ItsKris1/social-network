<template>
    <button class="btn" @click="follow">Follow<i class="uil uil-user-plus"></i></button>
</template>


<script>
export default {
    name: 'FollowBtn',
    props: ['profileId'],
    data() {
        return {
            userid: "",
            isMyProfile: false,
        }
    },

    created() {
        this.checkProfile();
    },


    watch: { //watching changes in route
        $route() {
            this.checkProfile()
        }
    },
    methods: {
        async follow() {
            // console.log('subscribe function:')
            await fetch("http://localhost:8081/follow?userId=" + this.$route.params.id, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json => console.log("server response:", json)))
        },
        unfollow() {
            console.log('unsubscribe function')
        },
        async getLoggedUserId() {
            await fetch("http://localhost:8081/currentUser", {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json => {
                    this.userid = json.users[0].id
                }))
        },

        async checkProfile() {
            await this.getLoggedUserId();
            this.isMyProfile = this.userid === this.profileId
        },
    }
}
</script>


<style>
#followBtn {
    height: 37px;
    width: 92px;
    left: 0px;
    top: 0px;
    border-radius: 8px;
    padding: 10px, 15px, 10px, 15px;
}
</style>