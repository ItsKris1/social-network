<template>

    <div class="relative-wrapper" @click="toggleShowNotifications">
        <span class="link-header">Notifications</span>
        <div class="item-list__wrapper" id="notifications" v-show="showNotifications">
            <ul class="item-list">
                <li v-for="notification in allNotifications"
                    v-if="
                    isDataValid(notificationsFromDB) &&
                    allNotifications.length > 0">

                    <div class="row1 ">
                        <img class="" src="../assets/icons/default-profile.svg">

                        <div>
                            <span class="who">{{ notification.user.nickname }}</span> {{ notification.content }}
                        </div>
                    </div>
                    <div class="row2">
                        <i class="uil uil-times decline" @click.stop="handleRequest(notification, 'decline')"></i>
                        <i class="uil uil-check accept" @click.stop="handleRequest(notification, 'accept')"></i>
                    </div>
                </li>

                <li v-else class="additional-info">No notifications</li>

                <!-- <li>
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
                </li> -->

            </ul>
        </div>
    </div>

</template>

<script>
import { mapState } from 'vuex';

export default {
    name: "notifications",
    data() {
        return {
            showNotifications: false,
            notificationsFromDB: {},
            // allNotifications: [],
        }
    },

    async created() {
        await this.fetchNotifications();
        // await this.initAllNotifications();
    },

    computed: {
        // allNotifications: {

        //     get() {
        //         return [...this.notificationsFromDB.notifications, ...this.$store.state.notifications.newNotifications]
        //     },

        //     set(newValue) {

        //     }
        // }

        ...mapState({
            allNotifications: state => state.notifications.allNotifications
        })
    },

    unmounted() {
        this.$store.commit("updateAllNotifications", [])
    },

    methods: {
        toggleShowNotifications() {
            this.showNotifications = !this.showNotifications
        },

        async fetchNotifications() {
            const response = await fetch('http://localhost:8081/notifications', {
                credentials: 'include'
            })

            const data = await response.json();
            this.notificationsFromDB = data;

            this.$store.commit("updateAllNotifications", data.notifications)
            // console.log("/notifications data", data)
        },

        async handleRequest(notification, reqResponse) {
            let endpoint;

            switch (notification.type) {
                case "FOLLOW":
                    endpoint = "responseFollowRequest";
                    break;
                case "GROUP_INVITE":
                    endpoint = "responseInviteRequest";
                    break;
            }

            const response = await fetch(`http://localhost:8081/${endpoint}`, {
                credentials: 'include',
                method: 'POST',
                body: JSON.stringify({
                    requestId: notification.id,
                    response: reqResponse,
                })
            })

            const data = await response.json();
            console.log("/acceptRequest data", data)

            // remove the notification
            // this.removeNotification(notification.id)
            this.$store.dispatch("removeNotification", notification.id)

        },


        isDataValid(resp) {
            return resp.type === "Success" ? true : false;
        },

        // initAllNotifications() {
        //     const notificationsFromDB = this.notificationsFromDB;

        // }
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