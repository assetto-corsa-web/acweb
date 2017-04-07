Vue.component('Settings', {
	template: '{{include=template/pages/settings.html}}',
	data: function(){
		return {
			err: 0,
			folder: '',
			cmd: '',
			saved: false
		}
	},
	mounted: function(){
		this._load();
	},
	methods: {
		_load: function(){
			this.$http.get('/api/getSettings')
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					return;
				}

				this.folder = resp.data.folder;
				this.cmd = resp.data.cmd;
			});
		},
		performSave: function(){
			this.saved = false;

			this.$http.post('/api/saveSettings', {folder: this.folder,
				cmd: this.cmd})
			.then(function(resp){
				if(resp.data.code){
					console.log(resp.data.code+': '+resp.data.msg);
					this.err = resp.data.code;
					return;
				}

				this.saved = true;
			});
		}
	}
});
