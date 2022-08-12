<template>
    <div class="item-list__wrapper" id="groups">
        <h3>Members</h3>
        <ul class="item-list">
            <li v-for="member in this.groupMembers">
                <img class="small" src="../assets/icons/default-profile.svg" alt="">
                <div class="item-text">{{ member.nickname }}</div>
            </li>
        </ul>
        <button class="btn form-submit" @click="toggleModal">Invite users<i class="uil uil-user-plus"></i></button>

        <Modal v-if="this.isOpen" @closeModal="toggleModal">
            <template #title>Invite users</template>
            <template #body>

                <MultiselectDropdown v-model:checkedOptions="checkedNames" placeholder="Select followers"
                                     :content="allFollowersNames" :clearInput="clearInput"
                                     @inputCleared="toggleClearInput" />
                <button class="btn form-submit" @click="toggleModal() ; inviteUsersToGroup()">Invite</button>
            </template>
        </Modal>
    </div>
</template>


<script>
import Modal from './Modal.vue';
import MultiselectDropdown from './MultiselectDropdown.vue';
export default {
    name: "GroupMembers",
    data() {
        return {
            groupMembers: null,
            isOpen: false,
            followers: [],
            listForShowing: [],
            allUsers: [],
            checkedNames: [],
            clearInput: false,
        };
    },
    created() {
        this.getGroupMembers();
        this.getFollowers();
    },

    computed: {
        allFollowersNames() {
            return this.listForShowing.map(user => user.nickname)
        }
    },

    watch: {
        $route() {
            this.getGroupMembers();
            this.getFollowers();
        }
    },
    methods: {

        async getFollowers() {
            let id
            await fetch('http://localhost:8081/currentUser', {
                credentials: "include"
            })
                .then((response => response.json()))
                .then((json => {
                    // console.log("CurrentUser: ", json);
                    id = json.users[0].id
                }));


            await fetch('http://localhost:8081/followers?userId=' + id, {
                credentials: "include"
            })
                .then((response => response.json()))
                .then((json => {
                    // console.log("Followers:", json);
                    this.followers = json.users;

                }))
                .then(() => {
                    if (this.followers !== null) {
                        this.createFollowersListForShowing(this.followers, this.groupMembers)
                    }
                })

        },
        async getGroupMembers() {
            await fetch("http://localhost:8081/groupMembers?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((response => response.json()))
                .then((json => {
                    // console.log("GroupMembers:", json);
                    this.groupMembers = json.users;
                }));
        },
        createFollowersListForShowing(followers, members) {

            let isUserInGroup = false
            for (let i = 0; i < Object.keys(followers).length; i++) {
                for (let j = 0; j < Object.keys(members).length; j++) {
                    if (followers[i].nickname === members[j].nickname) {
                        isUserInGroup = true
                    }
                }
                if (!isUserInGroup) {
                    this.listForShowing.push(followers[i])
                }
                isUserInGroup = false
            }
        },
        toggleModal() {
            this.isOpen = !this.isOpen;
        },

        getIds() {
            let arrOfIDS = [];
            for (let name of this.checkedNames) {
                for (let obj of this.listForShowing) {
                    if (obj.nickname === name) {
                        arrOfIDS.push(obj.id)
                    }
                }
            }

            return arrOfIDS

        },


        async inviteUsersToGroup() {
            await fetch("http://localhost:8081/newGroupInvite", {
                method: 'POST',
                credentials: 'include',
                body: JSON.stringify({ invitations: this.getIds(), id: this.$route.params.id })
            })
                .then((response => response.json()))
                .then((json => {
                    // console.log("new group invite response:", json);
                    this.clearInput = true;
                    // this.groupMembers = json.users;
                }));
        },

        toggleClearInput() {
            this.clearInput = !this.clearInput
        },
    },
    components: { Modal, MultiselectDropdown }
}
</script>


<style>
</style>