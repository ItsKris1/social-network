<template>

    <div id="navbar">
        <div>
            <a href="/" id="logo">Social Network</a>
        </div>
        <Search />
        <!-- <router-link to="/"><img src="../assets/home.png" alt="home.png"><span>Home</span></router-link> -->
        <Notifications />
        <div id="a">
            <router-link to="/profile">
                <img id="navbarUserPic" src="../assets/nouserimg.png" alt="profilePic">
            </router-link>
            <span id="navBarName">James Brown</span>
        </div>
        <img @click="logout" id="logoutBtn" src="../assets/logout.png">
    </div>

</template>


<script>
import Search from './Search.vue';
import Notifications from './Notifications.vue'
import Search1 from './Search.vue';
export default {
    name: 'NavBarOn',
    methods: {
        async logout() {
            // Test code
            await fetch('http://localhost:8081/logout', {
                credentials: 'include',
                headers: {
                    'Accept': 'application/json',
                }
            })
                .then((response => response.json()))
                .then((json => {console.log(json)}))
            // end of test code
            console.log("logout")
            this.$router.push("/");
        }
    },
    components: { Notifications, Search, Search1 }
}

</script>

<style scoped>
#navbar {
    display: flex;
    justify-content: space-around;
    margin-top: 10px;
    align-items: center;
}

#logo {
    font-family: Inter;
    font-size: 24px;
    font-weight: 649;
    line-height: 29px;
    letter-spacing: 0em;
    text-align: center;
    font-variation-settings: 'slnt' 0;
    text-decoration: none;
}

#logoutBtn {
    width: 16px;
    height: 20px;
    cursor: pointer;
}

#navbarUserPic {
    height: 47px;
    width: 47px;
    border-radius: 50%;
    margin-right: 8px;
}

#a {
    display: flex;
    align-items: center;
}
</style>