import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import axios from "axios";

import "../static/main.scss";
import * as pages from "./pages";

Vue.use(VueRouter);
Vue.use(Vuex);

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
	{path: "/", component: pages.Home},
	{path: "*", component: pages.Home} // TODO 404
];

let router = new VueRouter({routes, mode: "history"});

// router interceptor to check token for protected pages
router.beforeEach((to, from, next) => {
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

// main component
new Vue({
	el: "#app",
	router
});
