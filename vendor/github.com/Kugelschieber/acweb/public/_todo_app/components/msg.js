Vue.component('msg', {
	props: ['type', 'msg'],
	template: '{{include=template/components/msg.html}}'
});
