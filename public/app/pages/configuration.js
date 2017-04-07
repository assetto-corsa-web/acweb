Vue.component('Configuration', {
	template: '{{include=template/pages/configuration.html}}',
	data: function(){
		return {
			err: 0,
			addEditConfiguration: false,
			removeConfiguration: false,
			saved: false
		}
	},
	methods: {

	}
});
