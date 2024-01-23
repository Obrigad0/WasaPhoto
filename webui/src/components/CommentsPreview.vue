<script>
export default {
  data() {
    return{
        nomeAutoreCommento: "",
        tokenn: null,
        isMouseOver: false

    };
  },

  props: ["uId","cId","testo","iId","autore"],  //Autore e' l'autore della foto NON DEL COMMENTO!!!! uId e' l'autore del commento

  methods: {

    goToProfile(){
      this.$router.replace("/profile/"+this.uId)
    },

    checkSeMio(){
        //controlla se il commento e' stato creato dall'utewnte attuale
        if(this.tokenn == this.cId){
            return true
        }
        return false
    },

    cambiaColore(stato) {
        if(this.checkSeMio()){
            this.isMouseOver = stato;
        }
    },
    
    async removeComment() {
        if(this.tokenn == this.cId){
            try {
                await this.$axios.delete("/user/"+ this.autore +"/images/"+this.iId+"/comments/"+ this.cId) 
                //this.comments.pop(this.idCommento); //rimettere
            }catch (e){
                    this.errore = "{"+ e +"}"
            }
                this.comments.pop(this.idCommento); // levare!!
        }
    },

    async idToName(){
        try {
          let response = await this.$axios.get("/user/"+ this.uId) 
          this.nomeAutoreCommento = response.data.name
        }catch (e){
            this.errore = "{"+ e +"}"
      }
    }
  },

  mounted(){
    this.tokenn = localStorage.getItem("token")
    console.log(this.testo)
  }
}
</script>

<template>
    <div class="commentoP">        
      <div><p style="font-size: 14px;" class="usnme" @click="goToProfile()">{{ this.uId }} x</p> <p style="font-size: 14px;" >ha scritto:</p></div>
      <div><p @click="removeComment()" @mouseover="cambiaColore(true)" @mouseout="cambiaColore(false)" :class="{ rosso: isMouseOver }" title="Elimina">  {{ testo }} </p></div>
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