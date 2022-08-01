<template>

    <!-- <button @click="getAllMyPosts">get</button>
    <button @click="showAllMyPosts">show</button> -->
    <!-- <div v-if="this.myposts !== undefined"> -->
    <!-- {{this.posts}} -->
    <Post v-if="posts" v-for="post in this.posts" :key="post.id" v-bind:postData="post" />
    <p class="additional-info large" v-else>No posts to show</p>
    <!-- </div> -->

</template>


<script>
import Post from './Post.vue'
export default {
    name: 'AllMyPosts',
    components: { Post },
    props: {
        userid: ""
    },
    data() {
        return {
            posts: null
        }
    },
    created() {
        // this.getAllMyPosts()
        this.getPosts()
    },
    // computed: mapGetters(['myPosts']),
    watch: { //watching changes in route
        userid() {
            this.getPosts()
        }
    },
    methods: {
        async getPosts() {
            // console.log("USERID", this.userid);
            await fetch("http://localhost:8081/userPosts?id=" + this.userid, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((r) => {
                    // console.log("response", r);
                    this.posts = r.posts
                });
        },
    },
}
</script>


<style>
</style>