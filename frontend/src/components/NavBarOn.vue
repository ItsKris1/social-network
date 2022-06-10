<template>

    <div id="navbar">
        <div>
            <div>
                <a href="/" id="logo">Social Network</a>
            </div>
            <div>
                <input type="text" placeholder="Search">
                <img>
            </div>
        </div>
        <span>Home</span>
        <div>
            <router-link to="/profile">
                <img alt="profilePic">
            </router-link>
            <span>James Brown</span>
        </div>
        <div>
        </div>
    </div>

</template>


<script>

export default {
    name: 'NavBarOn'
}

</script>


<style>
#navbar {
    display: flex;
}
</style>