<template>
    <div class="login-container">
        <SideBar :auth="false"/>
        <LogIn/>
    </div>
    
</template>

<script>
import LogIn from '@/components/LogIn.vue';
import SideBar from '@/components/SideBar.vue'
export default {
    name: 'LoginView',
    components:{
        LogIn,
        SideBar,
    }  
}
</script>

<style>
.login-container{
  display:flex;
  width:100%;
  height:100%;
}
</style>
