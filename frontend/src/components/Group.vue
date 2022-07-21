<template>

    <!-- READ ME -->
    <!-- I am using other components just for designing the layout. Later we can replace those with the right components -->
    <div class="content">

        <!-- Start post -->
        <NewPost></NewPost>
        <!-- Group about -->
        <div class="about">
            <h2 class="about-title">Group name</h2>
            <p class="about-text">Group description</p>
        </div>

        <!-- Group posts -->
        <AllPosts></AllPosts>

        <!-- Members -->
        <Groups></Groups>

        <!-- Events -->
        <Groups></Groups>

        <!-- Group requests -->
        <div class="item-list__wrapper" id="requests">
            <h3>Requests</h3>
            <ul class="item-list">
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">

                        <div>
                            <span class="who">John Smith</span>
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">
                        <div>
                            <span class="who">Chris Brown</span>
                        </div>

                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">

                        <div>
                            <span class="who">Goblins</span>
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>

            </ul>
        </div>
    </div>

</template>


<script>
import AllPosts from './AllPosts.vue'
import Groups from './Groups.vue';
import Notifications from './Notifications.vue';
import NewPost from './NewPost.vue';
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
            groupData: null
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
    },
    components: { AllPosts, Groups, Notifications, NewPost }
}
</script>


<style>
.content {
    margin-top: 100px;
}

#requests .item-list {
    gap: 15px;
}

#requests .item-list li {
    justify-content: space-between;
    gap: 20px;
}
</style>