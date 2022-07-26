<template>

    <div class="messaging-wrapper">
        <ChatBox></ChatBox>

        <div class="messaging" @click="toggleShowContent">

            <div class="messaging-header">
                <p>Messaging</p>
                <i class="uil uil-angle-down"></i>
            </div>

            <div class="messaging-content" v-show="showContent">
                <ul class="item-list">
                    <li>
                        <div class="user-picture small"></div>
                        <div class="item-text">User 1</div>
                    </li>
                    <li>
                        <div class="user-picture small"></div>
                        <div class="item-text">User 2</div>
                    </li>
                    <li>
                        <div class="user-picture small"></div>
                        <div class="item-text">User 3</div>
                    </li>

                </ul>

                <ul class="item-list">
                    <li>
                        <img src="../../assets/icons/users-alt.svg" alt="" class="small">
                        <div class="item-text">Group 1</div>
                    </li>
                    <li>
                        <img src="../../assets/icons/users-alt.svg" alt="" class="small">
                        <div class="item-text">Group 2</div>
                    </li>
                    <li>
                        <img src="../../assets/icons/users-alt.svg" alt="" class="small">
                        <div class="item-text">Group 3</div>
                    </li>
                </ul>
            </div>

        </div>
    </div>

</template>


<script>
import ChatBox from './ChatBox.vue'
export default {
    name: '',
    components: { ChatBox },
    data() {
        return {
            showContent: false,
        }
    },

    methods: {
        toggleShowContent() {
            this.showContent = !this.showContent
        }
    }
}
</script>


<style scoped>
.messaging-wrapper {
    position: fixed;
    bottom: 0;
    right: 0;
    display: flex;
    align-items: flex-end;
}

.messaging {

    width: 315px;
    background-color: var(--color-white);
}

.messaging-header {
    display: flex;
    justify-content: space-between;
    padding: 12px 20px;
    background-color: var(--color-blue);
    border-top-left-radius: 5px;
    color: var(--color-white);
}

.messaging-header i {
    font-size: 18px;
}

.messaging-header p {
    font-weight: 300;
    letter-spacing: 0.35px;
}

.messaging-content {
    padding: 20px;
}

.messaging .item-list:first-of-type {
    padding-bottom: 20px;
}
</style>