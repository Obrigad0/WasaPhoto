<script>

import CommentsPreview from '@/components/CommentsPreview.vue';

export default {
  data() {
    return {
        imageUrl: "",
        descrizione: 'descrizione',
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
        idAutore: null,
        nomiLike: [],
        eliminato: false,
    };
    
  },
  components: {
    CommentsPreview,
  },

  props: ['autore','like','commenti',"data","iId","desc","isProfile"],

  methods: {

    isLike(){
        return this.likes.includes(parseInt(localStorage.getItem('token'),10))
    },

    async toggleLike() {  //funzione

      if(this.autore == localStorage.getItem('token')){
        //niente autolike sorry user
        console.log("L'Utente non può mettere like ad un proprio post.")
        const contenitoreDescrizione = this.descrizione
        const elementoDesc = document.getElementById("stiDesc")
        this.descrizione = "AUTO-LIKE NON CONSENTITO!"
        elementoDesc.style.color = 'red';
        setTimeout(() => {
          this.descrizione = contenitoreDescrizione
          elementoDesc.style.color = 'black';
        }, 3000) 

        return
      }else{
        console.log("is liked era "+this.isLiked)
        if(this.isLiked){
          //rimuovo il like
          try {
            console.log("levo il like")
            await this.$axios.delete("/user/"+ this.autore +"/images/"+this.iId+"/like/"+ localStorage.getItem('token')) 
            this.likes.pop(parseInt(localStorage.getItem('token'),10))
            this.nomiLike = this.nomiLike.filter(elemento => elemento.id !== parseInt(localStorage.getItem('token'),10));
            this.isLiked = !this.isLiked
          }catch (e){
            this.errore = "{"+ e +"}"
            console.log(e)

          }
        }else{
          //aggiungo il like
          try {
            console.log("metto il like")
            await this.$axios.put("/user/"+ this.autore +"/images/"+this.iId+"/like/"+ localStorage.getItem('token'))
            this.likes.push(parseInt(localStorage.getItem('token'),10))
            await this.idToName(parseInt(localStorage.getItem('token'),10))
            this.isLiked = !this.isLiked

          }catch (e){
            this.errore = "{"+ e +"}"
            console.log(e)

          }
        }
        console.log("is liked ora è "+this.isLiked)

      }
      console.log("lista like alla foto:"+this.likes)

    },

    async addComment() {  
      
      try {
        if (this.comment.trim() !== '') {
            await this.$axios.post("/user/"+ this.autore +"/images/"+this.iId+"/comments/", {text: this.comment.trim()}) 
            let com = { text: this.comment, idComment: await this.getCommentId(), commenter: parseInt(localStorage.getItem('token'),10) } 
            if(this.comments == null){
              this.comments.push(com)
            }else{
              this.comments.unshift(com)
            } 
            this.comment = ''
            console.log("Array di commenti: "+this.comments)
            console.log("pushato, chi ha commentato "+com.commenter)
          }
      }catch (e){
            this.errore = "{ "+ e +" }"
            console.log("Errore nel caricamento del commento:"+this.errore)
      }

    },
    
    toAuthorProfile(){
      this.$router.replace("/profile/"+this.idAutore)
    },

    toComment(){  //stile
        this.post = !this.post;
    },

    toLikes(){
        this.likeP = !this.likeP;
    },
    async getCommentId(){
      try{
        let response = await this.$axios.get("/user/"+localStorage.getItem('token')+"/images/"+this.iId+"/comments/");
        // mi preleva tutti i commenti in ordine discendente di id del commento, prendo il primo.
        return response.data[0].idComment
      }catch(e){  
        console.log("Errore"+e)
      }
    },

    async getPost(){
        //funzione
      try{
        this.token = localStorage.getItem('token')
        //this.likes = this.like
        console.log("like ricevuti :"+this.like)
        this.likes = this.like != null ? this.like : []
        console.log("lista like alla foto:"+this.likes)
        this.comments = this.commenti != null ? this.commenti : []
        //console.log(this.comments)
        this.isP = this.isProfile
        if(!this.isP){ 
          this.nomeAutore = await this.idToNameAutore(this.autore) 
        }
        console.log("il nome è:"+this.nomeAutore)
        this.descrizione = this.desc != null ? this.desc : ""

        console.log("id dell'immagine"+this.iId)
        console.log("url?:::"+__API_URL__)
        this.imageUrl = __API_URL__+ "/user/"+this.autore+"/images/"+this.iId
        //const response = await axios.get("/user/"+this.autore+"/images/"+this.iId);
        //this.imageUrl = response.data.file;
        //this.imageInfo = response.data.info;



        //get image, ritorna anche le informazioni dell'immagine, ma prendo solo il file mandato
        //this.imageUrl = __API_URL__+ "/user/"+this.autore+"/images/"+this.iId

        if(this.isLike()){
          this.isLiked = true
        }
      }catch(e){
        this.errore = "{"+ e +"}"
      }
    },

    goToProfile(){
      this.$router.replace("/profile/"+this.autore)
    },

    goToProfileComm(id){
      this.$router.replace("/profile/"+id)
    },

    async idToName(id){
        try {
          let response = await this.$axios.get("/user/"+ id) 
          const nuovoElemento = {
              nome: response.data.name,
              id: response.data.userId
          };
          this.nomiLike.push(nuovoElemento) 
        }catch (e){
            this.errore = "{"+ e +"}"
      }
    },

    async idToNameAutore(id){
        try {
          let response = await this.$axios.get("/user/"+ id) 
          this.idAutore = response.data.userId
          return response.data.name
        }catch (e){
            this.errore = "{"+ e +"}"
      }
      return "Errore"
    },

    async eliminaPost(){
      try {
        await this.$axios.delete("/user/"+ this.autore +"/images/"+this.iId) 
        this.eliminato = true
      }catch (e){
        this.errore = "{"+ e +"}"
      }
    },

  },

  async mounted(){
    await this.getPost();
    if(this.likes !== null){
      for (let i = 0; i < this.likes.length; i++) {
          await  this.idToName(this.likes[i])
    }}
    
  }

};
</script>

