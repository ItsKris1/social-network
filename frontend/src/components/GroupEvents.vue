<template>
    <div class="item-list__wrapper" id="groups">
        <h3>Events:</h3>
        <ul class="item-list">
            <li v-for="event in this.groupEvents">                
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text" style="cursor:pointer" @click="this.showEvent(event)">{{ event.title }}</div>

                <Modal v-if="this.eventIsOpen" @closeModal="closeEvent">
                    <template #title>
                        {{ this.eventData.title }}
                    </template>
                    <template #body>
                        <div>{{this.eventData.date}}</div>
                        <div>{{this.eventData.content}}</div>

                    </template>
                </Modal>
            </li>
        </ul>
        <button class="btn form-submit" @click="toggleModal">New event</button>

        <Modal v-if="this.isOpen" @closeModal="toggleModal">
            <template #title>Crete new event</template>
            <template #body>
                <span>Title</span>
                <input type="text" v-model="this.formData.title">
                <span>Description</span>
                <textarea v-model="this.formData.content" cols="30" rows="3"
                    placeholder="What is this about?"></textarea>
                <span>Date</span>
                <input v-model="this.formData.date" type="date">
                <span>Going?</span>
                <select v-model="this.formData.going" name="going">
                    <option value="YES">Yes</option>
                    <option value="NO">No</option>
                </select>
                <button class="btn form-submit" @click="toggleModal() ; createNewEvent()">Create</button>
            </template>

        </Modal>
    </div>
</template>


<script>
import Modal from './Modal.vue';
export default {
    name: "GroupEvents",
    data() {
        return {
            groupEvents: [],
            isOpen: false,
            eventIsOpen: false,
            eventData: {
                title: "",
                date: "",
                content: "",
                going: "",
            },
            formData: {
                title: "",
                content: "",
                date: null,
                going: null
            }
        };
    },
    created() {
        this.getGroupEvents();
    },
    watch:{
        $route(){
            this.getGroupEvents()
        }
    },
    methods: {
        async getGroupEvents() {
            await fetch("http://localhost:8081/groupEvents?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((response => response.json()))
                .then((json => {
                    console.log("All Events:", json);
                    this.groupEvents = json.events
                }));
        },
        async createNewEvent() {

            await fetch('http://localhost:8081/newEvent', {
                method: 'POST',
                credentials: 'include',
                body: JSON.stringify({
                    groupID: this.$route.params.id,
                    title: this.formData.title,
                    content: this.formData.content,
                    date: this.formData.date,
                    going: this.formData.going
                })
            })
                .then((r => r.json()))
                .then((json => {
                    console.log("newEvent", json);
                }))
            this.getGroupEvents()
        },
        toggleModal() {
            this.isOpen = !this.isOpen;
        },
        showEvent(event) {
            this.eventData.title = event.title
            this.eventData.date = event.date
            this.eventData.content = event.content
            this.eventData.going = event.going

            this.eventIsOpen = true
        },
        closeEvent() {
            this.eventIsOpen = false
        },
    },
    components: { Modal }
}
</script>


<style>
</style>