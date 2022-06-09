<template>
    <form @submit.prevent="submitRegData">
        <div id="frame1">
            <div id="frame3">
                <div id="email-input">
                    <label class="label">Email</label>
                    <input v-model="form.email" id="" class="rectangle1" type="email">
                </div>
                <div id="password-input">
                    <label class="label">Password</label>
                    <input v-model="form.password" class="rectangle1" type="password">
                </div>
                <div id="firstname-input">
                    <label class="label">First Name</label>
                    <input v-model="form.firstname" class="rectangle1" type="text">
                </div>
                <div id="lastname-input">
                    <label class="label">Last Name</label>
                    <input v-model="form.lastname" class="rectangle1" type="text">
                </div>
                <div id="birtday-input">
                    <label class="label">Date Of Birth</label>
                    <input v-model="form.dateofbirth" class="rectangle1" type="date">
                </div>
                <div id="nickname-input">
                    <label for="nickname" class="label">Nickname</label>
                    <input v-model="form.nickname" id="nickname" class="rectangle1" type="text">
                </div>
                <div id="avatar-input">
                    <label class="label" placeholder="avatar">Avatar</label>
                    <input @change="checkPicture" class="rectangle1" type="file"
                        accept="image/png, image/gif, image/jpeg">
                </div>
                <div id="about-me-input">
                    <label class="label">About me</label>
                    <input v-model="form.aboutme" class="rectangle1" type="text">
                </div>
                <button id="regBtn" type="submit">Register</button>
            </div>
        </div>
    </form>
</template>

<script>

export default {
    name: 'Register',
    data() {
        return {
            form: {
                email: "",
                password: "",
                firstname: "",
                lastname: "",
                dateofbirth: null,
                nickname: "",
                avatar: null,
                aboutme: "",
            },
        }
    },
    methods: {
        async submitRegData() {            
            try {
                await fetch('https://d0eb791a-34e8-410a-afb4-595a0ed0c134.mock.pstmn.io/reg', {
                    method: 'POST',
                    headers: {
                        // 'Accept': 'application/json',
                        "Content-Type": "multipart/form-data"
                    },
                    body: this.form
                })
                    // .then((response => response.json()))
                    // .then((json => console.log(json)))
            }
            catch { }
            this.$router.push("/");
        },
        checkPicture(e) {
            let files = e.target.files
            if (!files.length) {
                return;
            }
            const file = files[0]

            const [extension] = file.type.split("/")
            if ((!(extension == "image"))) {
                console.log('File is not an image.');
                return
            }
            if (file.size > 2048000) {
                console.log('File size is more than 2 MB.');
                return
            }
            this.form.avatar = file;
        }
    }
}
</script>

<style>
</style>