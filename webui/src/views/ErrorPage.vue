<script>
import PageComponents from '@/components/PageComponents.vue';

export default {
  data: function() {
		return {

        }
    },
    components: {
      PageComponents,
    },

    toHome(){
        this.$router.replace("/users/"+localStorage.getItem('token')+"/following/")
    },
}
</script>

<template>
    <PageComponents>
        <div class="errore">
             <img src="../assets/images/ErroR.jpg" title="Sfortunatamente non riusciamo a gestirlo...">
             <p style="color: red;">C'è stato un errore!</p>
             <p style="font-size: 30px;">Prova a tornare alla pagina <u class="uho" @click="toHome()">principale</u></p>
        </div>
    </PageComponents> 
</template>


<style scoped>
      @import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

    .errore{
        overflow-y: auto;
        margin-top: 30px;
        height: 95%;
        width: 95%;
        font-family: 'Rubik', sans-serif;
        font-size: 40px;
    }
    .errore img{
        padding-top: 7%;
    }
    .uho:hover{
        color: blue;
    }
</style>