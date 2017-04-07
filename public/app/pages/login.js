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
		performLogin: function(){
			this.$http.post('/api/login', {login: this.login, pwd: this.pwd})
			.then(function(resp){
				if(resp.data.code){
					this.err = resp.data.code;
					return;
				}

				SessionService.login(this.$router, resp.data);
			});
		}
	}
});
