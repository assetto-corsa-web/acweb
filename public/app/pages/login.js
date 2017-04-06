Vue.component('Login', {
	template: '{{include=template/pages/login.html}}',
	data: function(){
		return {
			erpversion: ERP_VERSION,
			copyright: COPYRIGHT,
			err: 0,
			client: '',
			login: '',
			pwd: ''
		}
	},
	methods: {
		perfomLogin: function(){
			
		}
	}
});
