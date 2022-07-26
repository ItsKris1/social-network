<template>
    <div class="item-list__wrapper" id="groups">
        <h3>Members:</h3>
        <ul class="item-list">
            <li v-for="member in this.groupMembers">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">{{ member.nickname }}</div>
            </li>
        </ul>
        <button class="btn form-submit" @click="toggleModal">Invite users</button>

        <Modal v-if="this.isOpen" @closeModal="toggleModal">
            <template #title>Choose users:</template>
            <template #body>
                {{ this.checkedNames }}
                <div v-for="user in this.allUsers">
                    <input type="checkbox" :value="user.nickname" v-model="checkedNames">{{ user.nickname }}
                </div>
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
        };
    },
    created() {
        this.getGroupMembers();
        this.getAllUsers();
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

            let formData = new FormData();
            formData.set("checkedNames", this.checkedNames)
            formData.set("groupId", this.$route.params.id)

            await fetch("http://localhost:8081/newGroupInvite", {
                method: 'POST',
                credentials: 'include',
                body: formData
            })
                .then((response => response.json()))
                .then((json => {
                    console.log("new group invite response:", json);
                    // this.groupMembers = json.users;
                }));
        },
    },
    components: { Modal, MultiselectDropdown }
}
</script>


<style>
</style>