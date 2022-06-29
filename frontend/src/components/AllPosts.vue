 <template>
    <!-- <button @click="showAllPosts">showFetch</button> -->
    <!-- <div id="all_posts" v-if="allPosts !== undefined">
           <Post v-for="post in allPosts" :key="post.id" v-bind:postData="post" />
       </div> -->

    <div class="post-wrapper">
        <div class="post">
            <div class="user-picture medium"></div>
            <div class="post-content">
                <p class="post-author">James Brown</p>
                <p class="post-body">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
                    incididunt
                    ut labore et dolore magna
                    aliqua. Ut enim ad minim veniam, quis nostrud exercitation.</p>
                <!-- <img class="post-image"> -->
                <button class="btn">Comments</button>
            </div>
        </div>

        <div class="comments">
            <div class="comment">
                <div class="user-picture medium"></div>
                <div class="comment-content">
                    <p class="comment-author">Roy Jones</p>
                    <p class="comment-body">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
                        tempor
                        incididunt
                        ut labore et dolore magna
                        aliqua. Ut enim ad minim veniam, quis nostrud exercitation.</p>
                    <!-- <img class="comment-image" src="../assets/pexels-phil-desforges-12576276.jpg" alt="" srcset=""> -->
                </div>
            </div>

            <div class="comment">
                <div class="user-picture medium"></div>
                <div class="comment-content">
                    <p class="comment-author">Roy Jones</p>
                    <p class="comment-body">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
                        tempor
                        incididunt
                        ut labore et dolore magna
                        aliqua. Ut enim ad minim veniam, quis nostrud exercitation.</p>
                    <img class="comment-image" src="../assets/pexels-phil-desforges-12576276.jpg" alt="" srcset="">
                </div>
            </div>

        </div>
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
    methods: {
        // showAllPosts() {
        //     this.$store.dispatch('fetchPosts')
        //     console.log('All posts: ', this.allPosts);
        // },
    },
}
</script>

<style>
#all_posts {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>