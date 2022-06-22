<template>
    <div id="profile">
        <div>
            <img id="profileImg" :src="'http://localhost:8081/'+this.user.avatar" alt="profilePic">
        </div>
        <div id="basicInformation">
            <span>{{this.user.nickname}}</span>
            <span>e-mail</span>
            <span>birthday date</span>
            <button id="followBtn" click="follow">Follow</button>
        </div>
        <div id="profileSettings">
            <span>Privacy change-button</span>
        </div>
    </div>
    <AllMyPosts />
</template>

<script>
import AllMyPosts from './AllMyPosts.vue'
export default {
    name: 'Profile',
    components:{AllMyPosts},
    data(){
        return{
            user:{
                avatar:"",
                id:"",
                nickname:""
            }
        }
    },
    created(){
        this.fetchBaseInfo()
    },
    methods: {
        follow() {
            console.log('subscribe function')
        },
        async fetchBaseInfo(){
            await fetch("http://localhost:8081/currentUser", {
                credentials: 'include',
            })
                .then((r => r.json()))
                .then((json => {
                    // console.log(json)
                    this.user = json.users[0]
                }))
        },
    }
}
</script>

<style>
#profile {
    display: flex;
    margin-top: 10px;
}

#profileImg {
    border-radius: 50%;
    margin-right: 20px;
    width: 200px;
    height: 200px;
}

#basicInformation {
    display: grid;
}

#profileSettings{
    margin-left: auto;
}

#followBtn {
    height: 37px;
    width: 92px;
    left: 0px;
    top: 0px;
    border-radius: 8px;
    padding: 10px, 15px, 10px, 15px;
}
</style>