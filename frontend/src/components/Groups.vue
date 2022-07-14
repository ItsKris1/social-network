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
        <NewGroup />
    </div>

</template>


<script>
import NewGroup from '@/components/NewGroup.vue'

export default {
    name: 'Groups',
    components: {
        NewGroup
    },

    methods: {
        createNewGroup() {

        }
    }
}
</script>


<style>
</style>