<template>
    <div class="item-list__wrapper" id="following">
        <h3>Following</h3>
        <ul class="item-list users">
            <li>
                <div class="user-picture small"></div>
                <div class="item-text">User 1</div>
            </li>
            <li>
                <div class="user-picture small"></div>
                <div class="item-text">User 2</div>
            </li>
            <li>
                <div class="user-picture small"></div>
                <div class="item-text">User 3</div>
            </li>
            <li>
                <div class="user-picture small"></div>
                <div class="item-text">User 4</div>
            </li>
        </ul>
    </div>
</template>


<script>
export default {
    name: 'Following',
    methods: {
        async getFollowing() {
            console.log("getFollowing");
            await fetch('http://localhost:8081/following?userId=' + this.$route.params.id, {
                credentials: 'include'
            })
                .then((response => response.json()))
                .then((json => console.log(json)))
        }
    }
}
</script>


<style>
</style>