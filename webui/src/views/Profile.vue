<script>
import PageComonents from '@/components/PageComponents.vue';
import ImageStructure from '@/components/ImageStructure.vue';
import LoadImage from '@/components/LoadImage.vue';


export default {
  data: function() {
		return {
      modifica: false,
      username: "Username",
      nuovoUsername: "",
      followerN: 0,
      followingN: 0,
      follower: [],
      following: [],
      banList: [],
      images: [],
      errore: null,
      seguito: false, //usati da utente che visita il profilo che non siamo noi
      bannato: false, //usati da utente che visita il profilo che non siamo noi
      //banned: false,
    }
  },

    components: {
      PageComonents,
      ImageStructure,
      LoadImage,
    },

  mounted() { 
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

      toBanList(){
        this.$router.push({ path: '/search', query: { searchValue: "@ban" } });
      },

      toFollow(){
        this.$router.push({ path: '/search', query: { searchValue: "@follow" } });
      },
      
      itsMe(){
        //controlla se l'utente del profilo visitato e' l'utente proprietario del profilo
        if(localStorage.getItem('token') == this.$route.params.idUser){
          console.log("Sono io!")
          return true
        }
        return false
        //return true // da levare
        // la funzione e' utilizzata per mostrare certi elementi solo all'utente proprietario della pagian
      },

      async cambiaUsername() { 
          let newUN = this.nuovoUsername
          if(newUN.length > 3 && newUN.length  <= 16){
            try{
              console.log("username inserito:"+newUN)
              await this.$axios.put("/user/"+this.$route.params.idUser, {username: newUN} , { headers: { 'Content-Type': 'application/json' }});
              location.reload(true)
            }catch(e){
              this.errore = e.toString();
            }
          }else{
            console.log("Errore username troppo corto")
          }
      },
        
      isFollowed(){
        //controlla se l'utente e' nell'array di followed dell'utente visitato 
        //se si il pulsante per seguire diventa per smettere di seguire
        //this.seguito = !this.seguito //da levare
        if(this.follower.includes(parseInt(localStorage.getItem('token'),10))){ console.log("i follower"+this.follower); return true }else{console.log("i follower"+this.follower); return false}
      },
       
      async follow(){ //usato da terzo
        try{
          console.log("cliccato segui") 
          if(this.seguito){
            //leva il follow
            await this.$axios.delete("/user/"+localStorage.getItem('token')+"/following/"+ this.$route.params.idUser);
            this.followerN -= 1
          }else{
            //seguo l'utente
            await this.$axios.put("/user/"+localStorage.getItem('token')+"/following/"+ this.$route.params.idUser);
            this.followerN += 1
          }
          this.seguito = !this.seguito
        }catch(e){
          this.errore = e.toString();
        }
      },
        
      async ban(){  
        try{
          console.log("cliccato banna") 
          if(this.bannato){
            //leva il ban
            await this.$axios.delete("/user/"+localStorage.getItem('token')+"/banned/"+ this.$route.params.idUser);
          }else{
            //banno l'utente
            await this.$axios.put("/user/"+localStorage.getItem('token')+"/banned/"+ this.$route.params.idUser);
          }
          this.bannato = !this.bannato
        }catch(e){
          this.errore = e.toString();
        }
      },
      
      async getBans(){
        try{
            let response = await this.$axios.get("/user/"+localStorage.getItem('token')+"/banned/");
            let dati = response.data
            if(dati){
              console.log("ok!")
            }
        }catch(e){
          this.errore = e.toString();
        }
      },

      async profileInfos(){ // usato al caricamento della pagina
        try{
          //prelevo le informazioni
          let response = await this.$axios.get("/user/"+this.$route.params.idUser);

          if (response.status === 401 || response.status === 500){
            
             return
          }

          this.username = response.data.name
          console.log("USername profilo visitato:"+this.username)
          //controlla se l'array ricevuto e' null, se si viene iserito in follower un array vuoto
          this.follower = response.data.follower != null ? response.data.follower : []
          this.following = response.data.following != null ? response.data.following : []

          // prendo tutte le immagini dell'utente
          console.log("recupero le immagini")
          let response2 = await this.$axios.get("/user/"+this.$route.params.idUser+"/images");
          this.images = response2.data != null ? response2.data : []
          console.log("Ecco i like che ho ricevuto dal server: "+this.images[0])

          //controllo se l'array e' null, se si inserisco 0 in followerN, altrimenti inserisco la lunghezza dell'array
          this.followerN = response.data.follower != null ? response.data.follower.length : 0
          this.followingN = response.data.following != null ? response.data.following.length : 0

          //usata solo se e' il nostro profilo
          this.banList = response.data.banList != null ? response.data.banList: [] 
          //forse

          //controllo se l'utente segue l'utente visitato
          this.seguito =  this.isFollowed();

          //faccio una chiamata al db per il mio profilo dove prendo i miei ban e vedo se ho bannato questo utente.
          if(this.$route.params.idUser !== parseInt(localStorage.getItem('token'),10)){
              console.log("Vedo se ho bannato questo utente")
              try{
                let response2 = await this.$axios.get("/user/"+localStorage.getItem("token"));
                const dban = response2.data.banList != null ? response2.data.banList: [] 
                if(dban.includes(parseInt(this.$route.params.idUser,10))){
                  // ho bannato questo profilo, quindi il pulsante cambia (e il tipo di operazione quando premuto)
                  console.log("Ho bannato questo profilo!")
                  this.bannato = true
                  //testoBottoneBan() ??
                }
            }catch(e){
              console.log(e)
            }

          }

        }catch(e){
           //this.banned = true
           console.log(e)
           console.log("Errore, informazioni non recuperabili o sei stato bannato dall'utente!")
             //mi sposto nella pagina errorpage
           this.$router.replace("/Error")
        }

      }, 
      
      unm(){
        this.modifica = !this.modifica
      }
      

  }


};

</script>

<template>
  <PageComonents> 
    <div class="box" style="height: 13vh;">
        <div class="titoli">
              <h1 v-if="!modifica" id="username">{{username}}</h1>
              <input v-model="nuovoUsername" v-if="modifica" class="input" type="text" maxlength="16" minlength="2" placeholder="inserisci nuovo username">
              <button @click="cambiaUsername()" v-if="modifica" >invia</button>
        </div>
        <div class="titoli2">
              <p style="color: black; margin-right: 10px; margin-top: 15px; font-size: 22px;">Follower: {{followerN}} Followed: {{followingN}}</p>
              <button @click="follow()" v-if="!itsMe()" class="operazioni">{{testoBottoneFollow()}}</button>
              <button @click="ban()" v-if="!itsMe()" class="operazioni">{{testoBottoneBan()}}</button>
              <!--<button @click="" v-if="itsMe()" class="operazioni">BANLIST</button>-->
              <button  @click="unm()" v-if="itsMe()" class="operazioni">Username</button>
              <button  @click="toBanList()" v-if="itsMe()" class="operazioni">Ban list</button>
              <!--<button  @click="toFollow()" v-if="itsMe()" class="operazioni">Follow</button>-->


        </div>
    </div>
    <div class="box2" style="height: 72vh;">
    
        <ImageStructure v-for="(post, index) in images"
            :key="index"
            :autore="post.author"
            :like="post.like"
            :commenti="post.comments"
            :data="post.data"
            :iId="post.imgId"
            :desc="post.descrizione"
            :isProfile=itsMe()>
        </ImageStructure>
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
    height: 30px;
    width: 125px;
    margin-right: 10px;
    border-radius: 10px;
  }


</style>

<!-- {
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
        }, -->