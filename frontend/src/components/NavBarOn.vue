<template>

    <div id="navbar">

        <div id="nav-titleSearch">
            <a id="nav-title" href="/">Social Network</a>
            <Search />
        </div>



        <ul class="nav-links">

            <li>
                <Notifications />
            </li>
            <li>
                <router-link v-if="typeof user.id !== 'undefined'" :to="{ name: 'Profile', params: { id: user.id } }">
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
#navbar {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 555;
    width: 100%;
    min-width: min-content;

    display: flex;
    align-items: center;
    justify-content: space-between;

    padding: 10px 40px;
    background-color: var(--color-blue);
    color: var(--color-white);


}


#navbar a {
    color: var(--color-white);
}

#nav-title {
    font-size: 24px;
    font-weight: 400;
}


.nav-links>li {
    font-weight: 300;
    display: inline-block;
    margin-left: 20px;
}


#nav-titleSearch {
    display: flex;
    gap: 25px;
    flex-grow: 1;
    align-items: center;


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