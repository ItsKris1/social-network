<template>

    <div id="navbar">
        <div>
            <a href="/" id="logo">Social Network</a>
        </div>
        <div id="searchDiv">
            <input id="inputBox" type="text" placeholder="Search for user or group">
            <img id="glass" src="../assets/glass.png" alt="glass.png">
        </div>
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
import Notifications from './Notifications.vue'
export default {
    name: 'NavBarOn',
    methods: {
        logout() {
            console.log("logout")
            this.$router.push("/");
        }
    },
    components: { Notifications, }
}

</script>

<style scoped>
#navbar {
    display: flex;
    justify-content: space-around;
    margin-top: 10px;
    align-items: center;
}

/* #searchDiv{
margin-right: auto;
} */
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

/* #navBarName {
    font-family: Inter;
    font-size: 16px;
    font-weight: 400;
    line-height: 19px;
    letter-spacing: 0em;
    text-align: center;
} */
#a {
    display: flex;
    align-items: center;
}

#searchDiv {
    display: flex;
    flex-direction: row;
    /* justify-content: space-between; */
    align-items: center;
    padding: 10px;
    /* gap: 53px; */

    width: 257px;
    height: 15px;

    background: #FFFFFF;
    border: 1px solid #706A6A;
    border-radius: 3px;
    border: 1px solid #706A6A;
}

#inputBox {
    border: none;
    width: 230px;
}
*:focus {
    outline: none;
}
#glass{
    margin-left: auto;
}
</style>