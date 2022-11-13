<template>
    <div class="sidebar">
        <h2>SN</h2>
        <div v-if="auth" id="logout">Log Out</div>
    </div>
</template>

<script>
export default{
    props:['auth'],
   computed:{
    isLoggedIn(){
        return false
    }
   }

}
</script>

<style scoped>
    .sidebar{
        width:75px;
        background-color: var(--color-blue);
        display:flex;
        flex-direction:column;
        color:var(--color-white);
        justify-content:space-between;
        align-items:center;
        text-align: center;
        padding: 15px 0 15px 0;
    }
    #logout{
        margin: 0  10px 0 10px;
        font-weight: 300;
        cursor: pointer;
        position:relative;
    }
    #logout::after{
    content: "";
    height: 2.5px;
    width: 0;
    background-color: rgb(132, 148, 236);
    display: block;
    position: absolute;
    left:10px;
    transition: all 0.35s ease-out;
    }
    #logout:hover::after{
        width: 60%;
        background-color: rgb(132, 148, 236);
    }
</style>