<script>
export default {
  data: function(){ //variabili usati nei metodi
    return {
      userName: "", //username inserito dall'utente
      errore: null, //errore ritornato dal server
      //token: null,  //token ritornato dal server
    }
  },
  methods: { //metodi
    async login(){ //metodo di login asincrono

      this.errore = null;
      console.log("metodo login avviato");

      try {

        let response = await this.$axios.post("/session",{
					name: this.userName.trim()
				});

        console.log("id inviato dal db:"+response.data);

        //memorizziamo il token 
        localStorage.setItem("token",response.data);
        //ci spostiamo nella pagina principale di wasaphoto, ovvero la pagina dello stream dell'utente
        this.$router.replace("/home")

        //this.$emit('updatedLoggedChild',true)
        console.log("Tutto ok, loggato");
      } catch(e){
            //cattura l'errore tornato dal server
            //this.errore = "Errore, login non effettuato!";
            this.errore = e.toString();
            console.log(this.errore);

      }
    }
  },
  
  mounted(){ 
    //verifica se l'utente e' gia loggato
		if (localStorage.getItem('token') && localStorage.getItem('token')!= 1){
      //se gia loggato non puo rimanere nella pagina di login e viene spostato alla sua home
      console.log("sei gia loggato!!");
      this.$router.replace("/home")
		}
	},
};
</script>


<template>
<div class="container">
  <div class="principale">
    <div>
      <h1 class="titolo">WASAPHOTO</h1>
      <div class="form">
        <input type="text" placeholder="Inserisci username" class="input" v-model="userName" maxlength="16" minlength="3"/>
        <button class="login" @click="login()" :disabled="userName == null || userName.length > 16 || userName.length < 3 || userName.trim().length < 3">Login</button>
        <p class="messaggioErrore" v-if="errore">{{ errore }}</p>
      </div>
    </div>
  </div>
  <footer class="footer">
      <p> 2024 WasaPhoto Flavio Corsetti 1997818</p>
  </footer>
</div>
</template>


<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

body,html {
    margin: 0;
    padding: 0;
}


.container{
    height: 90vh;
    margin: 0;
    border-radius: 6px;

}
.principale {
  font-family: 'Rubik', sans-serif;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 95%;
  border: 1px solid aquablue;
  box-sizing: border-box;
  background-image: url(../assets/images/loginBackground.jpg);
  background-position: center;
  background-size: cover;

}

.titolo {
  text-align: center;
  font-size: 8vw;
  margin-bottom: 20px;
  color: white; /* Colore del testo bianco */
  text-shadow: -2px -2px 0 #000, 2px -2px 0 #000, -2px 2px 0 #000, 2px 2px 0 #000; /* Ombra per simulare il bordo nero con dimensione 2px */
}

.form {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.input {
  width: 24vw;
  height: 21px;
  margin-bottom: 10px;
  padding: 8px;
  text-align: center; 
  font-size: 13px;
  border-radius: 6px;
}

.login {
  width: 150px;
  height: 30px;
  background-color: #540e23;
  color: white;
  border: 2px solid black;
  font-size: 18px;
  font-weight: bold;
  border-radius: 6px;
  cursor: pointer;
}

.messaggioErrore{
  color: white;
  font-size: 23px;
}

.footer {
  text-align: center;
  padding: 10px;
  background-color: #f1f1f1;
  height: 7%;
}

</style>