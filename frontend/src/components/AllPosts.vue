 <template>
    <div id="all_posts" v-if="Array.isArray(allPosts) && allPosts.length > 0">
        <Post v-for="post in allPosts" :key="post.id" v-bind:postData="post" />
    </div>

    <p v-if="!allPosts">No posts</p>

</template>

<script>
import Post from './Post.vue'
import { mapGetters } from 'vuex'
export default {
    name: 'AllPosts',
    created() {
        this.$store.dispatch('fetchPosts')
    },


    computed: mapGetters(['allPosts']),
    components: { Post },
}
</script>

<style>
#all_posts {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 50px;
}
</style>