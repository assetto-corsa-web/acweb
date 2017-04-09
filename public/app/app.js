Vue.use(VueRouter);
Vue.use(VueResource);

var router = new VueRouter({
	routes: [
		{
			path: '/',
			component: Vue.component('Login')
		},
		{
			path: '/instance',
			component: Vue.component('Instance')
		},
		{
			path: '/configuration',
			component: Vue.component('Configuration')
		},
		{
			path: '/settings',
			component: Vue.component('Settings')
		},
		{
			path: '/user',
			component: Vue.component('User')
		},
		{
			path: '/about',
			component: Vue.component('About')
		},
		{
			path: '*',
			component: Vue.component('Dashboard')
		}
	]
});

// Simple login check with redirection when not logged in.
router.beforeEach(function(to, from, next){
	Vue.http.get('/api/checkLogin')
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
});

window.onload = function(){
	new Vue({el: '#app', router: router});
};
