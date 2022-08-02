<template>

    <router-view>

    </router-view>

    <router-view name="Chat"></router-view>
</template>

<script>

import Chat from './components/Chat/Chat.vue'
export default {
    name: 'App',
    components: { Chat }
}
</script>

<style>
</style>