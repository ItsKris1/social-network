<template>
    <NavBarOn />
    <NewPost />

</template>

<script>
import NavBarOn from '@/components/NavBarOn.vue'
import NewPost from '@/components/NewPost.vue'
export default {
    name: 'MainView',
    components: { NavBarOn, NewPost, },
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