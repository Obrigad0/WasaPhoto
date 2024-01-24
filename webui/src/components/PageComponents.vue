<script>
  export default {
    data: function() {
		  return {
        token: localStorage.getItem('token')
      }
    },

    methods:{
      toHome(){
        this.$router.replace("/home")
      },
      toSearch(){
        this.$router.push({ path: '/search', query: { searchValue: this.username } });
      },

      logout(){
          localStorage.removeItem('token')
          this.$router.replace("/login")

      },

      toMyProfile(){
        this.$router.replace("/profile/"+localStorage.getItem('token'))

      },
    },

    mounted(){
      if(localStorage.getItem("token")<=0){
        console.log("Non sei loggato!!!")
        this.$router.replace("/login")
      }
    }
    
  };
</script>
  
<template>
  <div class="id">
      <div class="top">
        <div>
          <h1 @click="toHome()" class="titolo">WASAPHOTO</h1>
        </div>
        
        <div style="display: flex; align-items: center;">
            <button @click="logout()">logout</button>
            <button @click="toMyProfile()">il mio profilo</button>
            <input class="input" type="text" placeholder="cerca utenti" v-model="username">
            <button class="cerca" @click="toSearch()" style="padding: 2px 8px; background-color: #734f59; color: #fff; border: none; cursor: pointer;">find</button>
        </div>
      </div>
      <div class="box" style="height: 84vh; background-color: white;">
        <slot></slot>
      </div>
      <div class="box" style="height: 8vh; padding-top: 10px;background-color: #f1f1f1;" >
      <footer class="footer">
        <p style="color: #540e23">2024 WasaPhoto Flavio Corsetti 1997818</p>
        </footer>
    </div>
  </div>

</template>


<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

  body, html {
    margin: 0;
    padding: 0;
    height: 100%;
  }
  .app {
    height: 100%;
    display: flex;
    flex-direction: column;
    font-family: 'Rubik', sans-serif;

  }
  .top{
    display: flex; 
    justify-content: space-between; 
    align-items: center; 
    padding: 0 20px;
    height: 5vh;
    background-color: #830c2c;
    border-radius: 25px;
  }
  .box {
    justify-content: center;
    align-items: center;
    text-align: center;
    font-size: 20px;
    box-sizing: border-box;
  }
  .footer {
    text-align: center;
    
  }
  .titolo { 
    padding-left: 2vw;
    text-align: center;
    font-size: 3vw;
    color: white; 
    text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000; /* Ombra per simulare il bordo nero con dimensione 2px */
  }

  .input{
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
    height: 20px;
  }

  .cerca{
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
    height: 25px;

  }

</style>