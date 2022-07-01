<template>
    <div class="item-list">
        <h3>My groups</h3>
        <ul class="item-list__items">
            <li class="item-list__item">
                <i class="uil uil-users-alt"></i>
                <div class="item-list__text">Group 1</div>
            </li>
            <li class="item-list__item">
                <i class="uil uil-users-alt"></i>
                <div class="item-list__text">Group 2</div>
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