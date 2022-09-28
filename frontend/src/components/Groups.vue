<template>
    <div class="item-list__wrapper" id="groups" v-if="dataLoaded">

        <h3>My groups</h3>
        <!-- <div @click="show">HHHH</div> -->
        <ul class="item-list">

            <li v-for="group in userGroups" v-bind:key="group.id" v-if="this.userGroups !== null">
                <img class="small" src="../assets/icons/users-alt.svg" alt="">
                <div class="item-text"><router-link :to="{ path: `/group/${group.id}`}">{{ group.name }}
</router-link></div>
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
    // created() {
    //     this.$store.dispatch('getUserGroups')
    // },
    computed: mapState({
        userGroups: state => state.groups.userGroups,
        dataLoaded: state => state.dataLoaded.userGroups
    }),
    methods: {
        // show(){
        //     console.log(this.userGroups);
        // }
    }


}
</script>
