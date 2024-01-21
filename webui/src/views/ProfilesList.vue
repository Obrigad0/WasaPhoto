<script>
import PageComponents from '@/components/PageComponents.vue';
import ProfilePreview from '@/components/ProfilePreview.vue';

export default {
  data: function() {
		return {
            typeOfSearch: null,
            query: "",
            users: [
              { username: 'Utente1', follower: 100, following: 50 },
              { username: 'Utente2', follower: 150, following: 75 },
              { username: 'Utente3', follower: 80, following: 40 },
            ], //array utenti
        }
    },

    components: {
      PageComponents,
      ProfilePreview,
    },

    created() { 
    // al caricamento della pagina vengono recuperate tutti gli utenti

    if(this.typeOfSearch === 0){    //ricerca utenti
        this.getProfilesSearch();
    }else if(this.typeOfSearch === 1){
        this.getProfilesFollow();   //get follow 
    }

  },

  async getProfilesSearch(){
    //prend i profili della ricerca
    let response = await this.$axios.get("/user/",{ params: {name: this.query ,},}); //cambiare
          if (response.status === 401 || response.status === 500){
             console.log("Errore, informazioni non recuperabili")
             //mi sposto nella pagina errorpage
             return
          }
    this.users = response.data.profiles != null ? response.data.profiles : [] //attenzione da cambiare
  },

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

};

</script>

<template>
    <PageComponents>
        <div class="profili">
                <ProfilePreview  v-for="(users, index) in users"
                  :key = "index"
                  :nome = "users.username"
                  :follower = "users.follower"
                  :following = "users.following">
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