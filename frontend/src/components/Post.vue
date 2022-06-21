<template>

    <div id="post">
        <!-- <button @click="showPostId(postData.id)">Show post id!</button> -->
        <div>
            <img id="post_image" :src="'http://localhost:8081/' + postData.author.avatar" alt="user-avatar">
        </div>
        <div>
            <div><b>{{ postData.author.nickname }}</b></div>
            <div v-if="postData.image">
                <img id="postImage" :src="'http://localhost:8081/' + postData.image" alt="">
            </div>

            <div>{{ postData.content }}</div>
            <button v-if="!isCommentsOpen" @click="toggleComments">View comments</button>
            <div v-if="isCommentsOpen">
                <textarea name="" id="" cols="30" rows="5" placeholder="Add your comment here"></textarea>
                <div>
                    <button @click="toggleComments">Hide comments</button>
                    <div class="image-upload">
                        <label for="file-input">
                            <img src="../assets/addimg.png" />
                        </label>
                        <input @change="checkPicture" id="file-input" type="file"
                            accept="image/png, image/gif, image/jpeg" />
                    </div>
                    <button @click="submitComment(postData.id)">Add comment</button>
                </div>
                <div v-for=" comment in postData.Comments">
                    <div><b>{{ comment.author }}</b></div>
                    <div>{{ comment.content }}</div>
                </div>
            </div>
        </div>
    </div>
</template>


<script>
export default {
    name: 'Post',
    data() {
        return {
            isCommentsOpen: false,
            comment: {
                body: "",
                image: null
            }
        }
    },
    props: {
        postData: {
            type: Object,
            default() {
                return {}
            }
        }
    },
    methods: {
        toggleComments() {
            this.isCommentsOpen = !this.isCommentsOpen
        },
        async submitComment(post_id) {

            let commentData = new FormData();
            commentData.set('postid', post_id);
            commentData.set('body', this.comment.body);
            commentData.set('image', this.comment.image);

            await fetch('http://localhost:8081/newComment', {
                method: 'POST',
                credentials: 'include',
                body: commentData
            })
            console.log('Comment submitted.');
        },
        showPostId(postId) {

            console.log('post id: ', postId);
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
            this.comment.image = file;
        },
    }
}
</script>


<style>
#post_image {
    height: 47px;
    width: 47px;
    border-radius: 50%;
    margin-right: 8px;
}

#post {
    display: flex;
}

#postImage {
    height: 50%;
    width: 50%;
}
</style>