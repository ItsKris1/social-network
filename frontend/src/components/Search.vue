<template>
    <div id="searchDiv" @click.stop>
        <input @focus="toggleDropdown" v-model="searchQuery"
               :class="{ 'no-bottom-border': showDropdown }" type="text" placeholder="Search user or group">

        <div class="shadow-wrapper" v-show="showDropdown">
            <div id="dropdown">
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

    </div>

</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: 'Search',
    data() {
        return {
            filteredUsers: [],
            filteredGroups: [],
            showDropdown: false,
            searchQuery: "i"
        }
    },
    created() {
        this.$store.dispatch('getAllUsers')
        this.$store.dispatch('getAllGroups')
        window.addEventListener("click", this.hideDropdown)
    },

    watch: {
        searchQuery() {
            this.filteredUsers = this.filterUsers(this.searchQuery)
            this.filteredGroups = this.filterGroups(this.searchQuery)
            this.toggleDropdown();
        }
    },

    computed: {
        ...mapGetters(['allUsers', 'allGroups', 'filterUsers', 'filterGroups'])
    },
    methods: {
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

        // goToGroupPage(groupId) {
        //     this.$router.push({ name: 'Group', params: { id: groupId } })
        // },

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
    border-radius: 10px;
}

#searchDiv input[type="text"] {
    background-image: url(../assets/icons/glass.svg);
    background-repeat: no-repeat;
    background-position: left 10px center;
    border-radius: 10px;
    padding: 10px;
    padding-left: calc(17px + 20px);
    box-shadow: var(--container-shadow);
    background-color: var(--input-bg);
    cursor: pointer;

    transition: box-shadow 0.25s ease-in;

}



#searchDiv input[type="text"].no-bottom-border {
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
}

.shadow-wrapper {
    position: absolute;
    width: calc(100% + 8px);
    transform: translateX(50%);
    right: 50%;
    overflow: hidden;
    padding: 4px;
}

#dropdown {
    position: relative;
    background-color: var(--input-bg);
    color: var(--color-lg-black);
    margin-top: -4px;
    box-shadow: 0 0 5px 1px var(--hover-color);
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;
    width: 100%;

}


#dropdown .item-list {
    border-top: 1px solid rgb(211, 211, 211);
    padding: 15px;
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


#searchDiv input[type="text"]:hover,
#searchDiv input[type="text"]:focus {
    /* position: relative; */
    /* z-index: 2; */
    box-shadow: 0 0 4px 2px var(--hover-color);
}
</style>