<template>

    <div class="item-list__wrapper" id="groups">
        <h3>Groups</h3>
        <ul class="item-list">
            <li v-for="group in allGroups" v-if="allGroups !== null && allGroups.length > 0">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">{{ group.name }}</div>
            </li>

            <p class="additional-info" v-else>You are not part of any group</p>

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