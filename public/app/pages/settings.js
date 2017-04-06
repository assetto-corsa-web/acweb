Vue.component('Settings', {
	template: '{{include=template/pages/settings.html}}',
	data: function(){
		return {
			erpversion: ERP_VERSION,
			copyright: COPYRIGHT,
			tabs: [
				{label: 'Allgemein', show: true},
				{label: 'E-Mail', show: false},
				{label: 'Über emvi ERP', show: false}
			]
		}
	}
});
