<template>
    <div class="messaging">
        <div class="messaging-header">
            <h3>Messaging</h3>
            <i class="uil uil-angle-down"></i>
        </div>

        <div class="messaging-content">
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
                    <img src="../assets/icons/users-alt.svg" alt="" class="small">
                    <div class="item-text">Group 1</div>
                </li>
                <li>
                    <img src="../assets/icons/users-alt.svg" alt="" class="small">
                    <div class="item-text">Group 1</div>
                </li>
                <li>
                    <img src="../assets/icons/users-alt.svg" alt="" class="small">
                    <div class="item-text">Group 1</div>
                </li>

            </ul>
        </div>

    </div>
</template>


<script>
export default {
    name: '',
}
</script>


<style scoped>
.messaging {
    /* border: 2px solid blue; */
    width: 315px;
    position: fixed;
    bottom: 0;
    right: 0;
}

.messaging-header {
    display: flex;
    justify-content: space-between;
    padding: 15px 20px;
    background-color: var(--color-blue);
    border-top-left-radius: 5px;
}

.messaging-header i {
    color: var(--color-white);
    font-size: 18px;
}

.messaging-header h3 {
    color: var(--color-white);
    font-weight: 300;
    letter-spacing: 0.35px;

}

.messaging .item-list {
    /* padding: 0 20px;
    padding-top: 20px; */
}

.messaging-content {
    padding: 20px;
}

.messaging .item-list:first-of-type {
    padding-bottom: 20px;
}
</style>