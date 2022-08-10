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

                <MultiselectDropdown
                                     v-model:checkedOptions="checkedNames"
                                     placeholder="Select users"
                                     :content="allUserNames" :clearInput="clearInput"
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
            allUsers: [],
            checkedNames: [],
            clearInput: false,
        };
    },
    created() {
        this.getGroupMembers();
        this.getAllUsers();
    },

    computed: {
        allUserNames() {
            return this.allUsers.map(user => user.nickname)
        }
    },

    watch: {
        $route() {
            this.getGroupMembers();
        }
    },
    methods: {
        async getAllUsers() {
            await fetch("http://localhost:8081/allUsers", {
                credentials: "include"
            })
                .then((response => response.json()))
                .then((json => {
                    // console.log("Allusers:", json);
                    this.allUsers = json.users;
                }));
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
        toggleModal() {
            this.isOpen = !this.isOpen;
        },
        async inviteUsersToGroup() {
            await fetch("http://localhost:8081/newGroupInvite", {
                method: 'POST',
                credentials: 'include',
                body: JSON.stringify({ invitations: this.checkedNames, id: this.$route.params.id })
            })
                .then((response => response.json()))
                .then((json => {
                    console.log("new group invite response:", json);
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