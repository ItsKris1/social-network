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

    <div class="relative-wrapper">
        Notifications
        <div class="item-list__wrapper" id="notifications">
            <ul class="item-list">
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">

                        <div>
                            <span class="who">John Smith</span> sent you a follow request
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">
                        <div>
                            <span class="who">Chris Brown</span> wants to join your group
                        </div>

                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/default-profile.svg">

                        <div>
                            <span class="who">Goblins</span> invited you to join their group
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>
                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/users-alt.svg">

                        <div>
                            <span class="who">ItsKris</span> did something
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>

                <li>
                    <div class="row1">
                        <img class="small" src="../assets/icons/users-alt.svg">

                        <div>
                            <span class="who">ItsKris</span> did something very longworthy notification!
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times accept"></i>
                        <i class="uil uil-check decline"></i>
                    </div>
                </li>

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
.relative-wrapper {
    position: relative;

}

#notifications {

    position: absolute;
    transform: translateX(-50%);
    left: 50%;
    font-weight: 400;
    margin-top: 5px;
    min-width: 400px;

    /* display: none; */
}



#notifications .item-list {
    gap: 15px;
}

#notifications .item-list li {
    justify-content: space-between;
    gap: 20px;
}


.who {
    font-weight: 500;
}
</style>