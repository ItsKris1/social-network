<template>
    <div class="item-list__wrapper" id="groups">
        <h3>Events:</h3>
        <ul class="item-list">
            <li v-for="event in this.groupEvents">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">{{ event }}</div>
            </li>
        </ul>
        <button class="btn form-submit" @click="toggleModal">New event</button>

        <Modal v-if="this.isOpen" @closeModal="toggleModal">
            <!-- <template #title>Choose users:</template>
            <template #body>
                
                <button class="btn form-submit" @click="toggleModal() ; inviteUsersToGroup()">Invite</button>
            </template> -->
        </Modal>
    </div>
</template>


<script>
import Modal from './Modal.vue';
export default {
    name: "GroupEvents",
    data() {
        return {
            groupEvents: ["a", "b", "c"],
            isOpen:false
        };
    },
    created() {
        this.getGroupEvents();
    },
    methods: {
        async getGroupEvents() {
            await fetch("http://localhost:8081/groupEvents?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((response => response.json()))
                .then((json => {
                console.log("All Events:", json);
                // this.GroupEvents = json.users;
            }));
        },
        toggleModal() {
            this.isOpen = !this.isOpen;
        },
    },
    components: { Modal }
}
</script>


<style>
</style>