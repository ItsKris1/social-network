<template>

    <div class="item-list__wrapper" id="groups">
        <h3>My groups</h3>
        <ul class="item-list">
            <li v-for="group in allGroups">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">{{ group.name }}</div>
            </li>
            <!-- <li>
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
            </li> -->

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

    created() {
        this.$store.dispatch('getAllGroups');
    },

    computed: {
        allGroups() {
            return this.$store.state.groups.allGroups
        }
    }
}
</script>


<style>
/* .item-list__wrapper#groups {
    align-items: flex-start;
} */
</style>