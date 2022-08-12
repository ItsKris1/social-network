<template>

    <p class="notificationMsg">
        <!-- who sent -->
        <span class="who">{{ notification.user.nickname }}</span>

        <!-- additional text from database -->
        <span>{{ notification.content }}</span>

        <!-- additional event or group name -->
        <span v-if="notification.type === 'EVENT'">
            <span class="who">{{ notification.event.title }}</span> in group
            <span class="who">{{ notification.group.name }}</span>
        </span>
        <span v-if="notification.type === 'GROUP_INVITE'" class="who">{{ notification.group.name }}</span>
    </p>

</template>
<script>
export default {
    props: ['notification']
}
</script>


<style scoped>
.who {
    white-space: nowrap;
}
</style>