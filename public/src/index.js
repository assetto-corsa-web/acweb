import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import axios from "axios";

import "../static/main.scss";
import * as pages from "./pages";

Vue.use(VueRouter);
Vue.use(Vuex);
Vue.config.productionTip = false;
Vue.config.devtools = false;

// token interceptor for every request
axios.interceptors.request.use((config) => {
	const token = window.localStorage.getItem("token");

	if(token){
		config.headers.Authorization = `Bearer ${token}`;
	}

	return config;
}, (err) => {
	return Promise.reject(err);
});

// storage
// ...

// router
const routes = [
	{path: '/', component: Vue.component('Login')},
	{path: '/instance', component: Vue.component('Instance')},
	{path: '/configuration', component: Vue.component('Configuration')},
	{path: '/settings', component: Vue.component('Settings')},
	{path: '/user', component: Vue.component('User')},
	{path: '/about', component: Vue.component('About')},
	{path: '*', component: Vue.component('Dashboard')}
];

let router = new VueRouter({routes, mode: "history"});

// router interceptor to check token for protected pages
/*router.beforeEach((to, from, next) => {
	if(to.meta.protected){
		axios.get("http://localhost/auth/token") // TODO
		.then((r) => {
			next();
		})
		.catch((e) => {
			next("/");
		});
	}
	else{
		next();
	}
});
router.beforeEach(function(to, from, next){
	Vue.http.get('/api/session')
	.then(function(resp){
		// if not logged in and not on login page, redirect to login
		if(resp.data.code && to.path != '/'){
			next('/');
			return;
		}

		// if login or password page, but logged in, redirect to start page
		if(!resp.data.code && to.path == '/'){
			next('/instance');
			return;
		}

		SessionService.init(resp.data.data);
		next();
	});
});*/

// main component
new Vue({
	el: "#app",
	router
});
