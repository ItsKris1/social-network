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


<style scoped>
#navbar {
    display: flex;
    /* justify-content: space-around; */
    margin-top: 10px;
    margin-bottom: 10px;
    align-items: center;
}

#logo {
    font-family: Inter;
    font-size: 24px;
    font-weight: 649;
    line-height: 29px;
    letter-spacing: 0em;
    text-align: center;
    font-variation-settings: 'slnt' 0;
    text-decoration: none;

}
</style>