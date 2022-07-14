<template>
    <button class="btn" @click="toggle">New group<i class="uil uil-plus"></i></button>

    <Modal v-show="isOpen" @closeModal="toggle">
        <template #title>
            Create new group
        </template>

        <template #body>
            <form @submit.prevent="submitNewGroup" v-if="fetchedFollowers">
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

                <div class="form-input">
                    <p class="custom-label">Invite users</p>

                    <ul class="checkedFollowersList" v-if="checkedFollowers.length !== 0">
                        <li v-for="checkedFollower in checkedFollowers"> {{ checkedFollower }}</li>
                    </ul>

                    <div class="followers-dropdown">
                        <button type="button" @click="showDropdown = !showDropdown">
                            Select users
                            <img class="dropdown-arrow" src="../assets/icons/angle-down.svg" alt="" srcset="">
                        </button>

                        <ul class="item-list" v-show="showDropdown">
                            <li v-for="follower in fetchedFollowers">
                                <input type="checkbox"
                                       :id="follower"
                                       :value="follower"
                                       v-model="checkedFollowers" />
                                <label :for="follower">{{ follower }}</label>

                            </li>
                        </ul>
                    </div>

                </div>
                <button class="btn form-submit" type="submit">Create</button>
            </form>

        </template>
    </Modal>

</template>


<script>
import Modal from "@/components/Modal.vue"

export default {
    components: {
        Modal
    },
    data() {
        return {
            checkedFollowers: [],
            fetchedFollowers: null,

            // Element show/hide
            showDropdown: false,
            isOpen: false,
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
            formDataObject["invitations"] = this.checkedFollowers;

            const response = await fetch('http://localhost:8081/newGroup', {
                method: 'post',
                credentials: 'include',
                body: JSON.stringify(formDataObject)
            })

            form.reset()
            this.checkedFollowers = [];

            console.log("Submit new group", await response.json())


        },

        async getFollowers() {
            const response = await fetch('https://9ec93652-e031-4f3e-a558-86f8ed7d624a.mock.pstmn.io/getfollowers')
            const data = await response.json();
            this.fetchedFollowers = data.followers;

            // console.log("Fetched followers", data)

        },

        toggle() {
            this.isOpen = !this.isOpen
        },

    }
}

</script>

<style>
.followers-dropdown {
    background-color: var(--input-bg);
    box-shadow: var(--container-shadow);
    border-radius: 5px;
}

.followers-dropdown button {
    padding: 7.5px;

    border: none;
    font-family: 'Poppins', sans-serif;
    text-align: left;
    color: rgb(136, 136, 136);
    /* width: 250px; */
    position: relative;
    background-color: transparent;
    width: 100%;
}

.followers-dropdown .item-list {
    padding: 7.5px;
    width: 100%;
}



.checkedFollowersList {
    display: flex;
    gap: 5px;
    padding: 5px 0;
}

.checkedFollowersList li {
    background-color: rgb(179, 179, 179);
    border-radius: 5px;
    padding: 5px;
    font-size: 14px;
}
</style>