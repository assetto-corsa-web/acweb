Vue.component('User', {
	template: '{{include=template/pages/user.html}}',
	data: function(){
		return {
			tabs: [
				{label: 'Nutzer', show: true},
				{label: 'Rollen', show: false}
			]
		}
	}
});
