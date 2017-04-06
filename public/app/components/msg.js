Vue.component('msg', {
	props: ['type', 'msg'],
	template: '{{include=template/components/msg.html}}',
	data: function(){
		return {
			visible: true
		}
	},
	methods: {
		close: function(){
			this.visible = false;
		}
	}
});
