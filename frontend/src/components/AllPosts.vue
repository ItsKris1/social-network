<template>
    <button @click="getAllPosts">get</button>
    <button @click="showAllPosts">show</button>
    <div id="all_posts" v-if="this.posts !== undefined">
        <Post v-for="post in posts" :key="post.id" v-bind:postData="post" />
    </div>

</template>


<script>
import Post from './Post.vue'
export default {
    name: 'AllPosts',
    data() {
        return {
            posts: []
        }
    },
    created() {
        this.getAllPosts()
    },
    components: { Post },
    methods: {
        async getAllPosts() {
            // console.log('allposts');
            // await fetch("https://d1cd8cff-f1bc-4c79-913a-405f7cbbb7c3.mock.pstmn.io/allposts")
            await fetch("http://localhost:8081/allPosts", {
                credentials: 'include',
            })
                // .then((r => console.log(r)))
                .then((res => res.json()))
                .then((json => {
                    this.posts = json.posts
                    // return json.posts
                }))
            // .then((r => console.log(r)))

        },
        showAllPosts() {
            console.log('All posts: ', this.posts);
        },
    },
}
</script>


<style>
#all_posts {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>