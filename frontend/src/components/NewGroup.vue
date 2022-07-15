<template>
    <button class="btn" @click="toggleModal">New group<i class="uil uil-plus"></i></button>

    <Modal v-show="isOpen" @closeModal="toggleModal(); toggleClearInput();" v-if="fetchedFollowers">
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

                                     v-model="checkedFollowers"
                                     @inputCleared="toggleClearInput"
                                     :clear-input="clearInput"
                                     :content="fetchedFollowers"
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

            isOpen: false,
            clearInput: false

        }
    },

    created() {
        this.getFollowers();

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

        async getFollowers() {
            const response = await fetch('https://9ec93652-e031-4f3e-a558-86f8ed7d624a.mock.pstmn.io/getfollowers')
            const data = await response.json();
            this.fetchedFollowers = data.followers;

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