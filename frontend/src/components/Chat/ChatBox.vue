<template>
    <div class="wrapper">
        <div class="header">
            <p>Roy Jones</p>
            <i class="uil uil-times close"></i>
        </div>
        <div class="content">
            <div class="recieved-messages">
                <p class="message-author">Roy Jones</p>
                <p class="recieved-message">Hello there!</p>
                <p class="recieved-message">How are you?</p>
            </div>
            <div class="sent-messages">
                <p class="sent-message">I am good mate!</p>
            </div>

        </div>
        <div class="send-message">
            <form>
                <input type="text" name="sent-message" id="sent-message__input" placeholder="Send a message">
                <button><i class="uil uil-message"></i></button>
            </form>
        </div>
    </div>

</template>

<style scoped>
.box {
    height: 400px;
    width: 350px;
    border: 2px solid blue;
}
</style>