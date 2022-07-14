<template>
    <NavBarOn />

    <div id="layout">
        <NewPost />
        <Groups />
        <AllPosts />
    </div>

</template>

<script>
import NavBarOn from '@/components/NavBarOn.vue'
import NewPost from '@/components/NewPost.vue'
import AllPosts from '@/components/AllPosts.vue'
import Groups from '@/components/Groups.vue'
import NewGroup from '@/components/NewGroup.vue'

export default {
    name: 'MainView',
    components: { NavBarOn, NewPost, AllPosts, Groups, NewGroup },
}
</script>

<style>
#layout {
    display: grid;
    grid-template-columns: 1fr minmax(min-content, 500px) 1fr;
    grid-template-areas:
        "groups startpost ."
        "groups posts .";

    align-items: flex-start;
    row-gap: 50px;
    column-gap: 50px;
    margin-top: 100px;
    margin-bottom: 50px;
}


#groups {
    grid-area: groups;
    justify-self: end;
}

#all_posts {
    grid-area: posts;
}

.start-post {
    grid-area: startpost;
}
</style>