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

        <div class="form-input" id="followersDropdown">
            <p>Invite users</p>
            <ul class="checkedFollowersList">
                <li v-for="checkedFollower in checkedFollowers"> {{ checkedFollower }}</li>
            </ul>

            <button type="button" @click.self="showDropdown = !showDropdown">Select users</button>

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
        <button class="btn" type="submit">Post</button>
    </form>

</template>

<style>
#followersDropdown {
    width: 250px;
    background: var(--input-bg);
}

#followersDropdown button {
    padding: 7.5px;
    border-radius: 5px;
    background-color: var(--input-bg);
    border: none;
    box-shadow: var(--container-shadow);
    font-family: 'Poppins', sans-serif;
    text-align: left;
    color: rgb(136, 136, 136);
    width: 250px;


}

#followersDropdown .item-list {
    width: 100%;
}


.checkedFollowersList {
    display: flex;
    gap: 5px;
}

.checkedFollowersList li {
    background-color: rgb(179, 179, 179);
    border-radius: 5px;
    padding: 5px;
    font-size: 14px;
}
</style>
<script>
export default {

    data() {
        return {
            checkedFollowers: [],
            fetchedFollowers: null,

            showDropdown: false,
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