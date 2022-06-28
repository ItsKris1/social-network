<template>

    <div id="navbar">
        <div>
            <a href="/" id="logo">Social Network</a>
        </div>
        <Search />
        <Notifications />
        <div id="a">
            <router-link to="/profile">
                <!-- <img id="navbarUserPic" :src="'http://localhost:8081/' + user.avatar" alt="profilePic"> -->

                <!-- <span id="navBarName">{{ user.nickname }}</span> -->
                <span id="navBarName">My profile</span>

            </router-link>
        </div>
        <!-- <img @click="logout" id="logoutBtn" src="../assets/logout.png"> -->
        <span @click="logout" id="logoutBtn">Log out</span>
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
                    // console.log(json)
                    this.user = json.users[0]
                }))

        },
        async logout() {
            await fetch('http://localhost:8081/logout', {
                credentials: 'include',
                headers: {
                    'Accept': 'application/json',
                }
            })
                .then((response => response.json()))
                .then((json => { console.log(json) }))
            console.log("logout")
            this.$router.push("/");
        }
    },
    components: { Notifications, Search, }
}

</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@100;200&family=Water+Brush&display=swap');

#navbar {
    display: flex;
    justify-content: space-around;
    margin-top: 10px;
    align-items: center;

    background: #4B5283;
    font-family: 'Poppins';
    color: #EEEEEE;
    
}

#logo {
    /* font-family: Inter; */
    
    font-size: 24px;
    font-weight: 649;
    line-height: 29px;
    letter-spacing: 0em;
    text-align: center;
    font-variation-settings: 'slnt' 0;
    text-decoration: none;
    color: #EEEEEE;
}

#logoutBtn {
    /* width: 16px;
    height: 20px; */
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

#navBarName {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 400;
    font-size: 16px;
    line-height: 146.02%;
    /* or 23px */

    display: flex;
    align-items: center;
    text-align: center;
    text-decoration: none;

    color: #E4E3E3;
}
a:link { text-decoration: none; }
a:visited { text-decoration: none; }
a:hover { text-decoration: none; }
a:active { text-decoration: none; }
</style>