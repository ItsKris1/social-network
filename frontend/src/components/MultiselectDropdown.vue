<template>
    <div class="dropdown-wrapper">
        <p class="custom-label">{{ labelName }}</p>

        <ul class="checkedOptionsList" v-if="dropdownCheckedOptions.length !== 0">
            <li v-for="checkedOption in dropdownCheckedOptions"> {{ checkedOption }}</li>
        </ul>

        <div class="dropdown">
            <button type="button" @click="showDropdown = !showDropdown">
                {{ placeholder }}
                <img class="dropdown-arrow" src="../assets/icons/angle-down.svg" alt="" srcset="">
            </button>

            <ul class="item-list" v-show="showDropdown">
                <li v-for="option in content">
                    <input type="checkbox"
                           :id="option"
                           :value="option"
                           v-model="dropdownCheckedOptions" />
                    <label :for="option">{{ option }}</label>

                </li>
            </ul>
        </div>

    </div>

</template>


<script>
export default {
    props: ['labelName', 'placeholder', 'content', 'clearInput', 'checkedOptions'],
    emits: ['update:checkedOptions', 'inputCleared'],
    data() {

        return {
            dropdownCheckedOptions: [],
            showDropdown: false,
        }
    },

    watch: {
        clearInput() {
            if (this.clearInput) {
                this.dropdownCheckedOptions = [];
                this.showDropdown = false;
                console.log("Cleared input")

                this.$emit("inputCleared");
            }

        },

        dropdownCheckedOptions(checkedOption) {
            this.$emit('update:checkedOptions', checkedOption)
        }
    }
}

</script>


<style>
.dropdown-wrapper {
    display: flex;
    flex-direction: column;
    gap: 5px;

}

.dropdown {
    background-color: var(--input-bg);
    box-shadow: var(--container-shadow);
    border-radius: 5px;
}

.dropdown button {
    padding: 7.5px;

    border: none;
    font-family: 'Poppins', sans-serif;
    text-align: left;
    color: rgb(136, 136, 136);
    /* width: 250px; */
    position: relative;
    background-color: transparent;
    width: 100%;
    min-height: 35px;
}

.dropdown .item-list {
    padding: 7.5px;
    width: 100%;
}

.checkedOptionsList {
    display: flex;
    gap: 5px;
    padding: 5px 0;
}

.checkedOptionsList li {
    background-color: rgb(179, 179, 179);
    border-radius: 5px;
    padding: 5px;
    font-size: 14px;
}
</style>