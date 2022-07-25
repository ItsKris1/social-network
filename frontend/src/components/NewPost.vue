

<template>

    <button class="start-post" @click="toggleModal">
        <span>Start post</span>
        <i class="uil uil-edit"></i>
    </button>

    <Modal v-show="isOpen" @closeModal="toggleModal">
        <template #title>Create a post</template>
        <template #body>
            <form v-if="this.isGroupPage" @submit.prevent="submitGroupPost" id="newpost">
                <div class="form-input">
                    <label for="description">Description</label>
                    <textarea id="description" name="description" rows="4" cols="50" v-model="newpost.body"
                              placeholder="What are you thinking?" required></textarea>
                </div>

                <div class="btns-wrapper">

                    <label for="upload-image">
                        <img src="../assets/addimg.png" />
                    </label>
                    <input id="upload-image" @change="checkPicture" type="file"
                           accept="image/png, image/gif, image/jpeg" />

                    <button class="btn" type="submit">Post</button>

                </div>
            </form>

            <form v-else @submit.prevent="submitPost" id="newpost">
                <div class="form-input">
                    <label for="post_privacy">Post privacy</label>
                    <div class="select-wrapper">
                        <img src="../assets/icons/angle-down.svg" class="dropdown-arrow">

                        <select id="post_privacy" v-model="newpost.privacy" required>
                            <option value="" selected hidden>Choose here</option>
                            <option value="public">Everyone</option>
                            <!-- <select v-model="newpost.privacy" @change="getFollowers" id="post_privacy" required>
                            <option value="" selected disabled hidden>Choose here</option> -->
                            <option value="private">Followers</option>
                            <option value="almost-private">Choosen followers</option>
                        </select>

                    </div>

                    <MultiselectDropdown v-if="newpost.privacy === 'almost-private'"
                                         v-model:checkedOptions="newpost.checkedFollowers"
                                         placeholder="Select followers"
                                         :content="getMyFollowersNames" :clearInput="clearInput"
                                         @inputCleared="toggleClearInput" />
                </div>

                <div class="form-input">
                    <label for="description">Description</label>

                    <textarea id="description" v-model="newpost.body" rows="4" cols="50"
                              placeholder="What are you thinking?" required></textarea>

                </div>

                <div class="btns-wrapper">
                    <div class="add-image">
                        <div class="selected-image" v-if="fileAdded">
                            <p class="additional-info">{{ newpost.image.name }}</p>
                            <i class="uil uil-times close" @click="removeImage"></i>
                        </div>

                        <p class="additional-info" v-else>No file chosen
                        </p>

                        <label for="upload-image">
                            <input type="file" accept="image/png, image/gif, image/jpeg" style=""
                                   @change="checkPicture" ref="fileUpload" />
                        </label>
                    </div>

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
            isGroupPage: null,
            newpost: {
                privacy: "",
                body: "",
                checkedFollowers: null,
                image: {},
            },
            clearInput: false,
        }
    },

    created() {
        this.getMyFollowers();
        this.isGroupPageCheck()
    },

    computed: {
        getMyFollowersNames() {
            return this.$store.getters.getMyFollowersNames;
        },

        fileAdded() {
            return this.newpost.image.name !== undefined
        },
    },

    methods: {
        toggleModal() {
            // if modal was open, clear the form
            if (this.isOpen) { this.clearForm(); }
            this.isOpen = !this.isOpen
        },

        toggleClearInput() {
            this.clearInput = !this.clearInput
        },



        getMyFollowers() {
            this.$store.dispatch("getMyFollowers")
        },

        clearForm() {
            this.newpost.privacy = "";
            this.newpost.body = "";
            this.newpost.image = "";
            this.toggleClearInput();
        },

        isGroupPageCheck() {
            if (this.$route.path.includes("group")) {
                this.isGroupPage = true
            } else {
                this.isGroupPage = false
            }
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

        removeImage() {
            this.comment.image = {};
            this.$refs.fileUpload.value = "";
        },

        async submitPost() {

            let formData = new FormData();
            formData.set("body", this.newpost.body)
            formData.set("image", this.newpost.image)
            formData.set("privacy", this.newpost.privacy)
            formData.set("checkedfollowers", this.newpost.checkedFollowers)


            const response = await fetch('http://localhost:8081/newPost', {
                method: 'POST',
                credentials: 'include',
                body: formData
            })
            this.$store.dispatch('fetchPosts')
            console.log('Post submitted', await response.json());
            // console.log('Post submitted');
            this.toggleModal();
        },

        async submitGroupPost() {

            let formData = new FormData();
            formData.set('groupId', this.$route.params.id)
            formData.set('body', this.newpost.body);
            formData.set('image', this.newpost.image);

            this.toggleModal();
            await fetch('http://localhost:8081/newGroupPost', {
                method: 'POST',
                credentials: 'include',
                body: formData
            })
                .then((r => r.json()))
            // .then((json => console.log(json)))
            this.$store.dispatch('getGroupPosts')
            this.toggleModal();
            // console.log('Group Post Submitted');
        },
    }
}
</script>


<style scoped>
#newpost {
    display: flex;
    flex-direction: column;
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


.additional-info {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.btns-wrapper {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 10px;
}
</style>