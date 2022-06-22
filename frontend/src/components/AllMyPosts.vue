<template>

    <!-- <button @click="getAllMyPosts">get</button>
    <button @click="showAllMyPosts">show</button> -->
    <div v-if="this.posts !== undefined">
        <Post v-for="post in posts" :key="post.id" v-bind:postData="post" />
    </div>

</template>


<script>
import Post from './Post.vue'
export default {
    name: 'AllMyPosts',
    created() {
        this.getAllMyPosts()
    },
    components: { Post },
    methods: {
        getAllMyPosts() {
            this.$store.dispatch('fetchMyPosts')
        },


        // async getAllMyPosts() {
        //     // console.log('allposts');
        //     // await fetch("https://d1cd8cff-f1bc-4c79-913a-405f7cbbb7c3.mock.pstmn.io/allposts")
        //     await fetch("http://localhost:8081/userPosts", {
        //         credentials: 'include',
        //     })
        //         .then((r => console.log(r)))
        //         .then((res => res.json()))
        //         .then((json=>{
        //             this.posts = json.posts
        //             return json.posts
        //         }))
        //         .then((r => console.log(r)))

        // },
        showAllMyPosts() {
            console.log('All posts: ', this.posts);
        },
    },
}
</script>


<style>
</style>