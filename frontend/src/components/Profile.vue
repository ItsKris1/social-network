<template>

    <div v-if="user && this.$store.state.id !== ''">
        <div id="layout-profile">

            <div class="left-section ">
                <div class="user-profile__public">
                    <div class="user-picture" :style="{ backgroundImage: `url(http://localhost:8081/${user.avatar})` }">
                    </div>
                    <div class="user-profile__info">
                        <h3 class="username">{{ user.nickname }}</h3>
                        <p class="user-email" v-if="user.login">{{ user.login }}</p>
                        <p class="user-dateOfBirth" v-if="user.dateOfBirth">{{ user.dateOfBirth }}</p>
                    </div>

                    <div class="profile-btns">
                        <!-- Privacy and follow/unfollow button-->
                        <PrivacyBtn v-if="isMyProfile" :status="user.status" />
                        <FollowBtn v-else-if="!user.following" @follow="toggleFollowingThisUser" />
                        <UnfollowBtn v-else @unfollow="toggleFollowingThisUser" />
                        <!-- Send message button -->
                        <button v-if="showSendButton"
                                class="btn">Send message
                            <i class="uil uil-message"></i></button>
                    </div>

                </div>
                <div class="multiple-item-list" v-if="showProfileData">
                    <Following />
                    <Followers />
                </div>

            </div>

            <div class="middle-section" v-if="showProfileData">

                <div class="about" v-if="user.about !== ''">
                    <h2 class="about-title">About me</h2>
                    <p class="about-text">{{ user.about }}</p>
                </div>
                <AllMyPosts v-bind:userid="this.user.id" />

            </div>

            <p v-else class="additional-info large"> This profile is private</p>

        </div>
    </div>

</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
import Following from './Following.vue'
import Followers from './Followers.vue'
import FollowBtn from './FollowBtn.vue'
import PrivacyBtn from './PrivacyBtn.vue'
import UnfollowBtn from './UnfollowBtn.vue'
// import { mapGetters } from 'vuex'
export default {
    name: 'Profile',
    components: { AllMyPosts, Followers, Following, FollowBtn, PrivacyBtn, UnfollowBtn },
    data() {
        return {
            user: null,
            isMyProfile: false,
        }
    },
    created() {
        this.getUserData()
        this.checkProfile()
    },
    computed: {
        showProfileData() {
            return (this.user.following || this.isMyProfile || this.user.status === "PUBLIC") ? true : false
        },
        showSendButton() {
            return !this.isMyProfile && this.user.status === "PUBLIC" && !this.user.following
        }
    },



    methods: {
        async getUserData() {
            await fetch("http://localhost:8081/userData?userId=" + this.$route.params.id, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json) => {
                    this.user = json.users[0];
                    // console.log("User", this.user)

                });

        },

        async getMyUserID() {
            await this.$store.dispatch("getMyUserID")
        },

        async checkProfile() {
            await this.getMyUserID();
            const myID = this.$store.state.id;
            const profileID = this.$route.params.id;
            this.isMyProfile = (profileID === myID)
        },


        toggleFollowingThisUser() {
            this.getUserData();
            this.user.following = !this.user.following
        }
    },
    watch: { //watching changes in route
        $route() {
            if (this.$route.name === "Profile") {
                this.getUserData();
                this.checkProfile();
            }

        }
    }
}
</script>

<style>
#layout-profile {
    display: grid;
    grid-template-columns: 1fr minmax(min-content, 550px) 1fr;
    column-gap: 50px;
    margin-top: 50px;
    justify-items: center;

}

.middle-section {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 50px;
}

.left-section {
    display: flex;
    flex-direction: column;
    gap: 50px;
    max-width: max-content;
    min-width: min-content;
    justify-self: flex-end;

}


.user-profile__public,
.user-profile__private {
    display: flex;
    flex-direction: column;
    padding: var(--container-padding);
    background-color: var(--color-white);
    box-shadow: var(--container-shadow);
    border-radius: var(--container-border-radius);
    align-items: center;
    text-align: center;
    gap: 25px;


}

:is(.user-profile__public, .user-profile__private) .user-picture {
    background-image: url(../assets/pexels-jack-winbow-1559486.jpg);
    height: 185px;
    width: 185px;
}

.user-profile__info {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.profile-btns {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;
}
</style>