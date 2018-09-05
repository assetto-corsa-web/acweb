import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import axios from "axios";

import "../static/main.scss";
import "./global.js";
import * as pages from "./pages";

Vue.use(VueRouter);
Vue.use(Vuex);
Vue.config.productionTip = false;
Vue.config.devtools = false;

// router
const routes = [
	{path: '/', component: pages.Login},
	{path: '/instance', component: pages.Instance},
	{path: '/configuration', component: pages.Configuration},
	{path: '/settings', component: pages.Settings},
	{path: '/user', component: pages.User},
	{path: '/about', component: pages.About},
	{path: '*', component: pages.Instance}
];

let router = new VueRouter({routes, mode: "history"});

// router interceptor to check session for protected pages
router.beforeEach(function(to, from, next){
	axios.get('/api/session')
	.then(resp => {
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

		//SessionService.init(resp.data.data);
		next();
	});
});

// main component
new Vue({
	el: "#app",
	router
});
