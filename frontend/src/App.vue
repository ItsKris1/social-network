<template>
    <router-view></router-view>
    <Chat />
</template>

<script>

import Chat from './components/Chat.vue'
export default {
    name: 'App',
    components: { Chat }
}
</script>

<style>
</style>