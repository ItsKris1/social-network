<template>
    <router-view></router-view>
</template>

<script>

export default {
  name: 'App',
  mounted(){
    window.addEventListener("online", () => {
      this.$store.commit('updateOnlineStatus', true)
      this.$toast.success("You are connected!");
    });

    window.addEventListener("offline", () => {
      this.$store.commit('updateOnlineStatus', false)
      this.$toast.warning("Oh, boy! You have lost your connection!");
    });
  }
}
</script>

<style>
body, #app{
  width:100vw;
  height:100vh;
}
</style>
