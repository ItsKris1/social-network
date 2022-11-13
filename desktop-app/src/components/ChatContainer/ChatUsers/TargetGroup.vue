<template>
    <li class="group">
        <img src="../../../assets/icons/users-alt.svg" alt=""/>
        <div class="item-text">{{ group.name }}</div>
    </li>
</template>

<script>

export default {
    name:"TargetGroup",
    props: ['group'],
}
</script>
<style scoped>
.group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    justify-content: flex-start;
}
</style>