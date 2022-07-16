<template>
    <Post v-for="post in this.groupPosts" :key="post.id" v-bind:postData="post" />
</template>


<script>
import Post from './Post.vue'
export default {
    name: 'GroupPosts',
    components: { Post },
    data() {
        return {
            groupPosts: []
        }
    },
    created() {
        this.getGroupPosts();
    },
    methods: {
        async getGroupPosts() {
            await fetch("http://localhost:8081/groupPosts?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((r => r.json()))
                .then((json => {
                    console.log(json)
                    this.groupPosts = json.posts
                }))

        }
    }
}
</script>


<style>
</style>