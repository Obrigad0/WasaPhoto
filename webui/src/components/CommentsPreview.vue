<script>
export default {
  data() {
    return{
        nomeAutoreCommento: "",
        tokenn: null,
        isMouseOver: false,
        eliminato: false,
    };
  },

  props: ["uId","cId","testo","iId","autore"],  //Autore e' l'autore della foto NON DEL COMMENTO!!!! uId e' l'autore del commento

  methods: {

    goToProfile(){
      this.$router.replace("/profile/"+this.uId) //uId e' ; autore del commento
    },

    checkSeMio(){
        //controlla se il commento e' stato creato dall'utewnte attuale
        if(this.tokenn == this.uId){
            return true
        }
        return false
    },

    cambiaColore(stato) {
        if(this.checkSeMio()){
            console.log("Il commento è mio")
            this.isMouseOver = stato;
        }
    },
    
    async removeComment() {
        if(this.checkSeMio()){
            try {
                await this.$axios.delete("/user/"+ this.autore +"/images/"+this.iId+"/comments/"+ this.cId) 
                console.log("funzionerà pop??")
                this.eliminato = true
            }catch (e){
                this.errore = "{"+ e +"}"
            }
        }else{
          console.log("il commento non è il mio")
        }
    },

    async idToName(){
        try {
          console.log("id di chi ha commentato 2 : "+this.uId)
          let response = await this.$axios.get("/user/"+ this.uId) 
          this.nomeAutoreCommento = response.data.name
          console.log("nome di chi ha commentato: "+this.nomeAutoreCommento)
        }catch (e){
            this.errore = "{"+ e +"}"
            console.log(errore)
      }
    }
  },

  mounted(){
    this.tokenn = localStorage.getItem("token")
    this.idToName()
    console.log("nuovo commento!, id :"+this.cId)
    console.log("id di chi ha commentato 1 : "+this.uId)
    console.log("il commento : "+this.testo)

    //console.log(this.testo)
  }
}
</script>

<template>
    <div class="commentoP">        
      <div><p  style="font-size: 14px;" class="usnme" @click="goToProfile()">{{ this.nomeAutoreCommento }}</p> <p style="font-size: 14px;" >ha scritto:</p></div>
      <div v-show="!eliminato" style=" word-wrap: break-word; "><p @click="removeComment()" :id="this.cId" @mouseover="this.cambiaColore(true)" @mouseout="this.cambiaColore(false)" :class="{ rosso: isMouseOver }" title="Elimina">  {{ this.testo }} </p> </div>
      <div v-show="eliminato" style=" word-wrap: break-word; "><p  style="color: red;">  commento eliminato </p> </div>
    </div>
</template>

<style scoped>
      @import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

    .commentoP{
        font-family: 'Rubik', sans-serif;
        display: flex;
        flex-direction: column;
        width: 99%;
        background-color: whitesmoke ; 
    }
    .commentoP div{
      display: flex;
      flex-direction: row;
      border: 1px solid lightgrey ; 
;
    }
    .commentoP p{
      margin: 5px;
      font-size: 20px;
    }

    .usnme:hover{
      color: blue;
    }

    .rosso {
        color: red;
    }

</style>