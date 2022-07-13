<template>

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
                      required></textarea>
        </div>

        <div class="form-input">
            <div v-for="follower in fetchedFollowers">
                <input type="checkbox"
                       :id="follower"
                       :value="follower"
                       v-model="checkedFollowers" />

                <label :for="follower">{{ follower }}</label>
            </div>
            <div>Checked names: {{ this.checkedFollowers }}</div>
        </div>
        <button class="btn" type="submit">Post</button>
    </form>

</template>

<script>
export default {

    data() {
        return {
            checkedFollowers: [],
            fetchedFollowers: null
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

            console.log("Fetched followers", data)

        },



    }
}

</script>