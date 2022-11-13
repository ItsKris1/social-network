<template>
    <li class="user">
        <div class="user-picture"></div>
        <div class="item-text">{{ user.nickname }}</div>
    </li>
</template>

<script>

export default {
    name:"TargetUser",
    props: ['user'],
}
</script>
<style scoped>

.user {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    justify-content: flex-start;
}
</style>
