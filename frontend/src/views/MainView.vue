<template>
    <NavBarOn />
    <NewPost />
    <AllPosts />

</template>

<script>
import NavBarOn from '@/components/NavBarOn.vue'
import NewPost from '@/components/NewPost.vue'
import AllPosts from '@/components/AllPosts.vue'
export default {
    name: 'MainView',
    components: { NavBarOn, NewPost, AllPosts},
    data() {
        return {
            maindata: {
                posts: [],
                users: [],
                groups: [],
            }
        }
    },
}

</script>

<style>
</style>