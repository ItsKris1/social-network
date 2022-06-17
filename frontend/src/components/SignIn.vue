<template>
  <form @submit.prevent="signSubmit">
    <span>Welcome back!</span>
    <div style="input">
      <label>E-mail</label>
      <input v-model="signInForm.login" id="rectangle1" type="email" required>
    </div>

    <label>Password</label>
    <input v-model="signInForm.password" type="password" required>
    <button>Login</button>
  </form>
  <div id="reglink">
    <span>Need an account?</span>
    <router-link to="/reg">SIGN UP HERE</router-link>
    <b-icon-door-closed> </b-icon-door-closed>
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
    toast() {
      /*---------------           Here is toast example             --------------------*/
      // 
      this.$toast.open({
        message: 'Data sent!',
        type: 'default', //One of success, info, warning, error, default

        //optional options
        position: 'bottom-right', //One of top, bottom, top-right, bottom-right,top-left, bottom-left
        duration: 3000, //Visibility duration in milliseconds, set to 0 to keep toast visible
        dismissible: true, //Allow user dismiss by clicking
        onClick: null, //Do something when user clicks(function)
        onDismiss: null, //Do something after toast gets dismissed(function)
        queue: false, // Wait for existing to dismiss before showing new
        pauseOnHover: true, //Pause the timer when mouse on over a toast
      });
    },
    async signSubmit() {
      try {
        // await fetch('https://bfdf8b79-b1e1-40ce-8d02-896de58da3ca.mock.pstmn.io/signin', {
        await fetch('http://localhost:8081/signin', {
          credentials: 'include',  // Uncomment when testing on real-server
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
            if (json.message === "Login successful") {
              this.$toast.open({
                            message: 'Login success!',
                            type: 'success', //One of success, info, warning, error, default
                        })
              this.$router.push("/main");
            } else {
              this.$router.push("/");
              this.$toast.open({
                            message: json.message,
                            type: 'error', //One of success, info, warning, error, default
                        })
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