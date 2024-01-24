<script>

import CommentsPreview from '@/components/CommentsPreview.vue';

export default {
  data() {
    return {
        imageUrl: "",
        descrizione: 'Forza Forza Lazio!',
        isLiked: false,
        comments: [],
        likes: [],
        comment: '',
        post: true,
        likeP: false,
        errore: "",
        token: null,
        coloreCommento: "black",
        isP: true,
        nomeAutore: "",
        nomiLike: [],
    };
    
  },
  components: {
    CommentsPreview,
  },

  props: ['autore','like','commenti',"data","iId","desc","isProfile","username"],

  methods: {
    async toggleLike() {  //funzione

      if(this.autore == localStorage.getItem('token')){
        //niente autolike sorry user
        return
      }else{
        if(this.isLiked){
          //rimuovo il like
          try {
            await this.$axios.delete("/user/"+ this.autore +"/images/"+this.iId+"/like/"+ localStorage.getItem('token')) 
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

    },

    async addComment() {  //funzione
      
      try {
          await this.$axios.post("/user/"+ this.autore +"/images/"+this.iId+"/comments/", {text: this.comment.trim()}) 
          if (this.comment.trim() !== '') {
            let com = { testo: this.comment, cId: 1, uId: localStorage.getItem("token") } //cambiare
            this.comments.push(com);
            this.comment = '';
          }
      }catch (e){
            this.errore = "{"+ e +"}"
      }

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
        this.token = localStorage.getItem('token')
        this.likes = this.like
        this.comments = this.commenti
        //console.log(this.comments)
        this.isP = this.isProfile
        this.nomeAutore = this.username
        this.descrizione = this.desc
        //get image, ritorna anche le informazioni dell'immagine, ma prendo solo il file mandato
        this.imageUrl = __API_URL__+ "/user/"+this.autore+"/imgUs/"+this.iId
      }catch(e){
        this.errore = "{"+ e +"}"
      }
    },

    changeColor(){
      console.log("sopra")
      if(this.autore == localStorage.getItem('token')){
        console.log("cambio il colore")
        //se l'utente che ha postato il commento passa sul proprio commento
        //puo' elimnare il commento
        var commento = document.getElementById("comP");
        commento.style.color = "red";
      }
        var commento = document.getElementById("comP");
        commento.style.color = "red";
    },

    goToProfile(){
      this.$router.replace("/profile/"+this.autore)
    },

    async idToName(id){
        try {
          let response = await this.$axios.get("/user/"+ id) 
          this.nomiLike.push(response.data.name) 
        }catch (e){
            this.errore = "{"+ e +"}"
      }
    }
  },

  async mounted(){
    await this.getPost();
    for (let i = 0; i < this.likes.length; i++) {
      await idToName(this.likes[i])
    }
  }

};
</script>

<template>
  <div class="principale">
    <div v-if="!isP" @click="">{{this.nomeAutore}}</div>
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
      <p class="descrizione">{{ descrizione }}</p>
      <!-- v-if="errore === ''"  <p v-if="errore !== ''" class="descrizione">{{ errore }}</p> -->
      <div class="com">
        <button @click="toComment" >commenti</button>
        <button @click="toLikes" >like</button>
        <button  >elimina</button>
      </div>   
    </div>
    <div v-if="!post" class="commenti">
      <div  class="testo">
        <CommentsPreview v-for="(commentoa, index) in comments"
          :key="index"
          :uId="commentoa.uId"
          :cId="commentoa.cId"
          :testo="commentoa.testo"
          :iId="this.iId"
          :autore="this.autore">
        </CommentsPreview>
        <!-- <p  id="comP" class="cambiacolore" v-for="(comments, index) in comments" :key="index" @click="removeComment()">{{ comments }}</p> -->
      </div>
      <div class="com2">
        <button @click="toComment" >post</button>
      </div>   
    </div>
    <div v-if="likeP" class="commenti">
    
    <div  class="testo">
        <p  id="comP" class="cambiacolore" v-for="(likes, index) in nomiLike" :key="index">{{ likes }}</p>
      </div>
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
        color: black;
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
      color: black;
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