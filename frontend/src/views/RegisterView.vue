<template>
<NavBarOff />
    <Register />
</template>

<script>
import Register from "@/components/Register.vue";
import NavBarOff from "@/components/NavBarOff.vue";
export default {
    name: 'RegisterView',
    components: { Register, NavBarOff, },
}
</script>