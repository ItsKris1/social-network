<template>
    <div id="searchDiv">
        <input @input="filtered" v-model="this.searchQuery" id="inputBox" type="text"
            placeholder="Search user or group">

        <!-- <div style="color:black">{{ this.searchQuery }}</div> -->
        <!-- <img id="glass" src="../assets/glass.png" alt="glass.png"> -->

        <ul id="dropdownlist" v-if="this.dropdownList.length !== 0">
            <li @click="goToUserProfile(user.id)" id="dropdownitem" v-for="user in dropdownList">
                <img id="dropdownimage" :src="'http://localhost:8081/' + user.avatar" alt="dropdownimage">
                {{ user.nickname }}
            </li>

        </ul>
    </div>


    <!-- <div id="dropdownlist" v-if="this.dropdownList.length !== 0">
        <div @click="goToUserProfile(user.id)" id="dropdownitem" v-for="user in dropdownList">
            <img id="dropdownimage" :src="'http://localhost:8081/' + user.avatar" alt="dropdownimage">
            {{ user.nickname }}
        </div>
    </div> -->
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
/* #searchDiv {
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px;

    width: 257px;
    height: 15px;

    background: #FFFFFF;
    border: 1px solid #706A6A;
    border-radius: var(--container-border-radius);
} */

#dropdownlist {
    cursor: default;
    margin-left: 20px;
    display: block;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
    padding: 12px 16px;
    z-index: 1;
}

#dropdownitem {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 400;
    font-size: 16px;
    line-height: 146.02%;

    display: flex;
    align-items: center;
    text-align: center;
    color: #141111;
}

#dropdownimage {
    width: 20px;
    height: 20px;
}

#inputBox {
    background-image: url(../assets/glass.svg);
    background-repeat: no-repeat;
    background-position: right 7.5px center;
    padding-right: calc(17px + 7.5px + 7.5px);
}

*:focus {
    outline: none;
}

#glass {
    margin-left: auto;
}
</style>