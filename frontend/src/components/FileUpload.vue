<template>
    <p class="label">Avatar</p>

    <label for="fileUpload">
        <p class="placeholder">Choose a file...</p>
        <div class="btn">Browse</div>
    </label>

    <input type="file" name="fileUpload" id="fileUpload">

</template>


<style scoped>
.label {
    margin-bottom: 5px;
}

label {
    display: flex;
    justify-content: space-between;
    align-items: center;

    position: relative;

    background-color: var(--input-bg);
    box-shadow: var(--container-shadow);
    border-radius: 5px;

    cursor: pointer;
}

input {
    position: absolute;
    opacity: 0;
    width: 0.1px;
    height: 0.1px;
    z-index: -1;


}



.placeholder {
    font-size: 14px;
    padding: 7.5px;
    color: var(--color-placeholder);

}



.btn {
    border-radius: 5px;
    background-color: rgb(190, 190, 190);
    padding: 7.5px 10px;
    color: var(--color-lg-black);
}
</style>