<script>
import PageComonents from '@/components/PageComponents.vue';
import Image from '@/components/Image.vue';
import LoadImage from '@/components/LoadImage.vue';


export default {
  data: function() {
		return {
      arrayDiPost: [
        {
          autore: 'Autore1',
          like: [1,3,4,5,6,7,7,8],
          commenti: [
            { testo: 'Commento 1', cId: 1, uId: 101 },
            { testo: 'Commento 2', cId: 2, uId: 102 },
            { testo: 'Commento 2', cId: 2, uId: 102 },
            { testo: 'Commento 2', cId: 2, uId: 102 }
          ],
          data: '',
          iId: 1,
          desc: 'Descrizione del post 1',
          url: 'https://assets.goal.com/v3/assets/bltcc7a7ffd2fbf71f5/bltc1d5f3f76cbaa905/64ba66893452d45bae678aca/lazio_HD.jpg?auto=webp&format=pjpg&width=3840&quality=60'
        },
        {
          autore: 'Autore2',
          like: [1,5,45,71,5,33,8],
          commenti: [
            { testo: 'Commento 1', cId: 1, uId: 101 },
            { testo: 'Commento 2', cId: 2, uId: 102 }
          ],
          data: 'Data del post 2',
          iId: 1,
          desc: 'Descrizione del post 2',
          url: 'https://assets.goal.com/v3/assets/bltcc7a7ffd2fbf71f5/blt214f00e4cf476695/643a3a34ebdc0d4ea06f21e3/Lazio.jpg?auto=webp&format=pjpg&width=3840&quality=60'
        },
        {
          autore: 'Autore2',
          like: [1,5,45,71,5,33,8],
          commenti: [
            { testo: 'Commento 1', cId: 1, uId: 101 },
            { testo: 'Commento 2', cId: 2, uId: 102 }
          ],
          data: 'Data del post 2',
          iId: 1,
          desc: 'Descrizione del post 2',
          url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
        },
        {
          autore: 'Autore2',
          like: [1,5,45,71,5,33,8],
          commenti: [
            { testo: 'Commento 1', cId: 1, uId: 101 },
            { testo: 'Commento 2', cId: 2, uId: 102 }
          ],
          data: 'Data del post 2',
          iId: 1,
          desc: 'Descrizione del post 2',
          url: 'https://cdn.studenti.stbm.it/images/2022/10/07/ricerca-sulla-regione-lazio-orig.jpeg'
        },
      ],
      
      username: "Username",
      followerN: 0,
      followingN: 0,
      follower: [],
      following: [],
      images: [],
      errore: null,
      seguito: false, //usati da utente che visita il profilo che non siamo noi
      bannato: false, //usati da utente che visita il profilo che non siamo noi
      banned: false,

    }
  },
    components: {
      PageComonents,
      Image,
      LoadImage,
    },

  created() { 
    // al caricamento della pagina vengono recuperate tutte le informazioni sull'utente
    this.profileInfos();
  },

  methods:{
      testoBottoneFollow() {
        return this.seguito ? 'UnFollow' : 'Follow';
      },

      testoBottoneBan() {
        return this.bannato ? 'UnBan' : 'Ban';
      },
      
      itsMe(){
        //controlla se l'utente del profilo visitato e' l'utente proprietario del profilo
        //return localStorage.getItem('token') === this.$route.params.id
        return true // da levare
        // la funzione e' utilizzata per mostrare certi elementi solo all'utente proprietario della pagian
      },

      cambiaUsername() { 

      },
        
      isFollowed(){
        //controlla se l'utente e' nell'array di followed dell'utente visitato 
        //se si il pulsante per seguire diventa per smettere di seguire
        this.seguito = !this.seguito //da levare
        return this.follower.includes(localStorage.getItem('token'))
      },
       
      async follow(){ //usato da terzo
        this.seguito = !this.seguito //da levare
        try{
          console.log("cliccato segui") 
          if(this.seguito){
            //leva il follow
            this.followerN += 1
          }else{
            //seguo l'utente
            this.followerN -= 1
          }
          //this.seguito = !this.seguito
        }catch(e){
          this.errore = e.toString();
        }
      },
        
      async ban(){  //usato da terzo
        this.bannato = !this.bannato  //da levare
        try{
          console.log("cliccato banna") 
          if(this.bannato){
            //leva il ban
          }else{
            //banno l'utente
          }
          //this.bannato = !this.bannato
        }catch(e){
          this.errore = e.toString();
        }
      },
      
      
      async profileInfos(){ //usato al caricamento della pagina
        try{
          //prelevo le informazioni
          let response = await this.$axios.get("/users/"+this.$route.params.idUser);

          if (response.status === 401 || response.status === 500){
             console.log("Errore, informazioni non recuperabili")
             this.banned = true
             //mi sposto nella pagina errorpage
             return
          }

          this.username = response.data.username
          //controlla se l'array ricevuto e' null, se si viene iserito in follower un array vuoto
          this.follower = response.data.follower != null ? response.data.follower : []
          this.following = response.data.following != null ? response.data.following : []
          this.images = response.data.images != null ? response.data.images : []
          //controllo se l'array e' null, se si inserisco 0 in followerN, altrimenti inserisco la lunghezza dell'array
          this.followerN = response.data.follower != null ? response.data.follower.length : 0
          this.followingN = response.data.following != null ? response.data.following.length : 0
          //controllo se l'utente segue l'utente visitato
          this.seguito =  isFollowed();


        }catch(e){
           //mi sposto nella pagina errorpage

          this.banned = true
        }

      }, 
      
      

  }


};

