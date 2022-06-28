<template>
    <NavBarOn />
    <NewPost />
    <AllPosts />
    <Groups />
</template>

<script>
import NavBarOn from '@/components/NavBarOn.vue'
import NewPost from '@/components/NewPost.vue'
import AllPosts from '@/components/AllPosts.vue'
import Groups from '@/components/Groups.vue'

export default {
    name: 'MainView',
    components: { NavBarOn, NewPost, AllPosts, Groups },
}
</script>

<style>
</style>