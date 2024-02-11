<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {

	mounted(){
      if(!localStorage.getItem("token")){
        console.log("Non sei loggato")
        this.$router.replace("/login")
      }else{
        console.log("Questo è il tuo token: "+localStorage.getItem("token"))
      }
    }
}
</script>

<template>
	<main class="">
		<RouterView />
	</main>
</template>
<style>
</style>
