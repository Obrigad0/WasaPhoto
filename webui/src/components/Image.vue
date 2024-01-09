<script>
export default {
  data() {
    return {
        imageUrl:  'https://www.laziostylestore.com/images/lazio/products/large/LZ22A03_16.webp',
        descrizione: 'Forza Forza Lazio!',
        isLiked: false,
        comments: [],
        usersLike: [],
        comment: '',
        post: true,
        likeP: false,
        errore: "",
        token: null,
        coloreCommento: "black",
        /*
        idCommento: null,
        autore: null,
        like: null,
        commenti: null,
        data: null,
        iId: null,
        desc: null,
        */

    };
    
  },

  //props: ['autore','like','commenti',"data","iId","desc"], //,"token" ??

  methods: {
    async toggleLike() {  //funzione
      this.isLiked = !this.isLiked;//levare

      if(this.autore == localStorage.getItem('token')){
        //niente autolike sorry user
        return
      }else{
        if(this.isLiked){
          //rimuovo il like
          try {
            await this.$axios.delete("/user/"+ this.autore  +"/images/"+this.iId+"/like/"+ localStorage.getItem('token')) 
            this.like.pop(localStorage.getItem('token'))
          }catch (e){
            this.errore = "{"+ e +"}"
          }
        }else{
          //aggiungo il like
          try {
            await this.$axios.put("/user/"+ this.autore +"/images/"+this.iId+"/like/"+ localStorage.getItem('token'))
            this.like.push(localStorage.getItem('token'))
          }catch (e){
            this.errore = "{"+ e +"}"
          }
        }
      }
      console.log("qui")
      this.isLiked = !this.isLiked;

    },

    async addComment() {  //funzione
      
      try {
          await this.$axios.post("/user/"+ this.autore +"/images/"+this.iId+"/comments/", {text: comment}) 
          if (this.comment.trim() !== '') {
            this.comments.push(this.comment);
            this.comment = '';
          }
      }catch (e){
            this.errore = "{"+ e +"}"
      }
      
      if (this.comment.trim() !== '') { //LEVARE!!!
            this.comments.push(this.comment);
            this.comment = '';
      }
    },
    
    async removeComment() {
      try {
          await this.$axios.delete("/user/"+ this.autore +"/images/"+this.iId+"/comments/"+ this.idCommento) 
      }catch (e){
            this.errore = "{"+ e +"}"
      }
        this.comments.pop(this.idCommento);
    },
    
    toComment(){  //stile
        this.post = !this.post;
    },

    toLikes(){
        this.likeP = !this.likeP;
    },

    getPost(){
        //funzione
      try{
        //poi dovrei cambiare il js nel html
        this.likes = this.like
        this.comments = this.commenti
        //
        this.descrizione = this.desc
        this.imageUrl = __API_URL__+ "/user/"+this.autore+"/images/"+this.iId
      }catch(e){
        this.errore = "{"+ e +"}"
      }
    },

    changeColor(stato){
      if(this.autore == localStorage.getItem('token')){
        //se l'utente che ha postato il commento passa sul proprio commento
        //puo' elimnare il commento
      }else{

      }

    }
  },

  async mounted(){
    //await this.getPost();
  }

};
</script>

<template>
  <div class="principale">
    <div v-if="post && !likeP"  class="post">
      <img :src="imageUrl"  class="immagine">
      <div class="feedback">
        <button @click="toggleLike" class="like">
            <img v-if = "!isLiked" src="../assets/images/noLike.png">
            <img v-if = "isLiked"  src="../assets/images/like.png">
        </button>
        <input v-model="comment" @keyup.enter="addComment" placeholder="Aggiungi commento...">
        <button @click="addComment">
          <img  src="../assets/images/invia.png">
        </button>
      </div>
      <p v-if="errore === ''" class="descrizione">{{ descrizione }}</p>
      <p v-if="errore !== ''" class="descrizione">{{ errore }}</p>
      <div class="com">
        <button @click="toComment" >commenti</button>
        <button @click="toLikes" >like</button>
        <button  >elimina</button>
      </div>   
    </div>
    <div v-if="!post" class="commenti">
      <div  class="testo">
        <p  @mouseover="changeColor(true)" v-for="(comment, index) in comments" :key="index" @click="removeComment()">{{ comment }}</p>
      </div>
      <div class="com2">
        <button @click="toComment" >post</button>
      </div>   
    </div>
    <div v-if="likeP" class="commenti">
      <p>likes yeah</p>
      <div class="com2">
          <button @click="toLikes">post</button>
      </div>  
    </div>
  </div>
  </template>
  
  <style scoped>
      @import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

    .principale{
      width: 280px;
      border: 1px solid #ccc; 
      align-items: center;
      font-family: 'Rubik', sans-serif;
    }
    .post {
        width: 250px; /* Imposta la larghezza del singolo post */
        padding: 14px;
    }
    
    .immagine {
        width: 100%;
        height: auto;
    }
    .commenti{
      max-height: 240px; /* Imposta l'altezza massima del container dei commenti */
      width: 280px; /* Imposta la larghezza del singolo post */
      min-height: 40px;
    }
    .descrizione {
        margin-top: 10px;
    }
    
    .like{
        border: none;
        background-color: white;
        cursor: pointer;
    }
    .like img{
        height: 30px;
        width: 30px;
        padding-top: 3px;    
    }
    .feedback{
        width: 100%;
         display: flex;
         justify-content: center;
         align-items: center;
    }
    .feedback button {
        margin-right: 5px;
        border: none;
        background-color: white;
        cursor: pointer;
    }
    .com{
        display: flex;
        justify-content: center;

    }
    .testo{
      border: 1px solid #ccc; 
      max-height: 200px; /* Imposta l'altezza massima del container dei commenti */
      overflow-y: auto;
      width: 100%;
    }
    .com2 {
         display: flex;
        flex-wrap: nowrap;
        flex-direction: column;
      
    }
    .com button{
      border-bottom: 0;
        border-left: 0;
        border-right: 0;
        border-top: 0;
        background-color: white;
        cursor: pointer;

        }
    .com2 button{
        border-bottom: 0;
        border-left: 0;
        border-right: 0;
        border-top: 0;
        background-color: white;
        cursor: pointer;
        padding: 10px;
      }
    .feedback input{
        border-bottom: 1;
        border-left: 0;
        border-right: 0;
        border-top: 0;
        width: 60%;
        padding: 5px;
        margin-left: 10px;
        
    }

    .cambiacolore{
      color: black;
    }
    .cambiacolore:hover{
      color: red;
    }
  </style>