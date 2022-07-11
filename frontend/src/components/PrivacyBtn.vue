<template>
    {{ isMyProfile }}

    <div v-if="isMyProfile">
        <span>Private profile</span>
        <div class="hello-world">
            <div class="toggle-wrapper">
                <span class="toggle-indicator"></span>
            </div>
        </div>


    </div>


</template>


<script>
export default {
    name: 'PrivacyBtn',
    props: ['profileID'],

    data() {
        return {
            isMyProfile: false
        }
    },

    created() {
        this.getLoggedUserId()
    },
    methods: {
        async getLoggedUserId() {
            await fetch("http://localhost:8081/currentUser", {
                credentials: "include",
            })
                .then((r) => r.json())
                .then((json => {
                    const userid = json.users[0].id
                    this.isMyProfile = userid === this.profileID;
                }))
        },
    }

}


</script>


<style>
.toggle-wrapper {
    width: 60px;
    height: 30px;
    /* border: 1px solid black; */
    box-shadow: inset 0 0 0 1px black;
    border-radius: 50px;
    display: inline-block;
    position: relative;
}

.toggle-indicator {
    display: inline-block;
    height: 30px;
    width: 30px;
    background-color: black;
    border-radius: 50px;
    position: absolute;
    /* left: 0; */
    /* top: 0; */
    left: 0;
    transition: transform 0.4s ease;
}

/* .toggle-indicator:hover {
    transform: translateX(100%);
} */
</style>