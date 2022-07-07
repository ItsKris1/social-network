<template>

    <!-- 
    <div id="searchDiv">

        <input @input="filtered" v-model="this.searchQuery" id="inputBox" type="text"
            placeholder="Search user or group">



        <div style="color:black">{{ this.searchQuery }}</div>
        <img id="glass" src="../assets/glass.png" alt="glass.png">

        <ul id="dropdownlist" v-if="this.dropdownList.length !== 0" class="item-list">
            <li @click="goToUserProfile(user.id)" id="dropdownitem" v-for="user in dropdownList">
                <div class="user-picture small"
                    :style="{ backgroundImage: `url(http://localhost:8081/${user.avatar})` }">
                </div>
                {{ user.nickname }}


            </li>

        </ul>
    </div> -->


    <div id="searchDiv">

        <input @input="filtered" v-model="this.searchQuery" type="text" placeholder="Search user or group">

        <div id="dropdown" v-if="this.dropdownList.length !== 0">
            <ul class="item-list">
                <li @click="goToUserProfile(user.id)" id="dropdownitem" v-for="user in dropdownList">
                    <div class="user-picture small"
                        :style="{ backgroundImage: `url(http://localhost:8081/${user.avatar})` }"></div>
                    <div class="item-text">{{ user.nickname }}</div>
                </li>
                <li>
                    <div class="user-picture small"></div>
                    <div class="item-text">John Mayer</div>
                </li>

                <li>
                    <img class="small" src="../assets/icons/users-alt.svg" alt="">
                    <div class="item-text">Garrisonsttttttttttttttttttttttttttttttttttttttttttt</div>
                </li>
            </ul>
        </div>
    </div>

</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: 'Search',
    data() {
        return {
            allusers: [],
            searchQuery: "",
            dropdownList: [],
        }
    },
    created() {
        this.$store.dispatch('getAllUsers')
    },
    computed: mapGetters(['allUsers']),
    methods: {
        filtered() {
            this.dropdownList = this.$store.getters.filterUsers(this.searchQuery)
        }, goToUserProfile(userid) {
            this.$router.push({ name: 'Profile', params: { id: userid } })
        }
    },
}
</script>


<style>
#searchDiv {
    flex-grow: 1;
    position: relative;
    align-self: flex-start;
    min-width: 250px;
    max-width: 250px;
}

#searchDiv input[type="text"] {
    background-image: url(../assets/icons/glass.svg);
    background-repeat: no-repeat;
    background-position: left 10px center;
    padding: 10px;
    padding-left: calc(17px + 20px);
    border-radius: 10px;
    box-shadow: var(--container-shadow);
}


#dropdown {

    background-color: var(--input-bg);
    color: var(--color-lg-black);
    width: 100%;
    padding: 0 10px;
    position: absolute;
    box-shadow: 0 2px 5px -1px rgba(0, 0, 0, 0.27);

    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;

    /* display: none; */

}

#dropdown .item-list {
    padding: 15px 0;
    border-top: 1px solid rgb(211, 211, 211);
}

.open #inputBox {
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
}

.open #dropdownlist {
    display: block;
}


*:focus {
    outline: none;
}
</style>