<script>
import PageComponents from '@/components/PageComponents.vue';
import ImageStructure from '@/components/ImageStructure.vue';


export default {
  data: function() {
		return {
        arrayDiPost: [],
        nome: '',
        }
    },

    components: {
      PageComponents,
      ImageStructure,
    },

    methods:{
      async getMyStream(){
          try{

            let response = await this.$axios.get("/user/"+localStorage.getItem("token")+"/following/");
            let data = response.data
            
            //let user = data.utente; non servono
            // e' un array di immagini
            this.arrayDiPost = data.ultimaImmagine;

          }catch(e){
            console.log(this.e);
          }
      }
    },

    async mounted(){
      await this.getMyStream()
    }

}
</script>


<template>
    <PageComponents>
      <div class="tutto">
        <div class="destra">
          <p>Ecco i post dei tuoi amici:</p>
        </div>
        <div class="centro">
          <ImageStructure v-for="(post, index) in arrayDiPost"
                :key="index"
                :autore="post.author"
                :like="post.like"
                :commenti="post.comments"
                :data="post.data"
                :iId="post.imgId"
                :desc="post.descrizione"
                :isProfile="false">
          </ImageStructure>
        </div>
        <div class="sinistra">
        </div>
      </div>
    </PageComponents> 
</template>


<style>
@import url('https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap');

  .tutto{
    font-family: 'Rubik', sans-serif;
    margin-top: 10px;
    display: flex;    
    height: 100%;
    width: 100%;
  }

  .destra{
    height: 100%;
    width: 20%;
    text-align: center;
  }

  .centro{
    border: 1px solid #ccc; 
    border-top: 0px;
    border-bottom: 0px;
    height: 100%;
    width: 60%;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 2%;
  }


  .sinistra{
    height: 100%;
    width: 20%;
  }
</style>

<!-- {
                autore: 'Autore1',
                like: [1,3,4,5,6,7,7,8],
                commenti: ['Commento 1', 'Commento 2'],
                data: '',
                iId: 1,
                desc: 'Descrizione del post 1',
                url: 'https://www.laziostylestore.com/images/lazio/products/large/LZ22A03_16.webp'
             },
             {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 2',
                iId: 1,
                desc: 'Descrizione del post 2',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 3',
                iId: 1,
                desc: 'Descrizione del post 3',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 4',
                iId: 1,
                desc: 'Descrizione del post 4',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },
            {
                autore: 'Autore2',
                like: [1,5,45,71,5,33,8],
                commenti: ['Commento 3', 'Commento 4'],
                data: 'Data del post 5',
                iId: 1,
                desc: 'Descrizione del post 5',
                url: 'https://www.laziostylestore.com/images/lazio/products/small/LZ22A03_9.webp'
            },-->