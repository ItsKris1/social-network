<template>
    <div id="profile">
        {{ userInfo }}
        <div>
            <img id="profileImg" :src="'http://localhost:8081/' + userInfo.avatar" alt="profilePic">
        </div>
        <div id="basicInformation">
            <span>{{ userInfo.nickname }}</span>
            <span>{{ userInfo.login }}</span>
            <span>{{ userInfo.dateOfBirth }}</span>
            <button id="followBtn" click="follow">Follow</button>
        </div>
        <div id="profileSettings">
            <span>Privacy change-button</span>
        </div>
    </div>
    <div v-if="userInfo.about !== ''">
        <span>ABOUT ME<br></span>
        <span>{{ userInfo.about }}</span>
    </div>
    <AllMyPosts />
</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
import { mapGetters } from 'vuex'
export default {
    name: 'Profile',
    components: { AllMyPosts },
    created() {
        this.getUserInfo()
    },
    computed: mapGetters(['userInfo']),
    methods: {
        follow() {
            console.log('subscribe function')
        },
        getUserInfo() {
            this.$store.dispatch('getMyProfileInfo')
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