Vue.component('Dashboard', {
	template: '{{include=template/pages/dashboard.html}}',
	data: function(){
		return {
			erpversion: ERP_VERSION,
			copyright: COPYRIGHT
		}
	}
});