<template>
  <div class="principale">
    <div v-if="!isP" @click="toAuthorProfile()" >{{this.nomeAutore}}</div>
    <div v-if="post && !likeP"  class="post">
      <img v-if="!eliminato"  :src="imageUrl"  class="immagine" :title=this.data>
      <p v-if="eliminato" style="color: red;">ricarica la pagina</p>
      <div class="feedback">
        <button @click="toggleLike()" class="like">
            <img v-if = "!isLiked" src="../assets/images/noLike.png">
            <img v-if = "isLiked"  src="../assets/images/like.png">
        </button>
        <input v-model="comment" @keyup.enter="addComment()" placeholder="Aggiungi commento...">
        <button @click="addComment()">
          <img  src="../assets/images/invia.png">
        </button>
      </div>
      <p v-if="!eliminato" id="stiDesc" class="descrizione">{{ descrizione }}</p>      
      <p v-if="eliminato" style="color: red;">post elimanto</p>

      <!-- v-if="errore === ''"  <p v-if="errore !== ''" class="descrizione">{{ errore }}</p> -->
      <div class="com">
        <button @click="toComment()" >commenti</button>
        <button @click="toLikes()" >like</button>
        <button @click="eliminaPost()" v-if="isP">elimina</button>
      </div>   
    </div>
    <div v-show="!post" class="commenti">
      <div  class="testo">
        <CommentsPreview v-for="(commentoa, index) in comments"
          :key="index"
          :uId="commentoa.commenter"
          :cId="commentoa.idComment"
          :testo="commentoa.text"
          :iId="this.iId"
          :autore="this.autore">
        </CommentsPreview>
        <!-- <p  id="comP" class="cambiacolore" v-for="(comments, index) in comments" :key="index" @click="removeComment()">{{ comments }}</p> -->
      </div>
      <div class="com2">
        <button @click="toComment()" >post</button>
      </div>   
    </div>
    <div v-if="likeP" class="commenti">
    
    <div  class="testo">
        <p class="cambiacolore" v-for="(like, index) in nomiLike" :key="index" :id="like.id" @click="goToProfileComm(like.id)"> {{ like.nome }} </p>
      </div>
      <div class="com2">
          <button @click="toLikes()">post</button>
      </div>  
    </div>
  </div>
  </template>
  
  <style scoped>
      @import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

    .principale{
      box-sizing: inherit;
      border: 1px solid #ccc; 
      align-items: center;
      font-family: 'Rubik', sans-serif;
    }
    .post {
        width: 420px; /* Imposta la larghezza del singolo post */
        height: 400px;
        padding: 14px;
    }
    
    .immagine {
        width: 100%;
        height: auto;
        max-height: 250px;

    }
    .commenti{
      width: 420px;
      height: 400px;
      min-height: 40px;
    }
    .descrizione {
        margin-top: 10px;
        color: black;
        overflow-y: scroll;
        height: 35px;
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
      max-height: 85%;
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