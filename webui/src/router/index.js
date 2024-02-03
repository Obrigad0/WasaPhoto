import {createRouter, createWebHashHistory} from 'vue-router'

import Home from '../views/Home.vue'
import ErrorPage from '../views/ErrorPage.vue'
import Login from '../views/Login.vue'
import Profile from '../views/Profile.vue'
//import ProfileList from '../views/ProfileList.vue'
import Search from '../views/ProfilesList.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: "/login"},
		{path: '/login', component: Login},
		{path: '/home', component: Home},
		{path: '/profile/:idUser', component: Profile},
		{path: '/search', component: Search}, 
		{path: '/:catchAll(.*)', component: ErrorPage},

	]
})

export default router
