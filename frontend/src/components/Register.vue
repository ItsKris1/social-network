<template>
    <!-- <form @submit.prevent="submitRegData">
        <div id="frame1">
            <div id="frame3">
                <div id="form-input">
                    <label for="email">Email</label>
                    <input v-model="form.email" id="email" type="email">
                </div>
                <div id="form-input">
                    <label for="">Password</label>
                    <input v-model="form.password" type="password">
                </div>
                <div id="form-input">
                    <label for="">First Name</label>
                    <input v-model="form.firstname" type="text">
                </div>
                <div id="lastname-input">
                    <label for="">Last Name</label>
                    <input v-model="form.lastname" type="text">
                </div>
                <div id="birtday-input">
                    <label for="">Date Of Birth</label>
                    <input v-model="form.dateofbirth" type="date">
                </div>
                <div id="nickname-input">
                    <label for="nickname">Nickname</label>
                    <input v-model="form.nickname" id="nickname" type="text">
                </div>
                <div id="avatar-input">
                    <label placeholder="avatar">Avatar</label>
                    <input @change="checkPicture" type="file" accept="image/png, image/gif, image/jpeg">
                </div>
                <div id="about-me-input">
                    <label for="">About me</label>
                    <input v-model="form.aboutme" type="text">
                </div>
                <button id="regBtn" type="submit">Register</button>
            </div>
        </div>
    </form> -->

    <div class="register__wrapper">



        <img src="../assets/pexels-cottonbro-5053739.jpg" alt="man holding phone">




        <div class="register">
            <h1>Create your account</h1>
            <form @submit.prevent="submitRegData" id="register__form">

                <div class="form-group">
                    <div class="form-input">
                        <label for="firstname">First name</label>
                        <input v-model="form.firstname" type="text" name="firstname" id="firstname">
                    </div>
                    <div class="form-input">
                        <label for="email">Email</label>
                        <input v-model="form.email" type="email" name="email" id="email">
                    </div>

                    <div class="form-input">
                        <label for="date">Date of Birth</label>
                        <input v-model="form.dateofbirth" type="date" name="date" id="date">
                    </div>

                    <div class="form-input">
                        <label for="aboutme">About me</label>
                        <textarea v-model="form.aboutme" id="aboutme" name="aboutme" rows="4" cols="50"></textarea>
                    </div>
                </div>

                <div class="form-group">
                    <div class="form-input">
                        <label for="lastname">Last name</label>
                        <input v-model="form.lastname" type="text" name="lastname" id="lastname">
                    </div>

                    <div class="form-input">
                        <label for="password">Password</label>
                        <input v-model="form.password" type="password" name="password" id="password">
                    </div>


                    <div class="form-input">
                        <label for="nickname">Nickname</label>
                        <input v-model="form.nickname" type="text" name="nickname" id="nickname">
                    </div>
                    <div class="form-input">
                        <label for="avatar">Avatar</label>
                        <input id="avatar" @change="checkPicture" type="file" accept="image/png, image/gif, image/jpeg">
                    </div>
                </div>
            </form>

            <button class="btn" form="register__form" type="submit">Create account</button>

        </div>
    </div>

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
.register__wrapper {
    display: flex;
    align-items: center;
    height: 80vh;
    max-height: 700px;
    width: max-content;
    margin: auto;
    background-color: var(--color-white);
    box-shadow: var(--container-shadow);
    border-radius: 20px;
    overflow: hidden;

}

.register {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 45px;
    padding: 0 50px;

}

.register form {
    display: flex;
    gap: 40px;
    /* width: 100%; */
}

.register form>* {
    flex: 1;
}

.register .form-group {
    gap: 25px;
}
</style>