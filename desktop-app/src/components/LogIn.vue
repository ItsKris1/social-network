<template>
    <div class="sign-in">
      <h1>Sign in</h1>
      <form class="form-group" @submit.prevent="loginSubmit" id="sign-in__form">
        <div class="form-input">
          <label for="username">Email</label>
          <input type="email" id="email" v-model="signInForm.login" required>
        </div>
        <div class="form-input">
          <label for="password">Password</label>
          <input type="password" id="password" v-model="signInForm.password" required>
        </div>
      </form>
      <div>
        <button class="btn" form="sign-in__form" type="submit">Sign in</button>
        <p>Need an account?
          <a href="http://localhost:8080/reg" id="sign-up" target="_blank">Register here</a>
        </p>
      </div>
    </div>
</template>

<script>
export default{
    name:"LogIn",
    data() {
        return {
            signInForm: {
                login: "",
                password: "",
            },
        };
    },
    methods:{
        async loginSubmit(){
            await fetch("http://localhost:8081/signin", {
                credentials: "include",
                method: "POST",
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(this.signInForm)
            })
            .then(response => response.json())
            .then(json =>{
                if (json.message === "Login successful"){
                    // create web socket
                     this.$store.dispatch("createWebSocketConn")
                     .then(()=> this.$router.push('/chat'))
                }
                // else notify about wrong credentials
            })
        }
    }
}
</script>

<style >
  .sign-in{
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 40px;
  margin: 0 auto;
  padding: 0 70px;
  }

  .form-group {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.sign-in button {
  margin-bottom: 10px;
}

#sign-up {
  font-weight: 500;
}
</style>