<template>
    <div id="searchDiv">
        <input id="inputBox" type="text" placeholder="Search for user or group">
        <img id="glass" src="../assets/glass.png" alt="glass.png">
    </div>
</template>


<script>
export default {
    name: 'Search',
    data() {
        return {

        }
    },
    methods: {},
}
</script>


<style>
#searchDiv {
    display: flex;
    flex-direction: row;
    /* justify-content: space-between; */
    align-items: center;
    padding: 10px;
    /* gap: 53px; */

    width: 257px;
    height: 15px;

    background: #FFFFFF;
    border: 1px solid #706A6A;
    border-radius: 3px;
    border: 1px solid #706A6A;
}

#inputBox {
    border: none;
    width: 230px;
}

*:focus {
    outline: none;
}

#glass {
    margin-left: auto;
}
</style>