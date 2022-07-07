<template>

    <span>Private profile</span>
    <div class="hello-world">
        <div class="toggle-wrapper">
            <span class="toggle-indicator"></span>
        </div>
    </div>

</template>


<script>
export default {
    name: 'PrivacyBtn',
}


</script>


<style>
.toggle-wrapper {
    width: 60px;
    height: 30px;
    /* border: 1px solid black; */
    box-shadow: inset 0 0 0 1px black;
    border-radius: 50px;
    display: inline-block;
    position: relative;
}

.toggle-indicator {
    display: inline-block;
    height: 30px;
    width: 30px;
    background-color: black;
    border-radius: 50px;
    position: absolute;
    /* left: 0; */
    /* top: 0; */
    left: 0;
    transition: transform 0.4s ease;
}

/* .toggle-indicator:hover {
    transform: translateX(100%);
} */
</style>