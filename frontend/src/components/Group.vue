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
    },
    watch: {
        $route() {
            this.getGroupInfo();
        }
    },
    data() {
        return {
            groupData: {},
        };
    },
    methods: {
        async getGroupInfo() {
            await fetch("http://localhost:8081/groupInfo?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((r => r.json()))
                .then((json => {
                    // console.log("getGroupInfo",json);
                    this.groupData = json.groups[0];
                }));
        },
    },
    components: { NewPost, GroupPosts }
}
</script>


<style>
.temporary {
    margin-top: 100px;
}
</style>