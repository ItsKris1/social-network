<template>
    <div class="form-input">
        <p class="label">Avatar</p>

        <label for="fileUpload">
            <div class="btn">Browse</div>
            <p class="text" :class="textClass">{{ getFileName }}</p>
        </label>

        <div v-if="fileName !== ''" @click=removeImage class="btn" style="width: max-content">Remove image</div>

        <input type="file" name="fileUpload" id="fileUpload" @change="getFileName = $event" ref="fileUpload">
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
        // getFileName(e) {
        //     const file = e.target.files[0];
        //     this.fileName = file.name;
        // },

        removeImage() {
            this.fileName = "";
            this.$refs.fileUpload.value = "";
        }
    },

    computed: {
        textClass() {
            if (this.fileName === "") {
                return { placeholder: true }
            } else {
                return { selected: true }
            }
        },

        getFileName: {
            get() {
                if (this.fileName === "") {
                    return "Choose a file..."

                } else {
                    return this.fileName
                }
            },

            set(e) {
                const file = e.target.files[0];
                this.fileName = file.name;
            }
        },
    }
}
</script>

<style scoped>
.wrapper {}

label {
    /* display: flex;
    justify-content: flex-end;
    align-items: center; */

    display: block;
    position: relative;

    background-color: var(--input-bg);
    box-shadow: var(--container-shadow);
    border-radius: 5px;

    cursor: pointer;
    font-size: 14px;
    overflow: hidden;
    /* width: 100%; */
}

input {
    position: absolute;
    opacity: 0;
    width: 0.1px;
    height: 0.1px;
    z-index: -1;


}

.text {
    position: absolute;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    top: 0;
    width: calc(100% - 70.25px);
    padding: 7.5px;

}

.placeholder {
    color: var(--color-placeholder);
}

.selected {
    color: var(--color-lg-black);
}

.btn {
    border-radius: 5px;
    background-color: rgb(190, 190, 190);
    padding: 7.5px 10px;
    color: var(--color-lg-black);
    overflow: hidden;
    width: 70.25px;

    margin-left: auto;
}
</style>