</script>

<template>
  <PageComonents> 
    <div class="box" style="height: 13vh;">
        <div class="titoli">
              <h1 id="username">{{username}}</h1>
        </div>
        <div class="titoli2">
              <p style="color: black; margin-right: 10px;">follower: {{followerN}}     followed: {{followingN}}</p>
              <button @click="follow()" v-if="!itsMe()" class="operazioni">{{testoBottoneFollow()}}</button>
              <button @click="ban()" v-if="!itsMe()" class="operazioni">{{testoBottoneBan()}}</button>
              <button @click="" v-if="itsMe()" class="operazioni">BANLIST</button>
              <button @click="" v-if="itsMe()" class="operazioni">Username</button>
              <button @click="" v-if="itsMe()" class="operazioni">Posta Foto</button>
        </div>
    </div>
    <div class="box2" style="height: 72vh;">
        <Image v-for="(post, index) in arrayDiPost"
            :key="index"
            :autore="post.autore"
            :like="post.like"
            :commenti="post.commenti"
            :data="post.data"
            :iId="post.iId"
            :desc="post.desc"
            :url="post.url"
            :isProfile="true"
            :username="null">
        </Image>
        <LoadImage v-if="itsMe()"></LoadImage>
    </div>
  </PageComonents>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

  body, html {

    margin: 0;
    padding: 0;
    height: 100%;
  }

  .box {
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    color: #fff;
    font-size: 20px;
    box-sizing: border-box;
  }

  .box2 {
    border: 1px solid #ccc; 
    display: flex;
    justify-content: space-evenly;
    gap: 2%;
    box-sizing: border-box;
    overflow-y: auto;
    flex-wrap: wrap;
    align-items: center;
  }

  #username{
    text-align: center;
    font-size: 5vw;
    color: white; 
    text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000; /* Ombra per simulare il bordo nero con dimensione 2px */
  
  }
  .titoli {
    flex: 1;
    height: 13vh; 
    box-sizing: border-box;
    align-items: center;
    justify-content: center;
    display: flex;
  }

  .titoli2 {
    flex: 1;
    height: 13vh; 
    box-sizing: border-box;
    align-items: center;
    justify-content: center;
  }

  .operazioni{
    height: 25px;
    width: 100px;
    margin-right: 10px;
    border-radius: 10px;
  }

</style>
