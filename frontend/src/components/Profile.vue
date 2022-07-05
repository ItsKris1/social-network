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

                <div class="user-profile__privacy">
                    <!-- <label for="user-profile__setting">Private profile</label>
            <input type="radio" name="user-profile__setting" id="user-profile__setting"> -->

                    <span>Private profile</span>
                    <img src="../assets/Toggle.png">
                </div>

                <!-- <button class="btn">Follow<i class="uil uil-user-plus"></i></button> -->
            </div>
            <div class="multiple-item-list" id="following">
                <div class="item-list__wrapper">
                    <h3>Followers</h3>
                    <ul class="item-list users">
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 1</div>
                        </li>
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 2</div>
                        </li>
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 3</div>
                        </li>
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 4</div>
                        </li>
                    </ul>
                </div>

                <div class="item-list__wrapper">
                    <h3>Following</h3>
                    <ul class="item-list users">
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 1</div>
                        </li>
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 2</div>
                        </li>
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 3</div>
                        </li>
                        <li>
                            <div class="user-picture small"></div>
                            <div class="item-text">User 4</div>
                        </li>
                    </ul>
                </div>
            </div>
        </div>


        <div class="middle-section ">
            <div class="about" v-if="user.about !== ''">
                <h2 class="about-title">About me</h2>
                <p class="about-text">{{ user.about }}</p>
            </div>
            <AllMyPosts />
        </div>







    </div>

</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
// import { mapGetters } from 'vuex'
export default {
    name: 'Profile',
    components: { AllMyPosts },
    data() {
        return {
            user: {}
        }
    },
    created() {
        // this.getUserInfo()
        this.getUserId()
    },
    computed: {
        // ...mapGetters(['userInfo']),
        // ...getUserId()
    },
    methods: {
        follow() {
            console.log('subscribe function')
        },
        getUserInfo() {
            this.$store.dispatch('getMyProfileInfo')
        },
        async getUserId() {
            await fetch("http://localhost:8081/userData?userId=" + this.$route.params.id, {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json) => {
                    this.user = json.users[0];
                    // console.log(userInfo);
                    // this.commit("updateProfileInfo", userInfo);
                    console.log("user profile info -", json);
                });
        }
    },
    watch: { //watching changes in route
        $route() {
            this.getUserId()
        }
    }
}
</script>

<style>
#layout-profile {
    display: grid;
    grid-template-columns: 1fr minmax(min-content, 500px) 1fr;
    column-gap: 50px;
    margin-top: 100px;

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
    align-items: flex-end;
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
    margin-left: 15px;
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

    width: 500px;
    background: var(--color-white);
    box-shadow: var(--container-shadow);
    border-radius: var(--container-border-radius);

}
</style>