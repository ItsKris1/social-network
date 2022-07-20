<template>
    <button class="btn" @click="toggleModal">New group<i class="uil uil-plus"></i></button>

    <Modal v-show="isOpen" @closeModal="toggleModal(); toggleClearInput();"
           v-if="isMyFollowersFetched">
        <template #title>
            Create new group
        </template>

        <template #body>
            <form @submit.prevent="submitNewGroup" ref="theForm">
                <div class="form-input">
                    <label for="name">Name</label>
                    <input type="text" name="name" id="name">
                </div>

                <div class="form-input">
                    <label for="description">Description</label>
                    <textarea id="description"
                              name="description"
                              rows="4"
                              cols="50"
                              required
                              placeholder="Describe here"></textarea>
                </div>

                <MultiselectDropdown

                                     v-model:checkedOptions="checkedFollowers"
                                     @inputCleared="toggleClearInput"
                                     :clear-input="clearInput"
                                     :content="getMyFollowersNames"
                                     label-name="Invite users"
                                     placeholder="Select users" />

                <button class="btn form-submit" type="submit">Create</button>
            </form>

        </template>
    </Modal>
</template>


<script>
import Modal from "@/components/Modal.vue"
import MultiselectDropdown from "./MultiselectDropdown.vue";

export default {
    components: {
        Modal,
        MultiselectDropdown,
    },
    data() {
        return {
            fetchedFollowers: null,
            checkedFollowers: null,
            myFollowers: {},
            isOpen: false,
            clearInput: false

        }
    },

    created() {
        this.getMyFollowers();

    },

    computed: {
        getMyFollowersNames() {
            return this.$store.getters.getMyFollowersNames;
        },
    },


    methods: {
        async submitNewGroup(e) {
            const form = e.currentTarget;
            const formData = new FormData(form);
            const formDataObject = Object.fromEntries(formData.entries())
            formDataObject["invitations"] = Object.values(this.checkedFollowers);

            const response = await fetch('http://localhost:8081/newGroup', {
                method: 'post',
                credentials: 'include',
                body: JSON.stringify(formDataObject)
            })
            console.log("Submit new group", await response.json())

            form.reset()
            this.toggleModal();
            this.toggleClearInput();



        },

        getMyFollowers() {
            this.$store.dispatch("getMyFollowers")
        },

        isMyFollowersFetched() {
            const myFollowers = this.$store.state.myFollowers
            return Object.keys(myFollowers.length !== 0)
        },

        toggleModal() {
            // if modal was open: clear input
            if (this.isOpen) {
                // this.form.reset();
                this.$refs.theForm.reset();
            }
            this.isOpen = !this.isOpen


        },

        toggleClearInput() {
            this.clearInput = !this.clearInput
        },

    }
}

</script>