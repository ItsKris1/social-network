<template>

    <!-- <button @click="getAllPosts">get</button> -->
    <div v-if="this.posts !== undefined">
        <Post v-for="post in posts" :key="post.ID" v-bind:postData="post" />
    </div>

</template>


<script>
import Post from './Post.vue'
export default {
    name: 'AllPosts',
    data() {
        return {
            posts:

                [{
                    "ID": 1,
                    "Content": "Hello Vue",
                    "Author": {
                        "nickname": "VanDamm",
                        "first-name": "Hola",
                        "last-name": "Hup"
                    },
                    "Comments": [{
                        "id": 1,
                        "author": "Mac Donald",
                        "content": "Its time to write comments."
                    },
                    {
                        "id": 2,
                        "author": "Jackie Chan Jr",
                        "content": "KiiYeeaaa."
                    }]
                }
                    , {
                    "ID": 2,
                    "Content": "here is post body!",
                    "Author": {
                        "nickname": "",
                        "first-name": "Hell",
                        "last-name": "Angel"
                    }
                }, {
                    "ID": 3,
                    "Content": "Another post!",
                    "Author": {
                        "nickname": "jackson",
                        "first-name": "five",
                        "last-name": "clowns"
                    }
                }]



        }
    },
    components: { Post },
    methods: {
        async getAllPosts() {
            await fetch("https://d1cd8cff-f1bc-4c79-913a-405f7cbbb7c3.mock.pstmn.io/allposts")
                // .then((r => console.log(r)))
                .then((res => res.json()))
                .then((r => console.log(r)))
                .then((r => {
                    this.posts = r
                }))
        },
    },
    // computed: {
    //     async getAllPosts() {
    //         await fetch("https://d1cd8cff-f1bc-4c79-913a-405f7cbbb7c3.mock.pstmn.io/allposts")
    //             .then((res => res.json()))
    //             .then((r => {
    //                 this.posts = r
    //             }))
    //             .then((r => console.log(r)))
    //     },
    // },
    // props: [this.posts]


}
</script>


<style>
</style>