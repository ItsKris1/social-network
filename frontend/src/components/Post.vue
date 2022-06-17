<template>
    <div id="post">
        <div>
            <img id="post_image" :src="'http://localhost:8081/' + postData.author.avatar" alt="user-avatar">
        </div>
        <div>
            <div><b>{{ postData.author.nickname }}</b></div>
            <div v-if="postData.image">
            <img  id="postImage" :src="'http://localhost:8081/' + postData.image" alt="">
            </div>
            
            <div>{{ postData.content }}</div>
            <button v-if="!isCommentsOpen" @click="toggleComments">View comments</button>
            <div v-if="isCommentsOpen">
            <textarea name="" id="" cols="30" rows="5" placeholder="Add your comment here"></textarea>
            <div>
                <button @click="toggleComments">Hide comments</button>
                <button @click="submitComment">Add comment</button>
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
            isCommentsOpen: false
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
        submitComment(){
            console.log('Comment submitted.');
        }
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

#postImage{
    height: 50%;
    width: 50%;
}
</style>