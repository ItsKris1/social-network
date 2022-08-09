<template>

    <div class="item-list__wrapper" id="groups" v-if="userGroups.type">
        <h3>Groups</h3>
        <ul class="item-list">
            <li v-for="group in userGroups.groups" v-if="userGroups.groups !== null">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text">{{ group.name }}</div>
            </li>

            <p class="additional-info" v-else>You are not part of any group</p>

        </ul>
        <NewGroup />
    </div>

</template>


<script>
import NewGroup from '@/components/NewGroup.vue';
import { mapState } from 'vuex';

export default {
    name: 'Groups',
    components: { NewGroup },
    computed: mapState({
        userGroups: state => state.groups.userGroups
    })
}
    // created() {
    //     this.$store.dispatch('getAllGroups');
    // },

</script>
