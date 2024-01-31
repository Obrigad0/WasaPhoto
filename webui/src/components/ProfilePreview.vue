<script>
export default {
  data() {
    return{

    };
  },

  props: ["nome","follower","following","uId"], 

  methods: {
    goToProfile(){
      this.$router.replace("/profile/"+this.uId)
    }
  }
}
</script>

<template>
    <div class="profile">        
      <div><p class="usnme" @click="goToProfile()">{{ nome }}</p></div>
      <div><p> Follower: {{ follower != null ? follower.length : 0 }}  Following: {{ following != null ? following.length : 0}}</p></div>
    </div>
</template>

<style scoped>
      @import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

    .profile{
        font-family: 'Rubik', sans-serif;
        display: flex;
        flex-direction: row;
        width: 99%;
        background-color: #E6A4B4 ; 
        justify-content: space-evenly;
        margin: 5px;
        border-radius: 19px;
    }
    .profile p{
      font-size: 22px;
    }

    .usnme:hover{
      color: blue;
    }

</style>