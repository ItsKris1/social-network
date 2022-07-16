<template>
    <div class="temporary">
        <div>
            <div>{{ this.groupData.name }}</div>
            <div>{{ this.groupData.description }}</div>
        </div>
        <NewPost />
        <GroupPosts />
    </div>
</template>


<script>
import NewPost from './NewPost.vue';
import GroupPosts from './GroupPosts.vue';
export default {
    name: "Group",
    created() {
        this.getGroupInfo();
        // this.getGroupPosts();
    },
    watch: {
        $route() {
            this.getGroupInfo();
        }
    },
    data() {
        return {
            groupData: null,
            // groupPosts: []
        };
    },
    methods: {
        async getGroupInfo() {
            await fetch("http://localhost:8081/groupInfo?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((r => r.json()))
                .then((json => {
                    console.log(json);
                    this.groupData = json.groups[0];
                }));
        },
        // async getGroupPosts() {
        //     await fetch("http://localhost:8081/groupPosts?groupId=" + this.$route.params.id, {
        //         credentials: "include"
        //     })
        //     .then((r=>r.json()))
        //     .then((json=>console.log(json)))

        // }

    },
    components: { NewPost, GroupPosts }
}
</script>


<style>
.temporary {
    margin-top: 100px;
}
</style>