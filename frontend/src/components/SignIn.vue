<template>
  <div id="login">
    <!-- <div class="signIn"> -->
    <span id="wback">Welcome back!</span>
    <div style="input">
      <label id="labelId">E-mail</label>
      <input id="rectangle1" type="email">
    </div>

    <label>Password</label>
    <input type="password">
    <button type="submit">Login</button>
  <div id="reglink">
    <span>Need an account?</span>
    <router-link to="/reg">SIGN UP HERE</router-link>
  </div>
  </div>
  <!-- </div> -->
</template>


<script>
export default {
  name: 'SignIn'
}
</script>

<style>
#login {
  position: relative;
  width: 1440px;
  height: 1488px;

  background: #FFFFFF;
}

#wback {
  /* Welcome back! */
  position: absolute;
  width: 287px;
  height: 48px;
  left: 539px;
  top: 276px;

  font-family: 'Inter';
  font-style: normal;
  font-weight: 400;
  font-size: 40px;
  line-height: 48px;

  color: #000000;
}
.input{
  /* Input / text */


/* Auto layout */

display: flex;
flex-direction: column;
align-items: flex-start;
padding: 0px;
gap: 9px;

position: absolute;
width: 212px;
height: 57px;
left: 592px;
top: 396px;
}

#labelId {
  /* Label */
  width: 36px;
  height: 17px;

  font-family: 'Inter';
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  line-height: 17px;
  /* identical to box height */

  text-align: center;

  color: #000000;


  /* Inside auto layout */
  flex: none;
  order: 0;
  flex-grow: 0;
}

#rectangle1 {
  /* Rectangle 1 */
  width: 212px;
  height: 31px;

  background: #EDEDED;
  border: 1px solid #918F8F;
  border-radius: 4px;

  /* Inside auto layout */
  flex: none;
  order: 1;
  align-self: stretch;
  flex-grow: 0;
}
</style>