<script>
import PageComponents from '@/components/PageComponents.vue';
import ProfilePreview from '@/components/ProfilePreview.vue';

export default {
  data: function() {
		return {
            typeOfSearch: null,
            query: "",
            users: [], //array utenti
        }
    },

    components: {
      PageComponents,
      ProfilePreview,
    },


  methods: {
      async getProfilesSearch(){
        //prend i profili della ricerca
        let response = await this.$axios.get("/user", { params: {queryNomeUtente: this.query ,},} ); //cambiare
              if (response.status === 401 || response.status === 500){
                console.log("Errore, informazioni non recuperabili")
                //mi sposto nella pagina errorpage
                this.$router.replace("/Error")
                return
              }
        this.users = response.data != null ? response.data : [] //gestione se array vuoto?
        console.log("ecco gli utenti"+this.users)
      },
      
      async getProfilesBan(){
        //prende i profili seguiti
        let response = await this.$axios.get("/user/"+localStorage.getItem("token")+"/banned/"); //cambiare
              if (response.status === 401 || response.status === 500){
                console.log("Errore, informazioni non recuperabili")
                //mi sposto nella pagina errorpage
                return
              }
        this.users = response.data != null ? response.data : [] //attenzione da cambiare
      },
      
      //FUNZIONE NON UTILIZZABILE
      /*async getProfilesFollow(){
        //prende i profili seguiti
        let response = await this.$axios.get("/user/"+localStorage.getItem("token")+"/following/"); //cambiare
              if (response.status === 401 || response.status === 500){
                console.log("Errore, informazioni non recuperabili")
                //mi sposto nella pagina errorpage
                return
              }
        this.users = response.data.utente != null ? response.data.utente : [] //attenzione da cambiare
      },
      */
        
  },
 
  async mounted(){
    this.query = this.$route.query.searchValue || ''
    console.log("Ecco la query "+this.query)
    if(this.query == "@ban"){
      await this.getProfilesBan()
    }else{
      await this.getProfilesSearch()
    }
  }
};

</script>

<template>
    <PageComponents>
        <div v-if="this.users.length" class="profili">
                <ProfilePreview  v-for="(users, index) in users"
                  :key = "index"
                  :nome = "users.name"
                  :follower = "users.follower"
                  :following = "users.following"
                  :uId = "users.userId">
                </ProfilePreview>
        </div>
        <div v-if="!this.users.length" class="profili">
          <p>Non ho trovato nessun utente</p>
        </div>

    </PageComponents> 
</template>

<style scoped>
    .profili{
        overflow-y: auto;
        margin-top: 30px;
        height: 92%;
        width: 95%;
    }
</style>