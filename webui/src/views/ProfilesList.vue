<script>
import PageComponents from '@/components/PageComponents.vue';
import ProfilePreview from '@/components/ProfilePreview.vue';

export default {
  data: function() {
		return {
            typeOfSearch: null,
            query: "",
            users: [], //array utenti
            //{ username: 'Utente1', follower: 100, following: 50 },
            //{ username: 'Utente2', follower: 150, following: 75 },
            //{ username: 'Utente3', follower: 80, following: 40 },
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
        this.users = response.data //gestione se array vuoto?
        console.log("ecco gli utenti"+this.users)
      },
       /*
       async getProfilesFollow(){
        //prende i profili seguiti
        let response = await this.$axios.get("/user/"+this.$route.params.idUser+"/following/"); //cambiare
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
    await this.getProfilesSearch()
  }
};

</script>

<template>
    <PageComponents>
        <div class="profili">
                <ProfilePreview  v-for="(users, index) in users"
                  :key = "index"
                  :nome = "users.username"
                  :follower = "users.follower"
                  :following = "users.following"
                  :uId = "null">
                </ProfilePreview>
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