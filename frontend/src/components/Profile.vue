<template>
    <div id="profile">
        <div>
            <img id="profileImg" src="../assets/nouserimg.png" alt="profilePic">
        </div>
        <div id="basicInformation">
            <span>James Brown</span>
            <span>e-mail</span>
            <span>birthday date</span>
            <button id="followBtn" click="follow">Follow</button>
        </div>
        <div id="profileSettings">
            <span>Privacy change-button</span>
        </div>
    </div>
    <AllMyPosts />
</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
export default {
    name: 'Profile',
    components:{AllMyPosts},
    methods: {
        follow() {
            console.log('subscribe function')
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
}

#basicInformation {
    display: grid;
}

#profileSettings{
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