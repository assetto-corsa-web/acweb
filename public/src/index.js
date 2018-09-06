import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import axios from "axios";

import "../static/main.scss";
import "./global.js";
import NewSessionStore from "./store/session.js";
import * as pages from "./pages";

Vue.use(VueRouter);
Vue.use(Vuex);
Vue.config.productionTip = false;
Vue.config.devtools = false;

// router
const routes = [
	{path: "/", component: pages.Login},
	{path: "/instance", component: pages.Instance, meta: {protected: true}},
	{path: "/configuration", component: pages.Configuration, meta: {protected: true}},
	{path: "/settings", component: pages.Settings, meta: {protected: true}},
	{path: "/user", component: pages.User, meta: {protected: true}},
	{path: "/about", component: pages.About, meta: {protected: true}},
	{path: "*", component: pages.Error404}
];

let router = new VueRouter({routes, mode: "history"});

// session storage
let sessionStorage = NewSessionStore();

// router interceptor to check session for protected pages
router.beforeEach(function(to, from, next){
	window.scrollTo(0, 0);

	if(to.meta.protected){
		axios.get("/api/session")
		.then(resp => {
			// non 0 number means user is not logged in, then redirect to login
			if(resp.data.code){
				sessionStorage.commit("logout");
				next("/");
				return;
			}

			next();
		});
	}
	else{
		next();
	}
});

// main component
new Vue({
	el: "#app",
	router,
	store: sessionStorage
});
