import {createApp, reactive} from 'vue'

import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Image from  './components/Image.vue'
import PageComponents from  './components/PageComponents.vue'
import ProfilePreview from  './components/ProfilePreview.vue'
import CommentsPreview from  './components/CommentsPreview.vue'
import LoadImage from  './components/LoadImage.vue'


//import from './components/.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Image",Image);
app.component("PageComponents",PageComponents);
app.component("ProfilePreview",ProfilePreview);
app.component("CommentsPreview",CommentsPreview);
app.component("LoadImage",LoadImage);


//app.component("",);

app.use(router)
app.mount('#app')
