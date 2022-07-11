<template>
    <button class="btn" v-show="this.isMyProfile" @click="follow">Follow<i class="uil uil-user-plus"></i></button>
</template>


<script>
export default {
    name: 'FollowBtn',
    data() {
        return {
            userid: "",
            isMyProfile: false,
        }
    },
    props: {
        profileId: ""
    },
    created() {
        this.getLoggedUserId()
        this.checkProfile()
    },
    watch: { //watching changes in route
        $route() {
            this.checkProfile()
        }
    },
    methods: {
        async follow() {
            // console.log('subscribe function:')
            // console.log(this.profileId);
             await fetch("http://localhost:8081/follow?userId=" + this.profileId, {
                credentials: "include",
            })
                .then((r) => r.json())
                // .then((json=>console.log("server response:",json)))
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
        checkProfile() {

            if (this.userid === this.profileId) {
                this.isMyProfile = true
            } else {
                this.isMyProfile = false
            }
            // console.log("checkProfile", this.isMyProfile);
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