<template>
    <div id="searchDiv" @click.stop>
        <input @input="filterSearch" @focus="toggleDropdown" v-model="searchQuery"
               :class="{ 'no-bottom-border': showDropdown }" type="text" placeholder="Search user or group">

        <div id="dropdown" v-show="showDropdown">
            <ul class="item-list">
                <li @click="goToUserProfile(user.id)" id="dropdownitem" v-for="user in filteredUsers">
                    <div class="user-picture small"
                         :style="{ backgroundImage: `url(http://localhost:8081/${user.avatar})` }"></div>
                    <div class="item-text">{{ user.nickname }}</div>
                </li>

                <li @click="goToGroupPage(group.id)" id="dropdownitem" v-for="group in filteredGroups">
                    <img src="../assets/icons/users-alt.svg" alt="" class="small">
                    <div class="item-text">{{ group.name }}</div>
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
            searchQuery: "",
            filteredUsers: [],
            filteredGroups: [],

            showDropdown: false,
        }
    },
    created() {
        this.$store.dispatch('getAllUsers')
        this.$store.dispatch('getAllGroups')
        window.addEventListener("click", this.hideDropdown)
    },


    computed: mapGetters(['allUsers', 'allGroups', 'filterUsers', 'filterGroups']),
    methods: {

        filterSearch() {
            this.filteredUsers = this.filterUsers(this.searchQuery)
            this.filteredGroups = this.filterGroups(this.searchQuery)
            this.toggleDropdown();
        },

        goToUserProfile(userid) {
            this.$router.push({ name: 'Profile', params: { id: userid } })
            this.clearSearch();
            this.hideDropdown();

        },

        goToGroupPage(groupId) {
            this.$router.push({ name: 'Group', params: { id: groupId } })
            this.clearSearch();
            this.hideDropdown();
        },

        toggleDropdown() {
            this.filteredUsers.length > 0 || this.filteredGroups.length > 0 ? this.showDropdown = true : this.showDropdown = false

        },

        clearSearch() {
            this.searchQuery = "";
        },

        hideDropdown() {
            this.showDropdown = false;
        },
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

#searchDiv input[type="text"].no-bottom-border {
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
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