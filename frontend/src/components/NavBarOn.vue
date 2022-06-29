<template>

    <div id="navbar">
        <div class="nav-titleSearch">
            <a id="nav-title" href="/">Social Network</a>
            <Search />
        </div>

        <ul class="nav-links">
            <li>
                <Notifications />
            </li>
            <li>
                <router-link to="/profile">
                    My profile
                </router-link>
            </li>
            <li @click="logout">Log out</li>
        </ul>


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
    padding: 10px 40px 10px 60px;
    background-color: var(--color-blue);
    align-items: center;
    justify-content: space-between;
    font-weight: 300;
    color: var(--color-white);
    margin-bottom: 50px;

}

#navbar a {
    color: var(--color-white);
}

#nav-title {
    font-size: 24px;
}


.nav-links>li {
    display: inline-block;
    margin-left: 20px;
}


.nav-titleSearch {
    display: flex;
    gap: 25px
}



a:link {
    text-decoration: none;
}

a:visited {
    text-decoration: none;
}

a:hover {
    text-decoration: none;
}

a:active {
    text-decoration: none;
}
</style>