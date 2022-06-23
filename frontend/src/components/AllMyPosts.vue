<template>

    <!-- <button @click="getAllMyPosts">get</button>
    <button @click="showAllMyPosts">show</button> -->
    <!-- <div v-if="this.myposts !== undefined"> -->
        <Post v-for="post in myPosts" :key="post.id" v-bind:postData="post" />
    <!-- </div> -->

</template>


<script>
import { mapGetters } from 'vuex'
import Post from './Post.vue'
export default {
    name: 'AllMyPosts',
    created() {
        this.getAllMyPosts()
    },
    components: { Post },
    computed: mapGetters(['myPosts']),
    methods: {
        getAllMyPosts() {
            this.$store.dispatch('fetchMyPosts')
        },
        showAllMyPosts() {
            console.log('All posts: ', myPosts);
        },
    },
}
</script>


<style>
</style>