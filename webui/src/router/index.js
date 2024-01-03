import {createRouter, createWebHashHistory} from 'vue-router'

import Home from '../views/Home.vue'
import ErrorPage from '../views/ErrorPage.vue'
import Login from '../views/Login.vue'
import PostImagePage from '../views/PostImagePage.vue'
import Profile from '../views/Profile.vue'
//import ProfileList from '../views/ProfileList.vue'
import Search from '../views/Search.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: "/login"},
		{path: '/login', component: Login},
		{path: '/user/:idUser/following/', component: Home},
		{path: '/user/:idUser/images', component: PostImagePage},
		{path: '/user/:idUser', component: Profile},
		{path: '/search', component: Search},
		{path: '/:catchAll(.*)', component: ErrorPage},

	]
})

export default router
