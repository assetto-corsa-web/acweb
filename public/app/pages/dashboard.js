Vue.component('Dashboard', {
	template: '{{include=template/pages/dashboard.html}}',
	data: function(){
		return {
			err: 0,
			started: false,
			startInstance: false,
			stopInstance: false
		}
	},
	methods: {
		
	}
});
