<template>

    <div class="relative-wrapper" @click="toggleShowNotifications">
        <span class="link-header">Notifications</span>
        <div class="item-list__wrapper" id="notifications" v-show="showNotifications">
            <ul class="item-list">
                <li>
                    <div class="row1 ">
                        <img class="" src="../assets/icons/default-profile.svg">

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
                        <img class="" src="../assets/icons/default-profile.svg">
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
                        <img class="" src="../assets/icons/default-profile.svg">

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
                        <img class="" src="../assets/icons/users-alt.svg">

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
                        <img class="" src="../assets/icons/users-alt.svg">

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
    data() {
        return {
            showNotifications: false
        }
    },

    methods: {
        toggleShowNotifications() {
            this.showNotifications = !this.showNotifications
        }
    },

}

</script>

<style>
.relative-wrapper {
    position: relative;

}


#notifications .row1 :is(img, i) {
    height: 2.25em;
    width: 2.25em;
}

#notifications {

    position: absolute;
    transform: translateX(-50%);
    left: 50%;
    font-weight: 400;
    margin-top: 10px;
    /* min-width: 400px; */
    /* max-width: 450px; */
    /* right: 5px; */
    /* top: 50px; */
    /* display: none; */
    width: 400px;
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


.link-header::after {
    content: "";
    height: 2.5px;
    width: 0;
    display: block;

    position: absolute;

    transition: all 0.35s ease-in-out;
}

.link-header:hover::after {
    width: 100%;
    background-color: rgb(132, 148, 236);
}
</style>