<template>
    <i class="uil uil-smile" @click="toggleShowEmojis"></i>
    <div class="emojis" v-show="showEmojis">
        <p class="emoji" v-for="emoji in getEmojis"
           @click.stop="addEmoji(emoji)">{{ emoji }}</p>
    </div>

</template>

<script>
export default {
    props: ["input", "messagebox"],
    data() {
        return {
            emojiCodes: [128512, 128514, 128519, 128520],
            showEmojis: false
        }
    },

    computed: {
        getEmojis() {
            return this.emojiCodes.map((code) => String.fromCodePoint(code))
        }
    },

    methods: {
        addEmoji(emoji) {
            this.input.value += emoji;
        },

        toggleShowEmojis() {
            // console.log(this.messagebox.scrollHeight)
            // console.log(this.messagebox.scrollTop)
            // console.log(this.messagebox.offsetHeight)
            // if (this.showEmojis === false) {
            //     this.messagebox.scrollTop += 35;
            //     console.log("here")
            // }

            // if (this.showEmojis === false) {
            //     this.messagebox.scrollTop += 30;

            // } else {
            //     this.messagebox.scrollTop -= 30;
            // }
            // console.log(this.messagebox.scrollTop)

            // this.messagebox.scrollTop += 35;
            this.showEmojis = !this.showEmojis
        },



    }
}
</script>

<style scoped>
i {
    font-size: 1.35em;
    /* color: var(--color-lg-black); */
}

.emojis {
    display: flex;
    padding-top: 5px;
    width: 100%;
}
</style>