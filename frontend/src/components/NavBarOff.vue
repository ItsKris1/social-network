<template>
    <div id="navbar">
        <a id="nav-title" href="/">Social Network</a>
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
    padding: 10px 40px;
    background-color: var(--color-blue);
    align-items: center;
    justify-content: space-between;
    font-weight: 300;
    color: var(--color-white);
    margin-bottom: 50px;

}

#nav-title {
    font-size: 24px;
    color: var(--color-white);
    text-decoration: none;
}
</style>