<template>
    <div class="chatbox-wrapper">
        <div class="header">
            <p>{{ this.name }}</p>
            <i class="uil uil-times close" @click.stop="$emit('closeChat', this.name)"></i>
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

<script>
export default {
    props: ["name"],
    emits: ['closeChat']

}
</script>
<style scoped>
.chatbox-wrapper {
    height: 400px;
    width: 300px;
    display: flex;
    flex-direction: column;
    box-shadow: var(--container-shadow);
    border-radius: 5px 5px 0 0;
    overflow: hidden;
    --padding: 15px;
    --msg-border-rad: 13px;
    --msg-padding: 7.5px 10px;
    --msg-margin-b: 5px;
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
    color: var(--color-lg-black);
    font-size: 14px;
    overflow-y: auto;
}

.content>*:not(:last-child) {
    margin-bottom: calc(var(--msg-margin-b) * 4);
}

.recieved-messages,
.sent-messages {
    display: flex;
    flex-direction: column;
}

.sent-messages>*:not(:last-child),
.recieved-messages>*:not(:last-child) {
    margin-bottom: var(--msg-margin-b);
}

.sent-message,
.recieved-message {
    padding: var(--msg-padding);
    border-radius: var(--msg-border-rad);
}

.recieved-messages {
    align-items: flex-start;
}


.recieved-message {
    background-color: rgb(185, 185, 185);
}


.recieved-message .message-author {
    padding-bottom: 3.5px;
}

.sent-messages {
    align-items: flex-end;
}

.sent-message {
    background-color: rgb(212, 212, 212);

}


/* SEND MESSAGE FORM */
.send-message {
    background-color: var(--color-white);
    padding: 10px;
}

.send-message form {
    display: flex;
    border-radius: 30px;
    padding: 10px 20px;
    overflow: hidden;
    box-shadow: var(--container-shadow);
}

.send-message form input {
    padding: 0;
    box-shadow: none;
    background-color: inherit;
}

.send-message form button {
    border: none;
    background-color: inherit;
    font-size: 1.25em;
}
</style>