<template>
    <div id="profile">
        <!-- {{ user }} -->
        <div>
            <img id="profileImg" :src="'http://localhost:8081/' + user.avatar" alt="profilePic">
        </div>
        <div id="basicInformation">
            <span>{{ user.nickname }}</span>
            <span>{{ user.login }}</span>
            <span>{{ user.dateOfBirth }}</span>
            <button id="followBtn" click="follow">Follow</button>
        </div>
        <div id="profileSettings">
            <span>Privacy change-button</span>
        </div>
    </div>
    <div v-if="user.about !== ''">
        <span>ABOUT ME<br></span>
        <span>{{ user.about }}</span>
    </div>
    <Following />
    <Followers />
    <AllMyPosts v-bind:userid="this.user.id" />
</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
import Following from './Following.vue'
import Followers from './Followers.vue'
// import { mapGetters } from 'vuex'
export default {
    name: 'Profile',
    components: { AllMyPosts, Followers, Following },
    data() {
        return {
            user: {}
        }
    },
    created() {
        // this.getUserInfo()
        this.getUserId()
    },
    computed: {
        // ...mapGetters(['userInfo']),
        // ...getUserId()
    },
    methods: {
        follow() {
            console.log('subscribe function')
        },
        getUserInfo() {
            this.$store.dispatch('getMyProfileInfo')
        },
        async getUserId() {
            await fetch("http://localhost:8081/userData?userId=" + this.$route.params.id, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json) => {
                    // console.log("profile.vue/getuserid",json);
                    this.user = json.users[0];
                    // console.log(userInfo);
                    // this.commit("updateProfileInfo", userInfo);
                    // console.log("user profile info -", json);
                });
            // this.isOwnerProfile()
        },
        // isOwnerProfile() {
        //     console.log("cookie",document.cookie);
        //     console.log("user id",this.user.id);
        //     let activeCookie = document.cookie.slice(11)
        //     if (activeCookie === this.user.id) {
        //         console.log("It's a owner");
        //     } else { console.log("It's NOT a owner") }

        // }
    },
    watch: { //watching changes in route
        $route() {
            this.getUserId()
        }
    }
}
</script>

<style>
#profile {
    display: flex;
    margin-top: 10px;
}

#profileImg {
    border-radius: 50%;
    margin-right: 20px;
    width: 200px;
    height: 200px;
}

#basicInformation {
    display: grid;
}

#profileSettings {
    margin-left: auto;
}

#followBtn {
    height: 37px;
    width: 92px;
    left: 0px;
    top: 0px;
    border-radius: 8px;
    padding: 10px, 15px, 10px, 15px;
}
</style>