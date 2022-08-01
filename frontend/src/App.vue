<template>

    <router-view>

    </router-view>

    <Chat v-if="$route.name !== 'sign-in' && $route.name !== 'register' && $route.name !== 'auth'" />

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