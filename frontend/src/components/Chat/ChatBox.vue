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
                <p class="sent-message">This is a longer text</p>
                <p class="sent-message">This is even a lot more longer text helwhwehlwehwhlw</p>
            </div>

            <div class="recieved-messages">
                <p class="message-author">Roy Jones</p>
                <p class="recieved-message">Hello there!</p>
                <p class="recieved-message">How are you?</p>
            </div>

            <div class="sent-messages">
                <p class="sent-message">I am good mate!</p>
            </div>

            <div class="recieved-messages">
                <p class="message-author">Roy Jones</p>
                <p class="recieved-message">Good to know!</p>
            </div>

            <div class="sent-messages">
                <p class="sent-message">Indeed!</p>
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
.wrapper {
    height: 400px;
    width: 350px;
    display: flex;
    flex-direction: column;
    box-shadow: var(--container-shadow);
    --padding: 10px 12px;
    --msg-border-rad: 13px;
    --msg-padding: 7.5px 10px;
    --msg-margin-b: 7.5px;
}

.header {
    display: flex;
    justify-content: space-between;
    background-color: var(--color-blue);
    color: var(--button-content);
    padding: 12px 20px;
}

.content {
    flex: 1;
    background-color: var(--color-white);
    padding: var(--padding);
    font-size: 14px;
    overflow-y: auto;
}

.content>*:not(:last-child) {
    margin-bottom: calc(var(--msg-margin-b) * 2);
}


.recieved-message {
    background-color: gray;
    padding: var(--msg-padding);
    border-radius: var(--msg-border-rad);
}

.recieved-messages {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
}

.recieved-messages>*:not(:last-child) {
    margin-bottom: var(--msg-margin-b);
}

.recieved-messages :first-child {
    /* padding-left: 5px; */
    margin-bottom: 3.5px;
}

.sent-message {
    background-color: rgb(168, 168, 168);
    padding: var(--msg-padding);
    border-radius: var(--msg-border-rad);
}

.sent-messages {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
}

.sent-messages>*:not(:last-child) {
    margin-bottom: var(--msg-margin-b);
}

.send-message {
    background-color: orange;
    padding: var(--padding);
}

.send-message form {
    display: flex;
    /* padding: var(--padding); */

    background-color: var(--color-white);

    background-color: blue;
}

.send-message input {}

.send-message button {
    border: none;
    background-color: inherit;
}
</style>