<template>
    <div class="wrapper">
        <p class="label">Avatar</p>

        <label for="fileUpload">
            <div class="btn">Browse</div>

            <p :class="textClass">{{ fileName }}</p>
        </label>

        <input type="file" name="fileUpload" id="fileUpload" @change="printFile">
    </div>

</template>

<script>
export default {
    data() {
        return {
            fileName: ""
        }
    },

    methods: {
        printFile(e) {
            const file = e.target.files[0];
            this.fileName = file.name;
        }
    },


    computed: {
        fileName() {
            if (this.fileName === "") {
                return "Choose a file...gweeeeeeeeeeeeeeeeeeggggggggggggggggggggggggggggggggggggggggggggggggggasfffffffffffffffffffffffffff"
            } else {
                return this.fileName
            }
        },

        textClass() {
            if (this.fileName === "") {
                return { placeholder: true }
            } else {
                return { selected: true }
            }
        }
    }
}
</script>

<style scoped>
.wrapper {
    flex-grow: 0;
    max-width: 100%;
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
    font-size: 14px;
}

input {
    position: absolute;
    opacity: 0;
    width: 0.1px;
    height: 0.1px;
    z-index: -1;


}

.placeholder {
    padding: 7.5px;
    color: var(--color-placeholder);
    position: absolute;
    white-space: nowrap;

}

.selected {
    padding: 7.5px;
    color: var(--color-lg-black);
    width: 100%;
}

.btn {
    border-radius: 5px;
    background-color: rgb(190, 190, 190);
    padding: 7.5px 10px;
    color: var(--color-lg-black);
}
</style>