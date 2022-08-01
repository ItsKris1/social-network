<template>
    <router-view></router-view>

    <Chat v-if="$route.name !== 'sign-in' && $route.name !== 'register'" />
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