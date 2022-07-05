<template>

    <!-- <button @click="getAllMyPosts">get</button>
    <button @click="showAllMyPosts">show</button> -->
    <!-- <div v-if="this.myposts !== undefined"> -->
        <!-- {{this.posts}} -->
    <Post v-for="post in this.posts" :key="post.id" v-bind:postData="post" />
    <!-- </div> -->

</template>


<script>
import { mapGetters } from 'vuex'
import Post from './Post.vue'
export default {
    name: 'AllMyPosts',
    components: { Post },
    props:{
        userid:""
    },
    data() {
        return {
            posts:{}
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
                    // console.log("response",r);
                    this.posts = r.posts
                });
        },
    },
}
</script>


<style>
</style>