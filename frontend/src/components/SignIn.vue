<template>
  <form @submit.prevent="signSubmit">
    <span>Welcome back!</span>
    <div style="input">
      <label>E-mail</label>
      <input v-model="signInForm.login" id="rectangle1" type="email" >
    </div>

    <label>Password</label>
    <input v-model="signInForm.password" type="password" >
    <button @click="toast">Login</button>
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
    toast(){
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