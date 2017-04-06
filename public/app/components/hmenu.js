Vue.component('hmenu', {
	template: '{{include=template/components/hmenu.html}}',
	data: function(){
		return {
			version: VERSION
		}
	},
	methods: {
		performLogout: function(){
			SessionService.logout(this.$router);
		}
	}
});
