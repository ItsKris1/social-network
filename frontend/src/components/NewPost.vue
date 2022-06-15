<template>
    <div @click="toggle">Start a post</div>
    <div v-show="isOpen">
        <form @submit.prevent="submitPost" id="newpost">
            <span><u>Create a post</u></span>
            <div>Body</div>
            <input v-model="newpost.body" type="text">
            <div class="image-upload">
                <label for="file-input">
                    <img src="../assets/addimg.png" />
                </label>
                <input @change="checkPicture" id="file-input" type="file" accept="image/png, image/gif, image/jpeg" />
            </div>
            <button type="submit">Post</button>
        </form>
    </div>
</template>



<script>
export default {
    name: 'Newpost',
    data() {
        return {
            isOpen: false,
            newpost: {
                body: "",
                image: null,
            }
        }
    },
    methods: {
        toggle() {
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
        async submitPost() {            

            let formData = new FormData();
            formData.set('body', this.newpost.body);
            formData.set('image', this.newpost.image);

            await fetch('', {
                method: 'POST',
                body: formData
            })

            console.log('Post submitted');
        }
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
</style>