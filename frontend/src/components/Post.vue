<template>

    <!-- <div id="post">
        <button @click="showPostId(postData.id)">Show post id!</button>
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
                <textarea v-model="this.comment.body" name="" id="" cols="30" rows="5"
                    placeholder="Add your comment here"></textarea>
                <div>
                    <button @click="toggleComments">Hide comments</button>
                    <div class="comment-img">
                        <label for="comment-image">
                            <img src="../assets/addimg.png" />
                        </label>
                        <input id="comment-image" @change="checkPicture" type="file"
                            accept="image/png, image/gif, image/jpeg" />
                    </div>
                    <button @click="submitComment(postData.id)">Add comment</button>
                </div>
                <div id="commentsDiv" v-for=" comment in postData.comments">
                <img class="commentAuthorAvatar" :src="'http://localhost:8081/' + comment.author.avatar" alt="commentAuthorAvatar">
                    <div><b>{{ comment.author.nickname }}</b></div>
                    <div>{{ comment.content }}</div>

                    <div v-if="comment.image">
                        <img id="commentImage" :src="'http://localhost:8081/' + comment.image" alt="">
                    </div>
                </div>
            </div>
        </div>
    </div> -->

    <div class="post-wrapper">
        <div class="post">
            <div class="user-picture medium"
                 :style="{ backgroundImage: `url(http://localhost:8081/${postData.author.avatar})` }"></div>
            <div class="post-content">
                <p class="post-author">{{ postData.author.nickname }}</p>
                <p class="post-body">{{ postData.content }}</p>
                <img v-if="postData.image" class="post-image" :src="'http://localhost:8081/' + postData.image" alt="">
                <button v-if="!isCommentsOpen" @click="toggleComments" class="btn ">Comments</button>

            </div>

        </div>

        <div v-if="isCommentsOpen">

            <div class="create-comment">
                <textarea v-model="this.comment.body" name="" id="" cols="30" rows="4"
                          placeholder="Add your comment here"></textarea>

                <div class="create-comment__btns">
                    <button class="btn outline" @click="toggleComments">Hide comments</button>

                    <div class="btns-wrapper">

                        <label for="upload-image">
                            <img src="../assets/addimg.png" />
                        </label>
                        <input id="upload-image" @change="checkPicture" type="file"
                               accept="image/png, image/gif, image/jpeg" style="display: none" />

                        <button class="btn" @click="submitComment(postData.id)">Comment</button>

                    </div>

                </div>
            </div>

            <div class="comments" v-if="postData.comments">
                <div class="comment" lang="en" v-for="comment in postData.comments">
                    <div class="user-picture medium"
                         :style="{ backgroundImage: `url(http://localhost:8081/${comment.author.avatar})` }"></div>
                    <div class="comment-content">
                        <p class="comment-author">{{ comment.author.nickname }}</p>
                        <p class="comment-body">{{ comment.content }}</p>
                        <img class="comment-image" v-if="comment.image" :src="'http://localhost:8081/' + comment.image"
                             alt="">
                    </div>
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
            this.$store.dispatch('fetchPosts')
            this.$store.dispatch('fetchMyPosts')
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


<style scoped>
.post-wrapper {
    display: inline-block;
    box-shadow: var(--container-shadow);
    padding: 30px;
    background-color: var(--color-white);
    width: 550px;
    border-radius: 10px;
}



.post {
    display: flex;
    gap: 10px;
}



.post-author,
.comment-author {
    font-weight: 500;
}

.post-content,
.comment-content {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
    flex-grow: 1;
}

.post-image,
.comment-image {
    width: 100%;
    margin: 10px 0 10px 0;
    border-radius: 5px;
}

.post-body,
.comment-body {
    overflow-wrap: anywhere;
}

.post-content button {
    align-self: flex-end;
    margin-top: 10px;
}




.create-comment {
    padding-left: 58px;
}

.create-comment textarea {
    margin: 10px 0;
}

.create-comment__btns {
    display: flex;
    justify-content: space-between;
}

.btns-wrapper {
    display: flex;
    gap: 10px;
    align-items: center;
    justify-content: flex-end;

}

.btns-wrapper label {
    margin: 0;
}

.btns-wrapper input {
    display: none;
}

.btns-wrapper label img {
    vertical-align: middle;
}




.comments>* {
    display: flex;
    gap: 10px;
    border-top: 1px solid #DDDDDD;
    padding-top: 30px;
}

.comments {
    display: flex;
    flex-direction: column;
    gap: 30px;
    margin-top: 30px;
}
</style>