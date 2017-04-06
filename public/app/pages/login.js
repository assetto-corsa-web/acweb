Vue.component('Login', {
	template: '{{include=template/pages/login.html}}',
	data: function(){
		return {
			version: VERSION,
			err: 0,
			login: '',
			pwd: ''
		}
	},
	methods: {
		perfomLogin: function(){
			
		}
	}
});
