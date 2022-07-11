<template>

    <div id="layout-profile">

        <div class="left-section ">
            <div class="user-profile__public">
                <div class="user-picture" :style="{ backgroundImage: `url(http://localhost:8081/${user.avatar})` }">
                </div>
                <div class="user-profile__info">
                    <h3 class="username">{{ user.nickname }}</h3>
                    <p class="user-email">{{ user.login }}</p>
                    <p class="user-dateOfBirth">{{ user.dateOfBirth }}</p>
                </div>

                <PrivacyBtn v-if="isMyProfile" />
                <FollowBtn v-else />


            </div>
            <div class="multiple-item-list">
                <Following />
                <Followers />
            </div>
        </div>


        <div class="middle-section ">
            <div class="about" v-if="user.about !== ''">
                <h2 class="about-title">About me</h2>
                <p class="about-text">{{ user.about }}</p>
            </div>
            <AllMyPosts v-bind:userid="this.user.id" />
        </div>

    </div>
</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
import Following from './Following.vue'
import Followers from './Followers.vue'
import FollowBtn from './FollowBtn.vue'
import PrivacyBtn from './PrivacyBtn.vue'
// import { mapGetters } from 'vuex'
export default {
    name: 'Profile',
    components: { AllMyPosts, Followers, Following, FollowBtn, PrivacyBtn },
    data() {
        return {
            user: {},
            isMyProfile: false,
        }
    },
    created() {
        // this.getUserInfo()
        this.getUserId()
        this.checkProfile()
    },
    computed: {
        // ...mapGetters(['userInfo']),
        // ...getUserId()
    },

    methods: {
        getUserInfo() {
            this.$store.dispatch('getMyProfileInfo')
        },
        async getUserId() {
            await fetch("http://localhost:8081/userData?userId=" + this.$route.params.id, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json) => {
                    // console.log("profile.vue/getuserid",json);
                    this.user = json.users[0];
                    // console.log("user", this.user)

                    // if (this.$route.params.id === this.user.id) {
                    //     console.log("my user")
                    // }
                    // console.log(userInfo);
                    // this.commit("updateProfileInfo", userInfo);
                    // console.log("user profile info -", json);
                });
            // this.isOwnerProfile()
        },
        // isOwnerProfile() {
        //     console.log("cookie",document.cookie);
        //     console.log("user id",this.user.id);
        //     let activeCookie = document.cookie.slice(11)
        //     if (activeCookie === this.user.id) {
        //         console.log("It's a owner");
        //     } else { console.log("It's NOT a owner") }

        // }

        async getLoggedUserId() {
            const response = await fetch("http://localhost:8081/currentUser", {
                credentials: "include",
            })

            const data = await response.json();
            this.loggedUserID = data.users[0].id;

            return data.users[0].id

            // .then((r) => r.json())
            // .then((json => {
            //     this.loggedUserID = json.users[0].id
            // }))
        },

        async checkProfile() {
            const profileID = this.$route.params.id;
            const loggedUserID = await this.getLoggedUserId();

            console.log(profileID)
            console.log(loggedUserID)

            this.isMyProfile = (profileID === loggedUserID)
        }
    },
    watch: { //watching changes in route
        $route() {
            this.getUserId()
            this.checkProfile();
        }
    }
}
</script>

<style>
#layout-profile {
    display: grid;
    grid-template-columns: 1fr minmax(min-content, 550px) 1fr;
    column-gap: 50px;
    margin-top: 100px;
    justify-items: flex-end;

}

.middle-section {
    display: flex;
    flex-direction: column;
    gap: 50px;
}

.left-section {
    display: flex;
    flex-direction: column;
    gap: 50px;
    max-width: max-content;
    min-width: min-content;

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



.user-profile__privacy {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 7.5px;

    padding-top: 25px;
    border-top: 1px solid rgb(212, 212, 212);
    width: 100%;
    font-size: 14px;
}

.about {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    padding: var(--container-padding);
    gap: 15px;

    background: var(--color-white);
    box-shadow: var(--container-shadow);
    border-radius: var(--container-border-radius);

}
</style>