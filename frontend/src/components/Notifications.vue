<template>
    <!-- <div @click="toggle" id="notificationsBtn">
        <img src="../assets/bell.png" alt="bell.png">
        <span id="notifications_text">Notifications</span>
        <div v-if="this.$store.state.isNotificationsOpen" id="notificationsContent">
            <div>1</div>
            <div>2</div>
            <div>3</div>
        </div>
    </div> -->

    Notifications
    <div class="some-wrapper">
        <div class="item-list" id="requests">
            <ul class="item-list__items">
                <li class="item-list__item">

                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">

                        <div class="item-list__text">
                            James Brown sent you a following request
                        </div>

                    </div>

                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>
                <!-- <li class="item-list__item">

                    <div class="row1">
                        <img class="small" src="../assets/icons/arrow-circle-right.svg">
                        <div class="item-list__text">twoasfasfasfasf</div>

                    </div>

                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li> -->

            </ul>
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
.some-wrapper {
    position: relative;
}

#requests {
    position: absolute;
    min-width: 300px;
    transform: translateX(-50%);
    left: 50%;
}

.item-list__text {
    overflow-wrap: anywhere;
}

#requests ul {
    /* position: absolute; */
    background-color: orange;
}

#requests .item-list__item {
    justify-content: space-between;
}

#requests .item-list__item .row1 {
    display: flex;
    gap: 5px;
    align-items: center;
}

#requests .item-list__item .row2 {
    flex-shrink: 0;
}

#notificationsBtn {
    display: inline-block;
    position: relative;
    cursor: pointer;
    user-select: none;
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

.accept {
    color: green;
}

.decline {
    color: red;
}
</style>