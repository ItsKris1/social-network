<template>
    <router-view></router-view>
</template>

<script>
  // import * as _Offline from  '../public/offline.js'
export default {
  name: 'App',
  methods:{
    async checkOnlineStatus (){
      fetch("https://catfact.ninja/fact")
      .then((res) => res.json())
      .then(() => {
        let online = this.$store.state.online
        if (!online){
          this.$store.commit('updateOnlineStatus', true)
          this.$toast.success("You are connected!");
        }
      })
      .catch(()=> {
        // console.error('Oh, boy! You have lost your connection!')
        let online = this.$store.state.online
        if (online){
          this.$store.commit('updateOnlineStatus', false)
          this.$toast.warning("Oh, boy! You have lost your connection!");
        }
      })
    }
  },
  mounted(){

  setInterval(async () => {
    this.checkOnlineStatus()
  }, 1000);

  

//   window.addEventListener("load", () => {
//     if (!navigator.onLine){
//       this.$toast.warning("Oh, boy! You have lost your connection!");
//     }

// });
//   window.addEventListener("offline", () => {
//     this.$toast.warning("Oh, boy! You have lost your connection!");
  
// });

// window.addEventListener("online", () => {
//   this.$toast.success("You are connected!");
// });
  }
}
</script>

<style>
body, #app{
  width:100vw;
  height:100vh;
}
</style>
