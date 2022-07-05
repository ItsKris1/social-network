<template>

    <!-- <button @click="getAllMyPosts">get</button>
    <button @click="showAllMyPosts">show</button> -->
    <!-- <div v-if="this.myposts !== undefined"> -->
        {{this.userid}}
    <Post v-for="post in myPosts" :key="post.id" v-bind:postData="post" />
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
        }
    },
    created() {
        // this.getAllMyPosts()
        this.getPosts()
    },
    computed: mapGetters(['myPosts']),
    // computed:{
    //     // getId(){
    //     //     this.userId = this.$route.params.id
    //     //     console.log("id",this.userId);
    //     // }
    // },
    methods: {
        // getAllMyPosts() {
        //     this.$store.dispatch('fetchMyPosts')
        // },
        async getPosts() {
            console.log("USERID", this.userid);
            await fetch("http://localhost:8081/userPosts?id=" + this.userid, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((r) => {
                    console.log("response",r);
                    // const myposts = r.posts;
                    // this.commit("updateMyPosts", myposts);
                });
        },
        // showAllMyPosts() {
        //     console.log('All posts: ', myPosts);
        // },
    },
}
</script>


<style>
</style>