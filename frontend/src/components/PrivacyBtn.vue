<template>
    <div class="user-profile__privacy">
        <span>Private profile</span>

        <div class="toggle-wrapper" @click="private = !private" :class="{ active: private }">
            <span class="toggle-indicator"></span>
        </div>

    </div>


</template>


<script>
export default {
    data() {
        return {
            private: false
        }
    }
}

// export default {
//     name: 'PrivacyBtn',
//     props: ['profileId'],

//     data() {
//         return {
//             isMyProfile: false,
//             loggedUserID: ""
//         }
//     },

//     // created() {
//     //     this.checkProfile();

//     // },
//     methods: {
//         async getLoggedUserId() {
//             const response = await fetch("http://localhost:8081/currentUser", {
//                 credentials: "include",
//             })

//             const data = await response.json();
//             this.loggedUserID = data.users[0].id;

//             // .then((r) => r.json())
//             // .then((json => {
//             //     this.loggedUserID = json.users[0].id
//             // }))
//         },

//         async checkProfile() {
//             await this.getLoggedUserId();
//             this.isMyProfile = (this.loggedUserID === this.profileId)
//             console.log("1", this.loggedUserID)
//             console.log("2", this.profileId)
//         },

//     }

// }


</script>


<style>
.toggle-wrapper {
    width: 60px;
    height: 30px;
    /* border: 1px solid black; */
    box-shadow: 0 0 0 1px rgb(221, 221, 221);
    border-radius: 50px;
    position: relative;
    background-color: rgb(228, 228, 228);
}

.toggle-indicator {
    display: inline-block;
    height: 30px;
    width: 30px;
    background-color: rgb(184, 184, 184);
    border-radius: 50px;
    position: absolute;
    /* left: 0; */
    /* top: 0; */
    left: 0;
    transition: all 0.3s ease;
}

.active .toggle-indicator {
    transform: translateX(100%);
    background-color: rgb(0, 158, 0);
}



/* .toggle-indicator:hover {
    transform: translateX(100%);
} */
</style>