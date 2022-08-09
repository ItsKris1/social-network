<template>
    <button class="btn" @click="handleRequest">{{ buttonText }}</button>
</template>

<script>
export default {
    data() {
        return {
            buttonText: "Request to join"
        }
    },

    methods: {
        handleRequest() {

        }
    }
}
</script>