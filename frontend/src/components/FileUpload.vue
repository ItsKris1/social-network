<template>
    <div class="form-input">
        <p class="label">Avatar</p>

        <label for="fileUpload">
            <div class="btn">Browse</div>
            <p class="text" :class="textClass">{{ getFileName }}</p>
        </label>

        <div v-if="imageFile" @click=removeImage class="btn" style="width: max-content">Remove image</div>

        <input type="file" name="fileUpload" id="fileUpload" @change="checkPicture" ref="fileUpload">
    </div>

</template>

<script>

export default {
    emits: ['fileUploaded'],
    data() {
        return {
            imageFile: null,
        }
    },

    methods: {
        removeImage() {
            this.imageFile = null;
            this.$refs.fileUpload.value = "";
        },

        checkPicture(e) {
            let files = e.target.files;
            if (!files.length) {
                return;
            }
            const file = files[0];
            const [extension] = file.type.split("/");
            if ((!(extension == "image"))) {
                console.log("File is not an image.");
                this.$toast.open({
                    message: "File is not an image.",
                    type: "error", //One of success, info, warning, error, default
                });
                return;
            }
            if (file.size > 2048000) {
                console.log("File size is more than 2 MB.");
                this.$toast.open({
                    message: "File size is more than 2 MB.",
                    type: "error", //One of success, info, warning, error, default
                });
                return;
            }
            this.imageFile = file;
            this.$emit("fileUploaded", this.imageFile)

        }
    },

    computed: {
        textClass() {
            if (this.imageFile === null) {
                return { placeholder: true }
            } else {
                return { selected: true }
            }
        },

        getFileName() {
            if (this.imageFile === null) {
                return "Choose a file..."
            } else {
                return this.imageFile.name
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