<template>
  <form @submit.prevent="signSubmit">
    <span id="wback">Welcome back!</span>
    <div style="input">
      <label id="labelId">E-mail</label>
      <input v-model="signInForm.login" id="rectangle1" type="email">
    </div>

    <label>Password</label>
    <input v-model="signInForm.password" type="password" required>
    <button type="submit">Login</button>
  </form>
  <div id="reglink">
    <span>Need an account?</span>
    <router-link to="/reg">SIGN UP HERE</router-link>
  </div>
  <!-- <p>{{signInForm.login}}+{{signInForm.password}}</p> -->
</template>


<script>
export default {
  name: 'SignIn',
  data() {
    return {
      signInForm: {
        login: "",
        password: "",
      },
    }
  },
  methods: {
    async signSubmit() {
      try {
        await fetch('https://93e46479-d19c-41a8-83b3-9c33e3dbaeea.mock.pstmn.io/login', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            "Content-Type": "application/json"
          },
          body: JSON.stringify(this.signInForm)
        })
          .then((response => response.json()))
          .then((json => {
            console.log(json)
            if (json.response === "OK") {
              this.$router.push("/main");
            } else {
              this.$router.push("/");
            }
          }
          ))
      }
      catch { }

    },
  },
}
</script>

<style>
</style>