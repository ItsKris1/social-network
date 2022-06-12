<template>

    <div id="navbar">
        <div>
            <a href="/" id="logo">Social Network</a>
        </div>  
        <div>
            <input type="text" placeholder="Search">
            <img src="../assets/glass.png" alt="glass.png">
        </div>
        <router-link to="/"><img src="../assets/home.png" alt="home.png"><span>Home</span></router-link>
        <Notifications />
        <div>
            <router-link to="/profile">
                <img alt="profilePic">
                <span>James Brown</span>
            </router-link>
        </div>
        <img @click="logout" id="logoutBtn" src="../assets/logout.png">
    </div>

</template>


<script>
import Notifications from './Notifications.vue'
export default {
    name: 'NavBarOn',
    methods: {
        logout() {
            console.log("logout")
        }
    },
    components: { Notifications, }
}

</script>


<style>
#navbar {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;
}

#logoutBtn {
    width: 16px;
    height: 20px;
    cursor: pointer;
}
</style>