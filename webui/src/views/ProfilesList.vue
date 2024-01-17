<script>
import PageComponents from '@/components/PageComponents.vue';
import ProfilePreview from '@/components/ProfilePreview.vue';

export default {
  data: function() {
		return {
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
    // if e' una ricerca utente...
        //this.getProfilesSearch();
    // else e' una visualizzazione di follower/following
        //this.getProfilesFollow();

  },

  async getProfilesSearch(query){
    //prende i profili della ricerca
    //this.users = //
  },

  async getProfilesFollow(idUtente){
    //prende i profili seguiti
    //this.users = //
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