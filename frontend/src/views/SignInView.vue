<template>
    <NavBarOff />
    <SignIn />

</template>

<script>

// @ is an alias to /src
import NavBarOff from '@/components/NavBarOff.vue';
import SignIn from '@/components/SignIn.vue';

export default {
    name: 'SignInView',
    components: {
        NavBarOff,
        SignIn,
    }
}

</script>