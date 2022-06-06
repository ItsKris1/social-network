<template>
    <transition name="fade" mode="out-in">
        <router-view />
    </transition>
</template>

<script>
export default {
    name: 'App'
}
</script>

<style>
.fade-enter-active,
.fade-leave-active {
    transition: all 0.5s;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
    transform: translateY(30px)
}
</style>