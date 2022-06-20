<template>
    <div @click="toggle" id="notificationsBtn">
        <!-- <img src="../assets/bell.png" alt="bell.png"> -->
        <span id="notifications_text">Notifications</span>
        <div v-if="this.$store.state.isNotificationsOpen" id="notificationsContent">
            <div>1</div>
            <div>2</div>
            <div>3</div>
        </div>
    </div>

</template>

<script>
export default {
    name: "notifications",
    methods: {
        toggle() {
            this.$store.commit("CHANGE_NOTIFICATIONS_STATE")
            console.log(this.$store.state.isNotificationsOpen)
        },
    },
}

</script>

<style>
#notificationsBtn {
    display: inline-block;
    position: relative;
    cursor: pointer;
    user-select: none;
}

#notifications_text {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 400;
    font-size: 16px;
    line-height: 146.02%;
    /* or 23px */

    display: flex;
    align-items: center;
    text-align: center;
    color: #E4E3E3;
}

#notificationsContent {
    cursor: default;
    display: block;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
    padding: 12px 16px;
    z-index: 1;
}
</style>