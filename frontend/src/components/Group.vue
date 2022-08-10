<template>

    <div class="content" v-if="groupData">

        <div class="left-section">
            <GroupMembers />
            <GroupEvents />
        </div>

        <div class="middle-section">
            <!-- Group about -->
            <div class="about">
                <h2 class="about-title">{{ this.groupData.name }}</h2>
                <p class="about-text">{{ this.groupData.description }}</p>
            </div>

            <NewPost></NewPost>
            <GroupPosts></GroupPosts>

        </div>

        <div v-if="this.groupData.admin" class="right-section">
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

    </div>
</template>


<script>
import AllPosts from './AllPosts.vue'
import Groups from './Groups.vue';
import Notifications from './Notifications.vue';
import NewPost from './NewPost.vue';
import GroupPosts from './GroupPosts.vue';
import GroupMembers from './GroupMembers.vue';
import Modal from './Modal.vue';
import GroupEvents from './GroupEvents.vue';
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
            groupData: null,
            isMemberOfGroup:false
        };
    },
    methods: {
        async getGroupInfo() {
            await fetch("http://localhost:8081/groupInfo?groupId=" + this.$route.params.id, {
                credentials: "include"
            })
                .then((r => r.json()))
                .then((json => {
                    // console.log("/groupInfo response", json);
                    this.groupData = json.groups[0];
                }));
        },

    },
    components: { AllPosts, Groups, Notifications, NewPost, GroupPosts, GroupMembers, Modal, GroupEvents }
}
</script>


<style>
.content {
    margin-top: 100px;

    display: grid;
    grid-template-columns: minmax(0, 1fr) minmax(min-content, 550px) minmax(min-content, 1fr);
    column-gap: 50px;
}



.middle-section {
    justify-self: center;
    display: flex;
    flex-direction: column;
    gap: 35px;

}


.left-section {
    justify-self: flex-end;
    display: flex;
    flex-direction: column;
    gap: 35px;
}

.right-section {
    justify-self: flex-start;
}


#requests .item-list {
    gap: 10px;
}

#requests .item-list li {
    justify-content: space-between;
    gap: 20px;
}

.box {
    height: 300px;
    width: 550px;
    border: 2px solid blue;
}

@media only screen and (max-width: 1176px) {

    .content {
        grid-template-columns: repeat(2, max-content);
        grid-template-rows: repeat(2, minmax(auto, max-content));
        row-gap: 35px;
        justify-content: center;
        grid-template-areas:
            "groups middle-section"
            "requests middle-section"
            "... middle-section";

    }


    .middle-section {
        grid-area: middle-section;
    }

    .left-section {
        grid-area: groups;
    }

    .right-section {
        grid-area: requests;
    }

}
</style>