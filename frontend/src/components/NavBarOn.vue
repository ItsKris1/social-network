<template>

    <div id="navbar">

        <div id="nav-titleSearch">
            <!-- <a id="nav-title" href="/">Social Network</a> -->
            <router-link to="/" id="nav-title">Social Network</router-link>
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
    z-index: 1;
    width: 100%;
    min-width: min-content;

    display: flex;
    align-items: center;
    justify-content: space-between;

    padding: 10px 40px;
    background-color: var(--color-blue);
    color: var(--color-white);

    position: relative;




}


#navbar a {
    color: var(--color-white);
}





#nav-title {
    font-size: 24px;
    font-weight: 400;
}


.nav-links li {
    font-weight: 300;
    display: inline-block;
    margin-left: 20px;
    cursor: pointer;
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

/* .nav-links li:hover {
    color: rgb(241, 181, 181)
} */


#navbar li:hover {
    /* color: rgb(255, 255, 255); */
    /* background-color: rgb(95, 107, 175);
    padding: 5px 7.5px;
    border-radius: 10px; */

    /* text-decoration: underline;
    text-decoration-color: rgb(132, 148, 236);
    text-decoration-thickness: 2px; */

    color: red;
    line-height: 0;
    border-bottom: 2px solid rgb(132, 148, 236);
}

#nav-title {
    padding: 2px;
}

#nav-title:hover {
    /* color: rgb(95, 107, 175); */
    /* text-shadow: 2px 2px 5px rgb(95, 107, 175); */
    /* text-transform: uppercase; */
    /* color: white; */

    /* text-decoration: underline;
    text-decoration-color: rgb(132, 148, 236);
    text-decoration-thickness: 2px;
    text-underline-offset: 5px; */
    border-bottom: 2px solid rgb(132, 148, 236);
    padding-bottom: 0px;
    /* padding: 3px; */
    /* line-height: 1; */
    /* padding: 0px 7.5px; */
    /* border-radius: 10px; */

}
</style>