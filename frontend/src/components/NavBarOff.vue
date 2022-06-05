<template>

    <div id="navbar">
        <a href="/" id="logo">Social Network</a>
    </div>

</template>


<script>

export default {
    name: 'NavBarOff'
}

</script>


<style>
#navbar {
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px 50px;

    width: 1440px;
    height: 72px;

    background: #CCC9C9;
}

#logo {
    position: absolute;
    left: 9.59%;
    right: 10.05%;
    /* top: 0%;
    bottom: 0%; */

    font-family: 'Inter';
    font-style: normal;
    /* font-weight: 649; */
    font-size: 24px;
    line-height: 29px;
    
    /* identical to box height */

    /* text-align: center; */
    color: #000000;
    font-variation-settings: 'slnt' 0;

}
</style>