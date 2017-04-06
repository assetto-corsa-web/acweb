Vue.use(VueRouter);
Vue.use(VueResource);

// Configure vue-resource to parse response as JSON.
Vue.http.interceptors.push(function(req, next){
	next(function(resp){
		resp.body = JSON.parse(resp.body);
	});
});

var router = new VueRouter({
	routes: [
		{
			path: '/',
			component: Vue.component('Login')
		},
		{
			path: '/dashboard',
			component: Vue.component('Dashboard')
		},
		{
			path: '/user',
			component: Vue.component('User')
		},
		{
			path: '/settings',
			component: Vue.component('Settings')
		},
		{
			path: '/profile',
			component: Vue.component('Profile')
		},
		{
			path: '/password',
			component: Vue.component('Password')
		},
		{
			path: '*',
			component: Vue.component('Dashboard')
		}
	]
});

// Simple login check with redirection when not logged in.
/*router.beforeEach(function(to, from, next){
	Vue.http.post('/api/checkLogin', {})
	.then(function(resp){
		// if not logged in and not on login/password page, redirect to login
		if(resp.data.code && to.path != '/' && to.path != '/password'){
			next('/');
			return;
		}

		// if login or password page, but logged in, redirect to start page
		if(!resp.data.code && (to.path == '/' || to.path == '/password')){
			next('/dashboard');
			return;
		}

		SessionService.init(resp.data.data);
		next();
	});
});*/

window.onload = function(){
	new Vue({el: '#app', router: router});
};
