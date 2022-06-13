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

            // for multipart form to work body should be of FormData type
            let formData = new FormData();
            formData.set('avatar', this.form.avatar);
            formData.set('email', this.form.email);
            formData.set('password', this.form.password);
            formData.set('firstname', this.form.firstname);
            formData.set('lastname', this.form.lastname);
            formData.set('dateofbirth', this.form.dateofbirth);
            formData.set('nickname', this.form.nickname);
            formData.set('aboutme', this.form.aboutme);

            // await fetch('https://d0eb791a-34e8-410a-afb4-595a0ed0c134.mock.pstmn.io/reg', {
            await fetch('http://localhost:8081/register', {
                credentials: 'include',
                method: 'POST',
                body: formData
            })
                .then((res) => {
                    // console.log(res)
                    if (res.status === 409) {
                        this.$toast.open({
                            message: 'Email already taken',
                            type: 'error', //One of success, info, warning, error, default
                        })
                    } else if (res.status === 400) {
                        this.$toast.open({
                            message: 'Bad request',
                            type: 'error', //One of success, info, warning, error, default
                        })
                    } else {
                        this.$toast.open({
                            message: 'Sucessfully registered!',
                            type: 'success', //One of success, info, warning, error, default
                        })
                        this.$router.push("/");
                    }
                });
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
                this.$toast.open({
                            message: 'File is not an image.',
                            type: 'error', //One of success, info, warning, error, default
                        })
                return
            }
            if (file.size > 2048000) {
                console.log('File size is more than 2 MB.');
                this.$toast.open({
                            message: 'File size is more than 2 MB.',
                            type: 'error', //One of success, info, warning, error, default
                        })
                return
            }
            this.form.avatar = file;
        }
    }
}
</script>

<style>
</style>