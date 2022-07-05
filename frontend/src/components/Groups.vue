<template>


    <div class="item-list__wrapper" id="groups">
        <h3>My groups</h3>
        <ul class="item-list">
            <li>
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">Group 1</div>
            </li>
            <li>
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">Group 2</div>
            </li>
            <li>
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">Group 3</div>
            </li>
            <li>
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">Group 4</div>
            </li>
        </ul>
        <button class="btn">New group<i class="uil uil-plus"></i></button>
    </div>

</template>


<script>
export default {
    name: 'Groups',
    methods: {
        createNewGroup() {

        }
    }
}
</script>


<style>
</style>