

<template>
    <button class="start-post" @click="toggle">
        <span>Start post</span>
        <i class="uil uil-edit"></i>
    </button>

    <Modal v-show="isOpen" @closeModal="toggle">
        <template #title>Create a post</template>
        <template #body>
            <form @submit.prevent="submitPost" id="newpost" ref="theForm">
                <div class="form-input">
                    <label for="post_privacy">Post privacy</label>
                    <div class="select-wrapper">
                        <img src="../assets/icons/angle-down.svg" class="dropdown-arrow">

                        <select id="post_privacy" v-model="newpost.privacy" required>
                            <option value="" selected hidden>Choose here</option>
                            <option value="public">Everyone</option>
                            <option value="private">Followers</option>
                            <option value="almost-private">Choosen followers</option>
                        </select>

                    </div>

                    <MultiselectDropdown v-if="fetchedFollowers && newpost.privacy === 'almost-private'"
                                         placeholder="Select followers"
                                         :content="fetchedFollowers" />
                </div>

                <div class="form-input">
                    <label for="description">Description</label>
                    <textarea id="description" v-model="newpost.body" rows="4" cols="50"
                              placeholder="What are you thinking?" required></textarea>
                </div>

                <div class="btns-wrapper">

                    <label for="upload-image">
                        <img src="../assets/addimg.png" />
                    </label>
                    <input id="upload-image" type="file"
                           accept="image/png, image/gif, image/jpeg" @change="checkPicture" />

                    <button class="btn" type="submit">Post</button>

                </div>
            </form>

        </template>
    </Modal>

</template>



<script>
import Modal from './Modal.vue'
import MultiselectDropdown from './MultiselectDropdown.vue';
export default {
    components: {
        Modal,
        MultiselectDropdown
    },
    name: 'Newpost',
    data() {
        return {
            isOpen: false,
            newpost: {
                privacy: "",
                body: "",
                image: null,
            },
            fetchedFollowers: null,
        }
    },

    created() {
        this.getFollowers();
    },



    methods: {
        toggle() {
            // this.newpost.privacy = "";
            // this.newpost.body = "";
            // this.newpost.image = "";

            // if (this.$refs.theForm && this.isOpen) {
            //     console.log(this.$refs.theForm);
            //     this.$refs.theForm.reset()
            // }
            this.isOpen = !this.isOpen
        },

        checkPicture(e) {
            let files = e.target.files
            if (!files.length) {
                return;
            }
            const file = files[0]

            const [extension] = file.type.split("/")
            if ((!(extension == "image"))) {
                console.log('File is not an image.');
                this.$toast.open({
                    message: 'File is not an image.',
                    type: 'error', //One of success, info, warning, error, default
                })
                return
            }
            if (file.size > 2048000) {
                console.log('File size is more than 2 MB.');
                this.$toast.open({
                    message: 'File size is more than 2 MB.',
                    type: 'error', //One of success, info, warning, error, default
                })
                return
            }
            this.newpost.image = file;

        },

        async getFollowers() {
            // console.log('fetching followers');
            await fetch('https://9ec93652-e031-4f3e-a558-86f8ed7d624a.mock.pstmn.io/getfollowers')
                .then((response => response.json()))
                .then((json) => {
                    this.fetchedFollowers = json.followers
                })
            // .then((json => console.log(json)))


        },

        async submitPost() {

            let formData = new FormData();
            formData.set("body", this.newpost.body)
            formData.set("image", this.newpost.image)
            formData.set("privacy", this.newpost.privacy)
            f

            console.log("Data", Object.fromEntries(formData.entries()))

            const response = await fetch('http://localhost:8081/newPost', {
                method: 'POST',
                credentials: 'include',
                body: formData
            })
            this.$store.dispatch('fetchPosts')
            console.log('Post submitted', await response.json());


            this.toggle();
        },
    },
}
</script>


<style>
#newpost {
    display: flex;
    flex-direction: column;
}

.image-upload>input {
    display: none;
}

#postBtn {
    width: 600px;
    display: flex;
    justify-content: center;
    border: 1px solid #706A6A;
    border-radius: 3px;
    border: 1px solid #706A6A;
}

.start-post {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px 15px;
    background-color: var(--input-bg);
    border: none;
    box-shadow: var(--container-shadow);
    font-family: inherit;
    font-size: 16px;
    border-radius: var(--container-border-radius);
}

.start-post i {
    font-size: 1.25em;
}

.select-wrapper {
    position: relative;
}
</style>