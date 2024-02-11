<script>
export default {
    data() {
        return {
                descrizione: "",
        }
    },
    methods: {

      async creaPost(){

          let fileCaricato = document.getElementById('file')
          //let descrizione = document.getElementById('descrizione')

          if(fileCaricato !== undefined && this.$route.params.idUser == localStorage.getItem("token")){
              //prendo solo il primo file caricato
              fileCaricato = fileCaricato.files[0]
              //oggetto reader per leggere il contenuto del file
              //const reader = new FileReader();
              //Creo l'oggetto formdata che sara' inviato al server 
              let formData = new FormData();
              //memorizziamo l'oggetto come un file binario
              formData.append('file',fileCaricato);
              formData.append('descrizione', this.descrizione);
              
                  await this.$axios.post("/user/"+this.$route.params.idUser+"/images", formData, {
                      headers: {
                      'Content-Type': 'multipart/form-data',
                      },
                  })

          }else{
              //errore file non inserito
              console.log("errore file non inserito")
          }
          location.reload(true)

      }
  }
}   
</script>

<template>
  <div class="principale">
    <div class="ct"><p>Carica Foto</p></div>
    <div class="cf">
      <div class="im">
        <input type="file" id="file" accept=".jpg, .png">
      </div>
    </div>
    <div class="cd">
      <input id="descrizione" placeholder="Aggiungi descrizione..." v-model="descrizione">
    </div>
    <div class="go">
      <button @click="creaPost()">CREA</button>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

    .principale{
        width: 300px;
        height: 360px;
        border: 1px solid #ccc; 
        align-items: center;
        font-family: 'Rubik', sans-serif;
    }
    .ct{
        width: 100%;
        height: 13%;
        background-color: white;
        text-align: center;
    }
    .ct p{
        padding-top: 7px;
        margin-top: 0px;
        margin-bottom: 0px;
        font-size: 24px;
    }

    .cf{
        width: 100%;
        height: 55%;
        display: flex;
        justify-content: center;
        align-items: center;
        border: 1px solid #000; 

    }
    .im{
      background-color: white;
      width: 190px;
      height: 140px;
      border: 1px solid #000; 
      display: flex;
      justify-content: center;
      align-items: center;
    }
    .cd{
      width: 100%;
      height: 22%;
      display: flex;
      justify-content: center;
      align-items: center;
      border-top: 0px solid;
      border: 1px solid #000; 
    }

    .go{
      border: 1px solid #000; 
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 10%;
    }

    button{
    height: 30px;
    width: 100px;
    margin-right: 10px;
    border-radius: 10px;
  }

</style>
