 <template>
    <!-- <button @click="showAllPosts">showFetch</button> -->
    <div id="all_posts" v-if="allPosts !== undefined">
        <Post v-for="post in allPosts" :key="post.id" v-bind:postData="post" />
    </div>

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
    methods: {},
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