Vue.component('hmenu', {
	template: '{{include=template/components/hmenu.html}}',
	props: ['admin'],
	methods: {
		performLogout: function(){
			SessionService.logout(this.$router);
		}
	}
});
