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
                <img id="navbarUserPic" :src="'http://localhost:8081/' + user.avatar" alt="profilePic">
            </router-link>
            <span id="navBarName">{{ user.nickname }}</span>
        </div>
        <img @click="logout" id="logoutBtn" src="../assets/logout.png">
    </div>

</template>


<script>
import Search from './Search.vue';
import Notifications from './Notifications.vue'
export default {
    name: 'NavBarOn',
    data() {
        return {
            user: {}
        }
    },
    created() {
        this.getUserInfo()
    },
    methods: {
        async getUserInfo() {
            await fetch("http://localhost:8081/currentUser", {
                credentials: 'include',
            })
                .then((r => r.json()))
                .then((json => {
                    console.log(json)
                    this.user = json.users[0]
                }))

        },
        async logout() {
            // Test code
            await fetch('http://localhost:8081/logout', {
                credentials: 'include',
                headers: {
                    'Accept': 'application/json',
                }
            })
                .then((response => response.json()))
                .then((json => { console.log(json) }))
            // end of test code
            console.log("logout")
            this.$router.push("/");
        }
    },
    components: { Notifications, Search, }
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